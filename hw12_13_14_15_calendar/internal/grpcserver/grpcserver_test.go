package grpcserver

import (
	context "context"
	"log"
	"net"
	"reflect"
	"testing"
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

var tests_add = []TestCase{
	{
		Name:           "First test",
		UUID:           "UUID-1",
		DateTimeString: "2020-10-06 00:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           true,
		// ErrCode:        codes.OK,
	},
	{
		Name:           "Second test",
		UUID:           "UUID-2",
		DateTimeString: "2020-10-06 01:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           true,
	},
	{
		Name:           "Event to delete",
		UUID:           "UUID-3",
		DateTimeString: "2020-10-08 14:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           true,
	},
}

var tests_add_negative = []TestCase{
	{
		Name:           "First test",
		UUID:           "UUID-1",
		DateTimeString: "2020-10-06 00:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           false,
		ErrCode:        codes.Internal,
	},
	{
		Name:           "Second test",
		UUID:           "UUID-2",
		DateTimeString: "2020-10-06 01:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           false,
		ErrCode:        codes.Internal,
	},
	{
		Name:           "Event to delete",
		UUID:           "UUID-3",
		DateTimeString: "2020-10-08 14:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           false,
		ErrCode:        codes.Internal,
	},
}

var tests_modify = []TestCase{
	{
		Name:           "First test",
		UUID:           "UUID-1",
		DateTimeString: "2020-10-06 02:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Minute * 30,
		IsOk:           true,
	},
	{
		Name:           "Second test",
		UUID:           "UUID-2",
		DateTimeString: "2020-10-07 03:00:00",
		UserID:         "USER-UUID-1",
		Duration:       time.Hour,
		IsOk:           true,
	},
}

var tests_delete = []TestCaseDeletion{
	{
		Name: "Delete UUID-3",
		UUID: "UUID-3",
		IsOk: true,
	},
}

type TestCase struct {
	Name           string
	UUID           string
	DateTimeString string
	UserID         string
	Duration       time.Duration
	IsOk           bool
	ErrCode        codes.Code
}

type TestCaseDeletion struct {
	Name string
	UUID string
	IsOk bool
}

func dialer() func(context.Context, string) (net.Conn, error) {
	listener := bufconn.Listen(1024 * 1024)

	server := grpc.NewServer()

	RegisterCalendarServer(server, &CalendarService{strg: storage.NewInMemoryStorage()})

	go func() {
		if err := server.Serve(listener); err != nil {
			log.Fatal(err)
		}
	}()

	return func(context.Context, string) (net.Conn, error) {
		return listener.Dial()
	}
}

func TestCalendarService(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)
	defer conn.Close()

	client := NewCalendarClient(conn)
	require.NotNil(t, client)

	for _, tt := range tests_add {
		t.Run(tt.Name, func(t *testing.T) {
			dt, err := time.Parse("2006-01-02 15:04:05", tt.DateTimeString)
			require.NoError(t, err)

			event := &Event{
				Uuid:         tt.UUID,
				Header:       "Some header",
				Datetime:     (*timestamp.Timestamp)(timestamppb.New(dt)),
				Duration:     tt.Duration.String(),
				Description:  "Some description",
				UserUuid:     tt.UserID,
				NotifyBefore: (time.Hour * 24).String(),
			}

			response, err := client.AddEvent(ctx, event)
			if response != nil {
				require.Equal(t, tt.IsOk, response.GetOk())
			}
			require.NoError(t, err)
		})
	}

	for _, tt := range tests_modify {
		t.Run(tt.Name, func(t *testing.T) {
			dt, err := time.Parse("2006-01-02 15:04:05", tt.DateTimeString)
			require.NoError(t, err)

			event := &Event{
				Uuid:         tt.UUID,
				Header:       "Some header",
				Datetime:     (*timestamp.Timestamp)(timestamppb.New(dt)),
				Duration:     tt.Duration.String(),
				Description:  "Some description",
				UserUuid:     tt.UserID,
				NotifyBefore: (time.Hour * 24).String(),
			}

			response, err := client.ModifyEvent(ctx, event)
			require.NoError(t, err)
			require.Equal(t, tt.IsOk, response.GetOk())
		})
	}

	// DeleteEvent
	for _, tt := range tests_delete {
		t.Run(tt.Name, func(t *testing.T) {
			response, err := client.DeleteEvent(ctx, &EventUUID{Uuid: tt.UUID})
			require.NoError(t, err)
			require.NotNil(t, response)
			require.Equal(t, tt.IsOk, response.GetOk())
		})
	}

	// ListDay
	t.Run("ListDay-1", func(t *testing.T) {
		response, err := client.ListDay(ctx, &Day{Day: "2020-10-06"})
		require.NoError(t, err)

		require.NotNil(t, response)
		require.Equal(t, 1, len(response.Events))

		dt, err := time.Parse("2006-01-02 15:04:05", tests_modify[0].DateTimeString)
		require.NoError(t, err)
		event := &Event{
			Uuid:         tests_modify[0].UUID,
			Header:       "Some header",
			Datetime:     (*timestamp.Timestamp)(timestamppb.New(dt)),
			Duration:     tests_modify[0].Duration.String(),
			Description:  "Some description",
			UserUuid:     tests_modify[0].UserID,
			NotifyBefore: (time.Hour * 24).String(),
		}

		require.True(t, reflect.DeepEqual(event, response.Events[0]))
	})

	// ListWeek
	t.Run("ListWeek-1", func(t *testing.T) {
		response, err := client.ListWeek(ctx, &Day{Day: "2020-10-05"})
		require.NoError(t, err)

		require.NotNil(t, response)
		require.Equal(t, 2, len(response.Events))
	})

	// ListMonth
	t.Run("ListMonth-1", func(t *testing.T) {
		response, err := client.ListMonth(ctx, &Day{Day: "2020-10-01"})
		require.NoError(t, err)

		require.NotNil(t, response)
		require.Equal(t, 2, len(response.Events))
	})
}

func TestCalendarServiceNegative(t *testing.T) {
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, "", grpc.WithInsecure(), grpc.WithContextDialer(dialer()))
	require.NoError(t, err)
	defer conn.Close()

	client := NewCalendarClient(conn)
	require.NotNil(t, client)

	for _, tt := range tests_add {
		t.Run(tt.Name, func(t *testing.T) {
			dt, err := time.Parse("2006-01-02 15:04:05", tt.DateTimeString)
			require.NoError(t, err)

			event := &Event{
				Uuid:         tt.UUID,
				Header:       "Some header",
				Datetime:     (*timestamp.Timestamp)(timestamppb.New(dt)),
				Duration:     tt.Duration.String(),
				Description:  "Some description",
				UserUuid:     tt.UserID,
				NotifyBefore: (time.Hour * 24).String(),
			}

			response, err := client.AddEvent(ctx, event)
			if response != nil {
				require.Equal(t, tt.IsOk, response.GetOk())
			}
			require.NoError(t, err)
		})
	}

	for _, tt := range tests_add_negative {
		t.Run(tt.Name, func(t *testing.T) {
			dt, err := time.Parse("2006-01-02 15:04:05", tt.DateTimeString)
			require.NoError(t, err)

			event := &Event{
				Uuid:         tt.UUID,
				Header:       "Some header",
				Datetime:     (*timestamp.Timestamp)(timestamppb.New(dt)),
				Duration:     tt.Duration.String(),
				Description:  "Some description",
				UserUuid:     tt.UserID,
				NotifyBefore: (time.Hour * 24).String(),
			}

			response, err := client.AddEvent(ctx, event)
			require.Equal(t, tt.IsOk, response.GetOk())
			if err != nil {
				if er, ok := status.FromError(err); ok {
					require.Equal(t, tt.ErrCode, er.Code())
				}
			}
		})
	}
}

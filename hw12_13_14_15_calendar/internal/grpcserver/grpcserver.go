// go gen: protoc --go_out=plugins=grpc:internal/grpcserver --go_opt=paths=source_relative --proto_path=api api/calendar.proto
package grpcserver

import (
	context "context"
	"time"

	"google.golang.org/grpc"

	timestamppb "google.golang.org/protobuf/types/known/timestamppb"

	codes "google.golang.org/grpc/codes"
	peer "google.golang.org/grpc/peer"
	status "google.golang.org/grpc/status"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"
	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/storage"
	"github.com/pkg/errors"
)

func LogInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Before
		t1 := time.Now()

		h, err := handler(ctx, req)

		// After
		t2 := time.Now()
		remoteAddr := ""
		if p, ok := peer.FromContext(ctx); ok {
			remoteAddr = p.Addr.String()
		} else {
			remoteAddr = "UNKNOWN"
		}

		// logger.Info(r.RemoteAddr, r.Method, r.RequestURI, r.Proto, srw.statusCode, t2.Sub(t1), r.UserAgent())
		logger.Info(remoteAddr, "GRPC", t2.Sub(t1))

		return h, err
	}
}

type CalendarService struct {
	strg storage.Storager
}

func (cs *CalendarService) AddEvent(ctx context.Context, event *Event) (*Result, error) {
	stEvent, err := pbEventToStorage(event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = cs.strg.Add(stEvent)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &Result{Ok: true}, nil
}
func (cs *CalendarService) ModifyEvent(ctx context.Context, event *Event) (*Result, error) {
	stEvent, err := pbEventToStorage(event)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	err = cs.strg.Modify(stEvent.UUID, stEvent)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &Result{Ok: true}, nil
}
func (cs *CalendarService) DeleteEvent(ctx context.Context, eventUUID *EventUUID) (*Result, error) {
	err := cs.strg.Delete(eventUUID.Uuid)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &Result{Ok: true}, nil
}
func (cs *CalendarService) ListDay(ctx context.Context, day *Day) (*ListEvents, error) {
	return commonList(day, cs.strg.ListDay)
}
func (cs *CalendarService) ListWeek(ctx context.Context, day *Day) (*ListEvents, error) {
	return commonList(day, cs.strg.ListWeek)
}
func (cs *CalendarService) ListMonth(ctx context.Context, day *Day) (*ListEvents, error) {
	return commonList(day, cs.strg.ListMonth)
}

func commonList(day *Day, fn func(time.Time) []storage.Event) (*ListEvents, error) {
	dayTime, err := time.Parse("2006-01-02", day.Day)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}

	events := fn(dayTime)
	result := []*Event{}
	for _, e := range events {
		result = append(result, pbEventToGRPC(e))
	}

	return &ListEvents{Events: result}, nil
}

func pbEventToStorage(event *Event) (*storage.Event, error) {
	duration, err := time.ParseDuration(event.Duration)
	if err != nil {
		return nil, errors.Wrap(err, "error during parsing Duration")
	}
	notifyBefore, err := time.ParseDuration(event.NotifyBefore)
	if err != nil {
		return nil, errors.Wrap(err, "error during parsing NotifyBefore")
	}

	return &storage.Event{
		UUID:               event.Uuid,
		Header:             event.Header,
		DateTime:           event.Datetime.AsTime(),
		Duration:           storage.TDuration(duration),
		Description:        event.Description,
		UserID:             event.UserUuid,
		NotificationBefore: storage.TDuration(notifyBefore),
	}, nil
}

func pbEventToGRPC(event storage.Event) *Event {
	return &Event{
		Uuid:         event.UUID,
		Header:       event.Header,
		Datetime:     timestamppb.New(event.DateTime),
		Duration:     time.Duration(event.Duration).String(),
		Description:  event.Description,
		UserUuid:     event.UserID,
		NotifyBefore: time.Duration(event.NotificationBefore).String(),
	}
}

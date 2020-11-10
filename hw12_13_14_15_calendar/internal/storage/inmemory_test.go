package storage

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type EventCase struct {
	uuid     string
	datetime string
	duration TDuration
}

func TestBasicFunctionality(t *testing.T) {
	ims := NewInMemoryStorage()
	require.NotNil(t, ims)

	tt := []EventCase{
		{"UUID-1", "2020-10-10 14:15:15", TDuration(time.Hour)},
		{"UUID-2", "2020-10-10 16:00:00", TDuration(time.Hour + time.Minute*30)},
		{"UUID-3", "2020-10-10 18:00:00", TDuration(time.Minute * 25)},
	}
	ttFailedAdd := EventCase{
		"", "2020-10-10 15:15:15", TDuration(time.Hour),
	}

	// Adding events to inmemory storage
	for _, test := range tt {
		datetime, err := time.Parse("2006-01-02 15:04:05", test.datetime)
		require.Nil(t, err)

		err = ims.Add(&Event{
			UUID:               test.uuid,
			Header:             "Header",
			DateTime:           datetime,
			Duration:           test.duration,
			Description:        "very boring",
			UserID:             "User UUID-1",
			NotificationBefore: TDuration(24 * time.Hour),
		})
		require.Nil(t, err)
	}

	// Adding emtpy UUID event
	datetime, err := time.Parse("2006-01-02 15:04:05", ttFailedAdd.datetime)
	require.Nil(t, err)
	err = ims.Add(&Event{
		UUID:               ttFailedAdd.uuid,
		Header:             "Header",
		DateTime:           datetime,
		Duration:           ttFailedAdd.duration,
		Description:        "very boring",
		UserID:             "User UUID-1",
		NotificationBefore: TDuration(24 * time.Hour),
	})
	require.Error(t, err)

	// Check Events
	day, err := time.Parse("2006-01-02", "2020-10-10")
	require.Nil(t, err)
	events := ims.ListDay(day)
	require.NotNil(t, events)
	require.Equal(t, len(tt), len(events))

	// Delete and check deletion
	err = ims.Delete("UUID-2")
	require.Nil(t, err)
	events = ims.ListDay(day)
	require.Equal(t, len(tt)-1, len(events))

	// Test double deletion
	err = ims.Delete("UUID-2")
	require.NotNil(t, err)
	require.Equal(t, len(tt)-1, len(events))

	// Test empty ID event modification
	err = ims.Modify("", &events[0])
	require.Error(t, err)

	// Test non-existent event modification
	err = ims.Modify("UUID-0", &events[0])
	require.Error(t, err)

	// Delete third event and trying to modify last event
	err = ims.Delete("UUID-3")
	require.Nil(t, err)
	modifyedEvent := events[0] // UUID-1
	descriptionBefore := modifyedEvent.Description
	modifyedEvent.Description = "not so boring"
	err = ims.Modify("UUID-1", &modifyedEvent)
	require.Nil(t, err)
	events = ims.ListDay(day)
	descriptionAfter := events[0].Description
	require.NotEqual(t, descriptionBefore, descriptionAfter)
}

func TestFewDaysLists(t *testing.T) {
	ims := NewInMemoryStorage()
	require.NotNil(t, ims)

	ttday1 := []EventCase{
		{"UUID-1", "2020-10-10 15:00:00", TDuration(time.Hour)},
		{"UUID-2", "2020-10-10 16:00:00", TDuration(time.Hour)},
		{"UUID-3", "2020-10-10 18:00:00", TDuration(time.Hour)},
	}
	ttday2 := []EventCase{
		{"UUID-4", "2020-10-07 15:00:00", TDuration(time.Hour)},
	}
	ttday3 := []EventCase{
		{"UUID-5", "2020-10-12 15:00:00", TDuration(time.Hour)},
	}
	testDays := [][]EventCase{ttday1, ttday2, ttday3}
	for _, testDay := range testDays {
		for _, test := range testDay {
			datetime, err := time.Parse("2006-01-02 15:04:05", test.datetime)
			require.Nil(t, err)

			err = ims.Add(&Event{
				UUID:               test.uuid,
				Header:             "Header",
				DateTime:           datetime,
				Duration:           test.duration,
				Description:        "very boring",
				UserID:             "User UUID-1",
				NotificationBefore: TDuration(24 * time.Hour),
			})
			require.Nil(t, err)
		}
	}

	// Test ListMonth
	month, err := time.Parse("2006-01-02", "2020-10-01")
	require.Nil(t, err)
	events := ims.ListMonth(month)
	require.NotNil(t, events)
	require.Equal(t, len(ttday1)+len(ttday2)+len(ttday3), len(events))

	//Test ListWeek
	week1, err := time.Parse("2006-01-02", "2020-10-05")
	require.Nil(t, err)
	week2, err := time.Parse("2006-01-02", "2020-10-12")
	require.Nil(t, err)
	eventsWeek1 := ims.ListWeek(week1)
	require.NotNil(t, eventsWeek1)
	eventsWeek2 := ims.ListWeek(week2)
	require.NotNil(t, eventsWeek2)
	require.Equal(t, len(ttday1)+len(ttday2), len(eventsWeek1))
	require.Equal(t, len(ttday3), len(eventsWeek2))

	// Test ListWeek on empty week
	week3, err := time.Parse("2006-01-02", "2020-10-19")
	require.Nil(t, err)
	eventsWeek3 := ims.ListWeek(week3)
	require.NotNil(t, eventsWeek3)
	require.Equal(t, 0, len(eventsWeek3))
}

func TestEdgeCases(t *testing.T) {
	ims := NewInMemoryStorage()
	require.NotNil(t, ims)

	tt := []EventCase{
		{"UUID-1", "2020-10-10 00:00:00", TDuration(time.Hour)},                  // 00:00:00-01:00:00
		{"UUID-2", "2020-10-10 01:00:00", TDuration(time.Hour + time.Minute*30)}, // 01:00:00-02:30:00
		{"UUID-3", "2020-10-10 23:59:59", TDuration(time.Minute)},                // 23:59:59-00:00:59
	}
	ttFailed := []EventCase{
		{"UUID-4", "2020-10-10 00:00:00", TDuration(time.Hour)},       // Same time as UUID-1
		{"UUID-5", "2020-10-10 01:40:00", TDuration(time.Minute)},     // Event within UUID-2
		{"UUID-6", "2020-10-10 02:29:00", TDuration(time.Minute * 2)}, // Overlap last minute of UUID-3
		{"UUID-1", "2020-10-15 00:00:00", TDuration(time.Hour)},       // Existent UUID
	}
	ttFailedModify := []EventCase{
		{"UUID-1", "2020-10-10 00:01:00", TDuration(time.Hour)},           // Overlap original UUID-2 by 1 minute
		{"UUID-2", "2020-10-11 00:00:58", TDuration(time.Minute)},         // Overlap original UUID-3
		{"UUID-3", "2020-10-01 00:00:00", TDuration(time.Hour * 24 * 30)}, // Overlap all events
	}

	// Adding events to inmemory storage
	for _, test := range tt {
		datetime, err := time.Parse("2006-01-02 15:04:05", test.datetime)
		require.Nil(t, err)

		err = ims.Add(&Event{
			UUID:               test.uuid,
			Header:             "Header",
			DateTime:           datetime,
			Duration:           test.duration,
			Description:        "very boring",
			UserID:             "User UUID-1",
			NotificationBefore: TDuration(24 * time.Hour),
		})
		require.Nil(t, err)
	}

	// Check sequential events
	day, err := time.Parse("2006-01-02", "2020-10-10")
	require.Nil(t, err)
	events := ims.ListDay(day)
	require.NotNil(t, events)
	require.Equal(t, len(tt), len(events))

	// Check errors on addition overlapping events
	for _, test := range ttFailed {
		datetime, err := time.Parse("2006-01-02 15:04:05", test.datetime)
		require.Nil(t, err)

		err = ims.Add(&Event{
			UUID:               test.uuid,
			Header:             "Header",
			DateTime:           datetime,
			Duration:           test.duration,
			Description:        "very boring",
			UserID:             "User UUID-1",
			NotificationBefore: TDuration(24 * time.Hour),
		})
		require.Error(t, err)
	}

	// Check errors on modifying with new overlapping events
	for _, test := range ttFailedModify {
		datetime, err := time.Parse("2006-01-02 15:04:05", test.datetime)
		require.Nil(t, err)

		err = ims.Modify(test.uuid, &Event{
			UUID:               test.uuid,
			Header:             "Header",
			DateTime:           datetime,
			Duration:           test.duration,
			Description:        "very boring",
			UserID:             "User UUID-1",
			NotificationBefore: TDuration(24 * time.Hour),
		})
		require.Error(t, err, test.uuid)
	}
}

package storage

import (
	"errors"
	"sync"
	"time"
)

type InMemoryStorage struct {
	mu sync.RWMutex
	m  map[string]*Event
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		sync.RWMutex{},
		make(map[string]*Event),
	}
}

func (ims *InMemoryStorage) Add(event *Event) error {
	ims.mu.Lock()
	defer ims.mu.Unlock()

	if event.UUID == "" {
		return errors.New("event doesn't exists")
	}

	if ims.checkTimeBusy(event, false) {
		return errors.New("time is busy")
	}

	if _, ok := ims.m[event.UUID]; ok {
		return errors.New("ID exists")
	}

	ims.m[event.UUID] = event
	return nil
}
func (ims *InMemoryStorage) Modify(id string, event *Event) error {
	ims.mu.Lock()
	defer ims.mu.Unlock()

	if id == "" {
		return errors.New("empty ID")
	}
	if _, ok := ims.m[id]; !ok {
		return errors.New("event doesn't exists")
	}

	if ims.checkTimeBusy(event, true) {
		return errors.New("time is busy")
	}

	ims.m[id] = event
	return nil
}

func (ims *InMemoryStorage) Delete(id string) error {
	ims.mu.Lock()
	defer ims.mu.Unlock()

	if _, ok := ims.m[id]; !ok {
		return errors.New("event doesn't exists")
	}

	delete(ims.m, id)
	return nil
}

func (ims *InMemoryStorage) ListDay(day time.Time) []Event {
	ims.mu.RLock()
	defer ims.mu.RUnlock()

	result := []Event{}

	for _, event := range ims.m {
		// Reducing by nanosecond is required cause it return false at equal time
		if event.DateTime.After(day.Add(-time.Nanosecond)) && event.DateTime.Before(day.Add(time.Hour*24)) {
			result = append(result, *event)
		}
	}

	return result
}

func (ims *InMemoryStorage) ListWeek(week time.Time) []Event {
	ims.mu.RLock()
	defer ims.mu.RUnlock()

	result := []Event{}

	for _, event := range ims.m {
		if event.DateTime.After(week.Add(-time.Nanosecond)) && event.DateTime.Before(week.Add(time.Hour*24*7)) {
			result = append(result, *event)
		}
	}

	return result
}

func (ims *InMemoryStorage) ListMonth(month time.Time) []Event {
	ims.mu.RLock()
	defer ims.mu.RUnlock()

	result := []Event{}

	for _, event := range ims.m {
		if month.Year() == event.DateTime.Year() && month.Month() == event.DateTime.Month() {
			result = append(result, *event)
		}
	}

	return result
}

func (ims *InMemoryStorage) checkTimeBusy(event *Event, modify bool) bool {
	eTimeBegin := event.DateTime
	eTimeEnd := event.DateTime.Add(time.Duration(event.Duration))

	for _, e := range ims.m {
		if modify && event.UUID == e.UUID {
			continue
		}
		currEventTimeBegin := e.DateTime
		currEventTimeEnd := e.DateTime.Add(time.Duration(e.Duration))
		if eTimeBegin.After(currEventTimeEnd.Add(-time.Nanosecond)) || eTimeEnd.Before(currEventTimeBegin.Add(time.Nanosecond)) {
			continue
		} else {
			return true
		}
	}

	return false
}

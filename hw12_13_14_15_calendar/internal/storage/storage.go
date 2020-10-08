package storage

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/config"
)

type TDuration time.Duration

type User struct {
	ID   int    `db:"id"`
	UUID string `db:"uuid"`
}

type Event struct {
	// TODO: make time.Duration hidden fields,
	// 	     use string variable with psql
	ID                 int       `db:"id"`
	UUID               string    `db:"uuid"`
	Header             string    `db:"header"`
	DateTime           time.Time `db:"dt"`
	Duration           TDuration `db:"duration"`
	Description        string    `db:"description"`
	UserID             string    `db:"user_id"`
	NotificationBefore TDuration `db:"notify_before"`
}

type Storager interface {
	Add(event *Event) error
	Modify(id string, e *Event) error
	Delete(id string) error
	ListDay(day time.Time) []Event
	ListWeek(week time.Time) []Event
	ListMonth(month time.Time) []Event
}

type ErrUserDoesntExists struct{}
type ErrEventDoesntExists struct{}
type ErrTimeBusy struct{}

func GetStoragerByConfig(cfg *config.Config) Storager {
	switch cfg.Database {
	case "inmemory":
		return NewInMemoryStorage()
	case "postgresql":
		return NewPgSQLStorage(cfg)
	}
	return nil
}

func (e *ErrUserDoesntExists) Error() string {
	return "user doesn't exists"
}

func (e *ErrEventDoesntExists) Error() string {
	return "event doesn't exists"
}

func (e *ErrTimeBusy) Error() string {
	return "time is busy"
}

func (d TDuration) Value() (driver.Value, error) {
	return []byte(time.Duration(d).String()), nil
}

func (d *TDuration) Scan(src interface{}) error {
	var source string
	switch src := src.(type) {
	case string:
		source = src
	case []byte:
		source = string(src)
	default:
		return errors.New("incompatible type for TDuration")
	}

	value, err := time.ParseDuration(source)
	if err != nil {
		return err
	}

	*d = TDuration(value)
	return nil
}

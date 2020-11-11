package storage

import (
	"time"

	"github.com/Demacr/otus_golang_hw/hw12_13_14_15_calendar/internal/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Register postgres driver
	"github.com/pkg/errors"
)

type PgSQLStorage struct {
	db *sqlx.DB
}

func NewPgSQLStorageStruct(db *sqlx.DB) *PgSQLStorage {
	return &PgSQLStorage{
		db: db,
	}
}

func (pgs *PgSQLStorage) Add(event *Event) (addError error) {
	tx, err := pgs.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "error during creating transaction")
	}
	defer func() {
		if addError != nil {
			if err := tx.Rollback(); err != nil {
				addError = errors.Wrap(err, addError.Error())
			}
		}
	}()

	user := User{}
	logger.Debug("created transaction")
	err = tx.Get(&user, "SELECT * FROM users WHERE uuid=$1", event.UserID)
	if err != nil {
		return errors.Wrap(err, "error during getting user")
	}

	logger.Debug("check user.ID")
	if user.UUID == "" {
		return errors.New("user doesn't exists")
	}

	logger.Debug("insert new event")
	_, err = tx.Exec("INSERT INTO events (uuid, header, dt, duration, description, user_id, notify_before) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		event.UUID,
		event.Header,
		event.DateTime,
		event.Duration,
		event.Description,
		user.ID,
		event.NotificationBefore,
	)
	if err != nil {
		return errors.Wrap(err, "error during insert new event")
	}

	return tx.Commit()
}

func (pgs *PgSQLStorage) Modify(id string, event *Event) (modifyError error) {
	tx, err := pgs.db.Beginx()
	if err != nil {
		return errors.Wrap(err, "error during creating transaction")
	}
	defer func() {
		if modifyError != nil {
			if err := tx.Rollback(); err != nil {
				modifyError = errors.Wrap(err, modifyError.Error())
			}
		}
	}()

	user := User{}
	logger.Debug("created transaction")
	err = tx.Get(&user, "SELECT * FROM users WHERE uuid=$1", event.UserID)
	if err != nil {
		return errors.Wrap(err, "error during getting user")
	}

	logger.Debug("check user.ID")
	if user.UUID == "" {
		return errors.New("user doesn't exists")
	}

	_, err = tx.Exec("UPDATE events SET header = $2, dt = $3, duration = $4, description = $5, user_id = $6, notify_before = $7 WHERE uuid = $1",
		event.UUID,
		event.Header,
		event.DateTime,
		event.Duration,
		event.Description,
		user.ID,
		event.NotificationBefore,
	)
	if err != nil {
		return errors.Wrap(err, "error during update event")
	}

	return tx.Commit()
}

func (pgs *PgSQLStorage) Delete(id string) error {
	_, err := pgs.db.Exec("DELETE FROM events WHERE uuid = $1", id)
	if err != nil {
		return errors.Wrap(err, "error during deletion event")
	}
	return nil
}

func (pgs *PgSQLStorage) ListDay(day time.Time) []Event {
	return pgs.listCommon(day, day.Add(24*time.Hour))
}

func (pgs *PgSQLStorage) ListWeek(week time.Time) []Event {
	return pgs.listCommon(week, week.Add(7*24*time.Hour))
}

func (pgs *PgSQLStorage) ListMonth(month time.Time) []Event {
	return pgs.listCommon(month, month.AddDate(0, 1, 0))
}

func (pgs *PgSQLStorage) listCommon(t1, t2 time.Time) []Event {
	result := []Event{}

	rows, err := pgs.db.Queryx("SELECT id, uuid, header, dt, TO_CHAR(duration, 'HH24hMImSSs') as duration, description, user_id, TO_CHAR(notify_before, 'HH24hMImSSs') as notify_before FROM events WHERE dt >= $1 AND dt < $2;",
		t1,
		t2,
	)
	if err != nil {
		logger.Error(err)
		return nil
	}
	defer rows.Close()

	event := Event{}
	for rows.Next() {
		err = rows.StructScan(&event)
		if err != nil {
			logger.Error(err)
			return nil
		}

		// TODO: check copying object
		result = append(result, event)
	}

	if rows.Err() != nil {
		logger.Error(err)
		return nil
	}
	return result
}

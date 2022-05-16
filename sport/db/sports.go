package db

import (
	"database/sql"
	"strings"
	"sync"
	"time"

	"github.com/golang/protobuf/ptypes"
	_ "github.com/mattn/go-sqlite3"

	"git.neds.sh/matty/entain/sport/proto/sport"
)

// SportRepo provides repository access to events.
type SportRepo interface {
	// Init will initialise our sport repository.
	Init() error

	// List will return a list of events.
	List(filter *sport.ListEventsRequestFilter) ([]*sport.Event, error)
}

type sportRepo struct {
	db   *sql.DB
	init sync.Once
}

// NewSportRepo creates a new sport repository.
func NewSportRepo(db *sql.DB) *sportRepo {
	return &sportRepo{db: db}
}

// Init prepares the sport repository dummy data.
func (r *sportRepo) Init() error {
	var err error

	r.init.Do(func() {
		// For test/example purposes, we seed the DB with some dummy sports.
		err = r.seed()
	})

	return err
}

func (r *sportRepo) List(filter *sport.ListEventsRequestFilter) ([]*sport.Event, error) {
	var (
		err   error
		query string
		args  []interface{}
	)

	query = getSportQueries()[eventsList]

	query, args = r.applyFilter(query, filter)

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return r.scanEvents(rows)
}

func (r *sportRepo) applyFilter(query string, filter *sport.ListEventsRequestFilter) (string, []interface{}) {
	var (
		clauses []string
		args    []interface{}
	)

	if filter == nil {
		return query, args
	}

	if len(filter.Ids) > 0 {
		clauses = append(clauses, "id IN ("+strings.Repeat("?,", len(filter.Ids)-1)+"?)")

		for _, ID := range filter.Ids {
			args = append(args, ID)
		}
	}

	if len(clauses) != 0 {
		query += " WHERE " + strings.Join(clauses, " AND ")
	}
	return query, args
}

func (m *sportRepo) scanEvents(
	rows *sql.Rows,
) ([]*sport.Event, error) {
	var events []*sport.Event

	for rows.Next() {
		var event sport.Event
		var (
			advertisedStart time.Time
			eventStart      time.Time
			eventEnd        time.Time
		)

		if err := rows.Scan(&event.Id, &event.Name, &advertisedStart, &eventStart, &eventEnd); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}

			return nil, err
		}

		ts, err := ptypes.TimestampProto(advertisedStart)
		if err != nil {
			return nil, err
		}

		event.AdvertisedStartTime = ts

		events = append(events, &event)
	}

	return events, nil
}

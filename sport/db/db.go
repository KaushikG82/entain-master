package db

import (
	"strconv"
	"time"

	"syreclabs.com/go/faker"
)

func (r *sportRepo) seed() error {
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS events (id INTEGER PRIMARY KEY, name TEXT, advertised_start_time DATETIME, event_start_time DATETIME, event_end_time DATETIME)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO events(id, name, advertised_start_time, event_start_time, event_end_time) VALUES (?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Team().Name(),
				faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 8)).Format(time.RFC3339),
				faker.Time().Between(time.Now().AddDate(0, 0, 10), time.Now().AddDate(0, 0, 20)).Format(time.RFC3339),
				faker.Time().Between(time.Now().AddDate(0, 0, 21), time.Now().AddDate(0, 1, 2)).Format(time.RFC3339),
			)
		}
	}

	statement, err = r.db.Prepare(`CREATE TABLE IF NOT EXISTS events_races_maspping (FOREIGN KEY (event_id) REFERENCES events(id), meeting_id INTEGER)`)
	if err == nil {
		_, err = statement.Exec()
	}

	x := 1
	for i := 1; i <= 100 && x <= 100; i++ {
		max, _ := strconv.Atoi(faker.Number().Between(1, 3))
		for j := 1; j <= max && x <= 100; j++ {
			statement, err = r.db.Prepare(`INSERT OR IGNORE INTO events_races_maspping(event_id, meeting_id) VALUES (?,?)`)
			if err == nil {
				_, err = statement.Exec(
					i,
					x+(j-1),
				)
				x++
			}
		}

	}
	return err
}

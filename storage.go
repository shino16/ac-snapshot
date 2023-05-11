package main

import (
	"database/sql"
	"time"
)

type Storage struct {
	db *sql.DB
}

type History []Snapshot

func initStorage(db *sql.DB) Storage {
	const sql = `
		CREATE TABLE IF NOT EXISTS snapshots (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			rating INTEGER,
			ac_count INTEGER,
			rps INTEGER,
			time TEXT
		);
		CREATE INDEX IF NOT EXISTS snapshots_name_idx ON snapshots (name);
    `
	if _, err := db.Exec(sql); err != nil {
		panic(err.Error())
	}
	return Storage{db}
}

func (s *Storage) getHistory(name string) (History, error) {
	sql := `
		SELECT name, rating, ac_count, rps, time
		FROM snapshots
		WHERE name = ?
		ORDER BY time DESC;`
	rows, err := s.db.Query(sql, name)
	if err != nil {
		return nil, err
	}
	ret := make([]Snapshot, 0)
	for rows.Next() {
		shot := Snapshot{}
		var timestamp string
		if err := rows.Scan(
			&shot.Name, &shot.Rating, &shot.ACCount, &shot.RPS, &timestamp,
		); err != nil {
			return nil, err
		}
		if shot.Time, err = time.Parse(time.RFC3339, timestamp); err != nil {
			return nil, err
		}
		ret = append(ret, shot)
	}
	return ret, nil
}

func (s *Storage) addSnapshot(shot Snapshot) error {
	sql := `
		INSERT INTO snapshots (name, rating, ac_count, rps, time)
		VALUES (?, ?, ?, ?, ?);`
	timestamp := shot.Time.Format(time.RFC3339)
	_, err := s.db.Exec(
		sql, shot.Name, shot.Rating, shot.ACCount, shot.RPS, timestamp)
	return err
}

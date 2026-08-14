package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() error {
	const query = `
	INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	defer stmt.Close()
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	e.ID = id
	return nil
}

func GetAllEvents() []Event {
	return events
}

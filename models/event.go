package models

import (
	"time"

	"github.com/rest-api-event/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required" json:"name"`
	Description string    `binding:"required" json:"description"`
	StartTime   time.Time `binding:"required" json:"start_time"`
	EndTime     time.Time `binding:"required" json:"end_time"`
	Location    string    `binding:"required" json:"location"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int64     `json:"user_id"`
}

var events = []Event{}

func (e *Event) Save() error {
	query := `
    INSERT INTO events (name, description, location, start_time, end_time, user_id)
    VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	result, err := statement.Exec(e.Name, e.Description, e.Location, e.StartTime, e.EndTime, e.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
    SELECT id, name, description, location, start_time, end_time, created_at, updated_at, user_id 
    FROM events`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location,
			&event.StartTime, &event.EndTime, &event.CreatedAt,
			&event.UpdatedAt, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	// Don't use SELECT * - specify columns in the correct order to match the Scan() call
	query := `SELECT id, name, description, location, start_time, end_time, created_at, updated_at, user_id 
              FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location,
		&event.StartTime, &event.EndTime, &event.CreatedAt, &event.UpdatedAt, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, start_time = ?, end_time = ?, user_id = ?
	WHERE id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(event.Name, event.Description, event.Location,
		event.StartTime, event.EndTime, event.UserID, event.ID)
	return err
}

func (event Event) Delete() error {
	query := `DELETE FROM events WHERE id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()

	_, err = statement.Exec(event.ID)
	return err
}

func (e Event) Register(userId int64) error {
	query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?, ?)`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := `
	DELETE FROM registrations
	WHERE event_id = ? AND user_id = ?`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer statement.Close()
	_, err = statement.Exec(e.ID, userId)
	return err
}

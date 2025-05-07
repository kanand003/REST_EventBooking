package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date"`
	CreatedAt   time.Time `json:"created_at"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() {
	// later add to database
	events = append(events, e)
}

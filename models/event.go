package models


import (
	"time"
)

type Event struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Location string `json:"location"`
	UserID int `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

var events = []Event{}

func (e *Event) Save() error {
	events = append(events, *e)
	return nil
}

func GetAllEvents() []Event{
	return events
}
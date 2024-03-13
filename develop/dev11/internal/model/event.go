package model

import "time"

type CalendarEvent struct {
	UserID int       `json:"user_id"`
	Date   time.Time `json:"date"`
	Title  string    `json:"title"`
}
package service

import (
	"time"

	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/model"
)

type Calendar interface {
	CreateAnEvent(event model.CalendarEvent) error
	UpdateAnEvent(event model.CalendarEvent) error
	DeleteAnEvent(event model.CalendarEvent) error
	EventsForDay(date time.Time) (events []model.CalendarEvent, err error)
	EventsForWeek(date time.Time) (events []model.CalendarEvent, err error)
	EventsForMonth(date time.Time) (events []model.CalendarEvent, err error)
}

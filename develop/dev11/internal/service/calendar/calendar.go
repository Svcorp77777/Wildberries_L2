package service

import (
	"fmt"
	"time"

	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/model"
	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/service"
)

type calendar struct{}

var simulationDB = make(map[time.Time][]model.CalendarEvent)

func (s *calendar) CreateAnEvent(event model.CalendarEvent) error {
	records, exist := simulationDB[event.Date]

	if exist {
		for i := 0; i < len(records); i++ {
			if records[i].UserID == event.UserID {
				return fmt.Errorf("такая запись уже существует")
			}
		}
	}

	simulationDB[event.Date] = append(simulationDB[event.Date], event)

	return nil
}

func (s *calendar) UpdateAnEvent(event model.CalendarEvent) error {
	records, exist := simulationDB[event.Date]

	if exist {
		for i := 0; i < len(records); i++ {
			if records[i].UserID == event.UserID {
				records[i].Title = event.Title
			}
		}
	}

	return fmt.Errorf("такое событие не было найдено")
}

func (s *calendar) DeleteAnEvent(event model.CalendarEvent) error {
	records, exist := simulationDB[event.Date]

	if exist {
		for i := 0; i < len(records); i++ {
			if records[i].UserID == event.UserID {
				records[i] = records[len(records)-1]
				records = records[:len(records)-1]
			}
		}
	}

	return fmt.Errorf("такое событие не было найдено")
}

func (s *calendar) EventsForDay(date time.Time) (events []model.CalendarEvent, err error) {
	events, exist := simulationDB[date]

	if exist {
		return events, nil
	}

	return events, fmt.Errorf("такое событие не было найдено")
}

func (s *calendar) EventsForWeek(date time.Time) (events []model.CalendarEvent, err error) {
	beginningWeek := date.AddDate(0, 0, -int(date.Weekday())+1)
	endingWeek := beginningWeek.AddDate(0, 0, 6)

	for currentDay := beginningWeek; currentDay.Before(endingWeek.AddDate(0, 0, 1)); currentDay = currentDay.AddDate(0, 0, 1) {
		events = append(events, simulationDB[currentDay]...)
	}

	if events == nil {
		return events, fmt.Errorf("не было найдено событий")
	}

	return events, nil
}

func (s *calendar) EventsForMonth(date time.Time) (events []model.CalendarEvent, err error) {
	beginningMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, time.UTC)
	endingMonth := beginningMonth.AddDate(0, 1, -1)

	for currentDay := beginningMonth; currentDay.Before(endingMonth.AddDate(0, 0, 1)); currentDay = currentDay.AddDate(0, 0, 1) {
		events = append(events, simulationDB[currentDay]...)
	}

	if events == nil {
		return events, fmt.Errorf("не было найдено событий")
	}

	return events, nil
}

func ServiceCalendar() service.Calendar {
	return &calendar{}
}

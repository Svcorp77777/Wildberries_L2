package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/model"
	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/service"
	"github.com/Svcorp77777/Wildberries-L2/dev11/internal/transport"
)

type handlerСalendar struct {
	serviceCalendar service.Calendar
}

func (h *handlerСalendar) Init(router *http.ServeMux) {
	// TODO? --- GET запросы
	router.HandleFunc("/events_for_day", h.eventsForDay)
	router.HandleFunc("/events_for_week", h.eventsForWeek)
	router.HandleFunc("/events_for_month", h.eventsForMonth)

	// TODO? --- POST запросы
	router.HandleFunc("/create_event", h.createEvent)
	router.HandleFunc("/update_event", h.updateEvent)
	router.HandleFunc("/delete_event", h.deleteEvent)
}

// TODO? -- GET
func (h *handlerСalendar) eventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		answerIncorrect := AnswerIncorrect{
			Error: "Недопустимый метод запроса",
		}

		JsonResponse(w, http.StatusMethodNotAllowed, answerIncorrect)

		return
	}

	date, err := checkRequestData(r)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при проверке данных: " + err.Error(),
		}

		JsonResponse(w, http.StatusBadRequest, answerIncorrect)

		return
	}

	events, err := h.serviceCalendar.EventsForDay(date)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Ошибка при поиске событий за день: " + err.Error(),
		}

		JsonResponse(w, http.StatusServiceUnavailable, answerIncorrect)

		return
	}

	answerСorrect := AnswerСorrect{
		Result: events,
	}

	JsonResponse(w, http.StatusOK, answerСorrect)
}

// TODO? -- GET
func (h *handlerСalendar) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		answerIncorrect := AnswerIncorrect{
			Error: "Недопустимый метод запроса",
		}

		JsonResponse(w, http.StatusMethodNotAllowed, answerIncorrect)

		return
	}

	date, err := checkRequestData(r)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при проверке данных: " + err.Error(),
		}

		JsonResponse(w, http.StatusBadRequest, answerIncorrect)

		return
	}

	events, err := h.serviceCalendar.EventsForWeek(date)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Ошибка при поиске событий за неделю: " + err.Error(),
		}

		JsonResponse(w, http.StatusServiceUnavailable, answerIncorrect)

		return
	}

	answerСorrect := AnswerСorrect{
		Result: events,
	}

	JsonResponse(w, http.StatusOK, answerСorrect)
}

// TODO? -- GET
func (h *handlerСalendar) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		answerIncorrect := AnswerIncorrect{
			Error: "Недопустимый метод запроса",
		}

		JsonResponse(w, http.StatusMethodNotAllowed, answerIncorrect)

		return
	}

	date, err := checkRequestData(r)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при проверке данных: " + err.Error(),
		}

		JsonResponse(w, http.StatusBadRequest, answerIncorrect)

		return
	}

	events, err := h.serviceCalendar.EventsForMonth(date)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Ошибка при поиске событий за месяц: " + err.Error(),
		}

		JsonResponse(w, http.StatusServiceUnavailable, answerIncorrect)

		return
	}

	answerСorrect := AnswerСorrect{
		Result: events,
	}

	JsonResponse(w, http.StatusOK, answerСorrect)
}

// TODO! -- POST
func (h *handlerСalendar) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		answerIncorrect := AnswerIncorrect{
			Error: "Недопустимый метод запроса",
		}

		JsonResponse(w, http.StatusMethodNotAllowed, answerIncorrect)

		return
	}

	event, err := checkAnalyzParameters(r)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при проверке данных: " + err.Error(),
		}

		JsonResponse(w, http.StatusBadRequest, answerIncorrect)

		return
	}

	if err := h.serviceCalendar.CreateAnEvent(event); err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при создании события: " + err.Error(),
		}

		JsonResponse(w, http.StatusServiceUnavailable, answerIncorrect)

		return
	}

	answerСorrect := AnswerСorrect{
		Result: "Success!",
	}

	JsonResponse(w, http.StatusOK, answerСorrect)
}

// TODO! -- POST
func (h *handlerСalendar) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		answerIncorrect := AnswerIncorrect{
			Error: "Недопустимый метод запроса",
		}

		JsonResponse(w, http.StatusMethodNotAllowed, answerIncorrect)

		return
	}

	event, err := checkAnalyzParameters(r)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при проверке данных: " + err.Error(),
		}

		JsonResponse(w, http.StatusBadRequest, answerIncorrect)

		return
	}

	if err := h.serviceCalendar.UpdateAnEvent(event); err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при обновлении события: " + err.Error(),
		}

		JsonResponse(w, http.StatusServiceUnavailable, answerIncorrect)

		return
	}

	answerСorrect := AnswerСorrect{
		Result: "Success!",
	}

	JsonResponse(w, http.StatusOK, answerСorrect)
}

// TODO! -- POST
func (h *handlerСalendar) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		answerIncorrect := AnswerIncorrect{
			Error: "Недопустимый метод запроса",
		}

		JsonResponse(w, http.StatusMethodNotAllowed, answerIncorrect)

		return
	}

	event, err := checkAnalyzParameters(r)
	if err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Возникла ошибка при проверке данных: " + err.Error(),
		}

		JsonResponse(w, http.StatusBadRequest, answerIncorrect)

		return
	}

	if err := h.serviceCalendar.DeleteAnEvent(event); err != nil {
		answerIncorrect := AnswerIncorrect{
			Error: "Ошибка при удалении события: " + err.Error(),
		}

		JsonResponse(w, http.StatusServiceUnavailable, answerIncorrect)

		return
	}

	answerСorrect := AnswerСorrect{
		Result: "Success!",
	}

	JsonResponse(w, http.StatusOK, answerСorrect)
}

func checkAnalyzParameters(r *http.Request) (event model.CalendarEvent, err error) {
	userID := r.FormValue("user_id")
	dateStr := r.FormValue("date")
	title := r.FormValue("title")

	path := r.URL.Path

	if path == "delete_event" {
		if userID == "" || dateStr == "" {
			return event, fmt.Errorf("неверно переданы параметры запроса")
		}
	} else {
		if userID == "" || dateStr == "" || title == "" {
			return event, fmt.Errorf("неверно переданы параметры запроса")
		}
	}

	userIDInt, err := strconv.Atoi(userID)
	if err != nil {
		return event, fmt.Errorf("неверный user_id")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return event, fmt.Errorf("неверный формат даты")
	}

	return model.CalendarEvent{
		UserID: userIDInt,
		Date:   date,
		Title:  title,
	}, nil
}

func checkRequestData(r *http.Request) (time.Time, error) {
	dateStr := r.URL.Query().Get("date")

	if dateStr == "" {
		return time.Time{}, fmt.Errorf("неверно переданы параметры запроса")
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("неверный формат даты")
	}

	return date, nil
}

func HandlerCalendar(serviceCalendar service.Calendar) transport.Handler {
	return &handlerСalendar{
		serviceCalendar: serviceCalendar,
	}
}

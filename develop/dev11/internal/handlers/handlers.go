package handlers

import (
	"dev11/internal/event"
	"dev11/internal/middleware"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// POST /create_event
// POST /update_event
// POST /delete_event
// GET /events_for_day
// GET /events_for_week
// GET /events_for_month

type EventHandler struct {
	ec *event.EventsCalendar
}

func (eh *EventHandler) HandleHTTPRequests() {
	http.HandleFunc("/create_event", middleware.Middleware(http.HandlerFunc(eh.createEvent)))
	http.HandleFunc("/update_event", middleware.Middleware(http.HandlerFunc(eh.updateEvent)))
	http.HandleFunc("/delete_event", middleware.Middleware(http.HandlerFunc(eh.deleteEvent)))
	http.HandleFunc("/events_for_day", middleware.Middleware(http.HandlerFunc(eh.getEventsForDay)))
	http.HandleFunc("/events_for_week", middleware.Middleware(http.HandlerFunc(eh.getEventsForWeek)))
	http.HandleFunc("/events_for_month", middleware.Middleware(http.HandlerFunc(eh.getEventsForMonth)))
}

func (eh *EventHandler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		result, err := eh.ec.AddEvent(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		response := map[string]string{"result": result}
		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, `{"error": "Wrong method"}`, http.StatusBadRequest)
}

func (eh *EventHandler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		result, err := eh.ec.UpdateEvent(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		response := map[string]string{"result": result}
		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, `{"error": "Wrong method"}`, http.StatusBadRequest)
}

func (eh *EventHandler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		result, err := eh.ec.DeleteEvent(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		response := map[string]string{"result": result}
		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, `{"error": "Wrong method"}`, http.StatusBadRequest)
}

func (eh *EventHandler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var (
			uid  uint
			date time.Time
		)
		val1, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")
		date, err2 := time.Parse("2006-01-02", dateStr)
		if err1 != nil || err2 != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		uid = uint(val1)
		eh.ec.GetDayEvents(uid, date)

	}
}
func (eh *EventHandler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var (
			uid  uint
			date time.Time
		)
		val1, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")
		date, err2 := time.Parse("2006-01-02", dateStr)
		if err1 != nil || err2 != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		uid = uint(val1)
		// eh.ec.GetDayEvents(uid, date) сделать цикл?

	}
}
func (eh *EventHandler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var (
			uid  uint
			date time.Time
		)
		val1, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")
		date, err2 := time.Parse("2006-01-02", dateStr)
		if err1 != nil || err2 != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		uid = uint(val1)
		// eh.ec.GetDayEvents(uid, date) сделать цикл?

	}
}

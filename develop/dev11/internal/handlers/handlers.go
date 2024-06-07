package handlers

import (
	"dev11/internal/event"
	"dev11/internal/middleware"
	"fmt"
	"net/http"
)

// POST /create_event
// POST /update_event
// POST /delete_event
// GET /events_for_day
// GET /events_for_week
// GET /events_for_month

func HandleHTTPRequests() {
	http.HandleFunc("/create_event", middleware.Middleware(http.HandlerFunc(createEvent)))
	http.HandleFunc("/update_event", middleware.Middleware(http.HandlerFunc(updateEvent)))
	http.HandleFunc("/delete_event", middleware.Middleware(http.HandlerFunc(deleteEvent)))
	http.HandleFunc("/events_for_day", middleware.Middleware(http.HandlerFunc(getEventsForDay)))
	http.HandleFunc("/events_for_week", middleware.Middleware(http.HandlerFunc(getEventsForWeek)))
	http.HandleFunc("/events_for_month", middleware.Middleware(http.HandlerFunc(getEventsForMonth)))
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		ev := &event.Event{}
		err := ev.Parse(r.Body)
		if err != nil{
			http.Error(w, `{"error": "Failed to decode request data"}`, 400)
		}
		fmt.Println(ev)
	}
}
func updateEvent(w http.ResponseWriter, r *http.Request)       {}
func deleteEvent(w http.ResponseWriter, r *http.Request)       {}
func getEventsForDay(w http.ResponseWriter, r *http.Request)   {}
func getEventsForWeek(w http.ResponseWriter, r *http.Request)  {}
func getEventsForMonth(w http.ResponseWriter, r *http.Request) {}

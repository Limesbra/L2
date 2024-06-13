package handlers

import (
	"dev11/internal/event"
	"dev11/internal/middleware"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

// EventHandler - структура, представляющая обработчик событий.
type EventHandler struct {
	Ec *event.EventsCalendar
}

// HandleHTTPRequests - метод структуры EventHandler, обрабатывающий HTTP запросы.
func (eh *EventHandler) HandleHTTPRequests() {
	http.HandleFunc("/create_event", middleware.Middleware(http.HandlerFunc(eh.createEvent)))
	http.HandleFunc("/update_event", middleware.Middleware(http.HandlerFunc(eh.updateEvent)))
	http.HandleFunc("/delete_event", middleware.Middleware(http.HandlerFunc(eh.deleteEvent)))
	http.HandleFunc("/events_for_day", middleware.Middleware(http.HandlerFunc(eh.getEventsForDay)))
	http.HandleFunc("/events_for_week", middleware.Middleware(http.HandlerFunc(eh.getEventsForWeek)))
	http.HandleFunc("/events_for_month", middleware.Middleware(http.HandlerFunc(eh.getEventsForMonth)))
}

// createEvent - метод типа POST структуры EventHandler для создания события.
func (eh *EventHandler) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		//добавляем ивент вызывая метод структуры EventsCalendar
		result, err := eh.Ec.AddEvent(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		// после успешного добавления возвращаем JSON с информацией о результате запроса
		response := map[string]string{"result": result}
		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, `{"error": "Wrong method"}`, http.StatusBadRequest)
}

// updateEvent - метод типа POST структуры EventHandler для изменения события
func (eh *EventHandler) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		//изменяем ивент вызывая метод структуры EventsCalendar
		result, err := eh.Ec.UpdateEvent(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		// после успешного изменения возвращаем JSON с информацией о результате запроса
		response := map[string]string{"result": result}
		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, `{"error": "Wrong method"}`, http.StatusBadRequest)
}

// deleteEvent - метод типа POST структуры EventHandler для удаления события
func (eh *EventHandler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == "Post" {
		//удаляем ивент вызывая метод структуры EventsCalendar
		result, err := eh.Ec.DeleteEvent(r.Body)
		if err != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		// после успешного удаления возвращаем JSON с информацией о результате запроса
		response := map[string]string{"result": result}
		json.NewEncoder(w).Encode(response)
		return
	}
	http.Error(w, `{"error": "Wrong method"}`, http.StatusBadRequest)
}

// getEventsForDay - метод типа GET структуры EventHandler для получения информации о события за день
func (eh *EventHandler) getEventsForDay(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var (
			uid  uint
			date time.Time
		)
		// получаем информацию о запросе и конвертируем данные в нужный тип
		val1, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")
		date, err2 := time.Parse("2006-01-02T15:04:05", dateStr)
		if err1 != nil || err2 != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		uid = uint(val1)
		date, _ = time.Parse("2006-01-02", (date.Format("2006-01-02")))
		// получаем все события за указанный период
		result, _ := eh.Ec.GetDayEvents(uid, date)
		// если длина слайса равна нулю => событий в этот период нет
		if len(result) == 0 {
			response := map[string]string{"result": "No events"}
			json.NewEncoder(w).Encode(response)
			return
		}
		// возвращаем JSON с информацией о результате запроса
		response := map[string][]event.Event{"result": result}
		fmt.Println(response)
		json.NewEncoder(w).Encode(response)
	}
}

func (eh *EventHandler) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var (
			uid  uint
			date time.Time
		)
		// получаем информацию о запросе и конвертируем данные в нужный тип
		val1, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")
		date, err2 := time.Parse("2006-01-02T15:04:05", dateStr)
		if err1 != nil || err2 != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		uid = uint(val1)
		date, _ = time.Parse("2006-01-02", (date.Format("2006-01-02")))
		// получаем все события за указанный период
		result, _ := eh.Ec.GetWeekEvents(uid, date)
		// если длина слайса равна нулю => событий в этот период нет
		if len(result) == 0 {
			response := map[string]string{"result": "No events"}
			json.NewEncoder(w).Encode(response)
			return
		}
		// возвращаем JSON с информацией о результате запроса
		response := map[string][]event.Event{"result": result}
		json.NewEncoder(w).Encode(response)
	}
}

func (eh *EventHandler) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var (
			uid  uint
			date time.Time
		)
		// получаем информацию о запросе и конвертируем данные в нужный тип
		val1, err1 := strconv.Atoi(r.URL.Query().Get("user_id"))
		dateStr := r.URL.Query().Get("date")
		date, err2 := time.Parse("2006-01-02T15:04:05", dateStr)
		if err1 != nil || err2 != nil {
			http.Error(w, `{"error": "Failed to decode request data"}`, http.StatusBadRequest)
			return
		}

		uid = uint(val1)
		date, _ = time.Parse("2006-01-02", (date.Format("2006-01-02")))
		// получаем все события за указанный период
		result, _ := eh.Ec.GetMonthEvents(uid, date)
		// если длина слайса равна нулю => событий в этот период нет
		if len(result) == 0 {
			response := map[string]string{"result": "No events"}
			json.NewEncoder(w).Encode(response)
			return
		}
		// возвращаем JSON с информацией о результате запроса
		response := map[string][]event.Event{"result": result}
		json.NewEncoder(w).Encode(response)
	}
}

// curl -H 'Accept: application/json' 'http://localhost:8080/events_for_day?user_id=3&date=2019-09-09'
// curl -H 'Accept: application/json' 'http://example.com/events_for_day?user_id=3&date=2019-09-09'

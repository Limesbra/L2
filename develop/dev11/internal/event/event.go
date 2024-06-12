package event

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"
)

// Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).

type EventsCalendar struct {
	mu     *sync.Mutex
	events map[uint][]event
}

type event struct {
	UserID   uint      `json:"user_ID"`
	EventID  uint      `json:"event_ID"`
	Date     time.Time `json:"Date"`
	Overview string    `json:"Overview"`
}

func CreateCalendar() EventsCalendar {
	return EventsCalendar{mu: &sync.Mutex{}, events: make(map[uint][]event)}
}

func (es *EventsCalendar) AddEvent(r io.ReadCloser) (string, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	var e event

	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bodyBytes, &e)
	if err != nil {
		return "", err
	}

	e.EventID = uint(len(es.events[e.UserID]) + 1)
	es.events[e.UserID] = append(es.events[e.UserID], e)
	status := fmt.Sprintf("event %d created", e.EventID)

	return status, nil
}

func (es *EventsCalendar) UpdateEvent(r io.ReadCloser) (string, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	var e event

	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bodyBytes, &e)
	if err != nil {
		return "", err
	}

	for i, item := range es.events[e.UserID] {
		if item.EventID == e.EventID {
			es.events[e.UserID][i].Date = e.Date
			es.events[e.UserID][i].Overview = e.Overview
			break
		}
	}

	status := fmt.Sprintf("event %d updated", e.EventID)

	return status, nil
}

func (es *EventsCalendar) DeleteEvent(r io.ReadCloser) (string, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	type dEvent struct {
		UserID  uint `json:"user_ID"`
		EventID uint `json:"event_ID"`
	}

	var dev dEvent
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	err = json.Unmarshal(bodyBytes, &dev)
	if err != nil {
		return "", err
	}

	idx := -1
	for i, item := range es.events[dev.UserID] {
		if item.EventID == dev.EventID {
			idx = i
			break
		}
	}

	if idx != -1 {
		es.events[dev.UserID] = append(es.events[dev.UserID][:idx], es.events[dev.UserID][idx+1:]...)
	}

	status := fmt.Sprintf("event %d deleted", dev.EventID)

	return status, nil
}

func (es *EventsCalendar) GetDayEvents(uid uint, date ) (string, error) {
	return "", nil
}

func (es *EventsCalendar) GetWeekEvents(r io.ReadCloser) (string, error) {
	return "", nil
}

func (es *EventsCalendar) GetMonthEvents(r io.ReadCloser) (string, error) {
	return "", nil
}

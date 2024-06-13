package event

import (
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"
)

// EventsCalendar - структура, представляющая календарь событий.
type EventsCalendar struct {
	mu     *sync.Mutex
	events map[uint][]Event
}

// Event - структура, представляющая событие.
type Event struct {
	UserID   uint      `json:"user_ID"`
	EventID  uint      `json:"event_ID"`
	Date     time.Time `json:"Date"`
	Overview string    `json:"Overview"`
}

// конструктор для структуры EventsCalendar
func CreateCalendar() EventsCalendar {
	return EventsCalendar{mu: &sync.Mutex{}, events: make(map[uint][]Event)}
}

// метод структуры EventsCalendar для добавление событий
func (es *EventsCalendar) AddEvent(r io.ReadCloser) (string, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	var e Event

	//преобразуем тело запроса в слайс байтов
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	// парсим массив байтов и записываем в структура ивента
	err = json.Unmarshal(bodyBytes, &e)
	if err != nil {
		return "", err
	}

	//присваиваем номер событию
	e.EventID = uint(len(es.events[e.UserID]) + 1)
	es.events[e.UserID] = append(es.events[e.UserID], e)

	//возвращаем информацию о выполненой операции
	status := fmt.Sprintf("event %d created", e.EventID)

	return status, nil
}

// метод структуры EventsCalendar для изменения событий
func (es *EventsCalendar) UpdateEvent(r io.ReadCloser) (string, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	var e Event

	//преобразуем тело запроса в слайс байтов
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	// парсим массив байтов и записываем в структура ивента
	err = json.Unmarshal(bodyBytes, &e)
	if err != nil {
		return "", err
	}

	// находим нужное событие и заменяем дату и описание события
	for i, item := range es.events[e.UserID] {
		if item.EventID == e.EventID {
			es.events[e.UserID][i].Date = e.Date
			es.events[e.UserID][i].Overview = e.Overview
			break
		}
	}

	//возвращаем информацию о выполненой операции
	status := fmt.Sprintf("event %d updated", e.EventID)

	return status, nil
}

// метод структуры EventsCalendar для удаления событий
func (es *EventsCalendar) DeleteEvent(r io.ReadCloser) (string, error) {
	es.mu.Lock()
	defer es.mu.Unlock()

	// временная структура для поиска нужной
	type dEvent struct {
		UserID  uint `json:"user_ID"`
		EventID uint `json:"event_ID"`
	}

	var dev dEvent

	//преобразуем тело запроса в слайс байтов
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}

	// парсим массив байтов и записываем в структура ивента
	err = json.Unmarshal(bodyBytes, &dev)
	if err != nil {
		return "", err
	}

	// выполняем поиск нужного события
	idx := -1
	for i, item := range es.events[dev.UserID] {
		if item.EventID == dev.EventID {
			idx = i
			break
		}
	}

	// возвращаем результат функции
	status := "event not found"
	if idx != -1 {
		es.events[dev.UserID] = append(es.events[dev.UserID][:idx], es.events[dev.UserID][idx+1:]...)
		status = fmt.Sprintf("event %d deleted", dev.EventID)
	}

	return status, nil
}

// метод структуры EventsCalendar для получения информации о событиях за день
func (es *EventsCalendar) GetDayEvents(uid uint, date time.Time) ([]Event, error) {
	es.mu.Lock()
	defer es.mu.Unlock()
	weekEvents := make([]Event, 0)
	// выполняем поиск подходящих событий
	for _, i := range es.events[uid] {
		short := i.Date.Truncate(24 * time.Hour)
		if short.Equal(date) {
			weekEvents = append(weekEvents, i)
		}
	}
	return weekEvents, nil
}

// метод структуры EventsCalendar для получения информации о предстаящих событиях
// период - неделя
func (es *EventsCalendar) GetWeekEvents(uid uint, date time.Time) ([]Event, error) {
	es.mu.Lock()
	defer es.mu.Unlock()
	stopDate := date.AddDate(0, 0, 8)
	weekEvents := make([]Event, 0)
	// выполняем поиск подходящих событий
	for date != stopDate {
		for _, i := range es.events[uid] {
			short := i.Date.Truncate(24 * time.Hour)
			if short.Equal(date) {
				weekEvents = append(weekEvents, i)
			}
		}
		date = date.AddDate(0, 0, 1)
	}
	return weekEvents, nil
}

// метод структуры EventsCalendar для получения информации о предстаящих событиях
// период - месяц
func (es *EventsCalendar) GetMonthEvents(uid uint, date time.Time) ([]Event, error) {
	es.mu.Lock()
	defer es.mu.Unlock()
	stopDate := date.AddDate(0, 0, 31)
	weekEvents := make([]Event, 0)
	// выполняем поиск подходящих событий
	for date != stopDate {
		for _, i := range es.events[uid] {
			short := i.Date.Truncate(24 * time.Hour)
			if short.Equal(date) {
				weekEvents = append(weekEvents, i)
			}
		}
		date = date.AddDate(0, 0, 1)
	}
	return weekEvents, nil
}

package event

import (
	"encoding/json"
	"io"
	"time"
)

// Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).

type Event struct {
	ID       uint      `json:"ID"`
	Date     time.Time `json:"Date"`
	Overview string    `json:"Overview"`
}

func (e *Event) Parse(r io.ReadCloser) error {
	bodyBytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bodyBytes, &e)
	if err != nil {
		return err
	}
	return nil
}

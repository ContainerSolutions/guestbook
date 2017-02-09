package entries

import (
	"errors"
	"time"
)

type Entry struct {
	Name    string
	Message string
	Date    time.Time
}

func New(name, message string) Entry {
	return Entry{
		Name:    name,
		Message: message,
		Date:    time.Now(),
	}
}

type Retriever interface {
	GetAll() ([]Entry, error)
	Write(Entry) error
}

type Mocker struct{}

func (m Mocker) GetAll() ([]Entry, error) {
	return []Entry{
		New("test1", "message1"),
		New("test2", "message2"),
	}, nil
}

func (m Mocker) Write(e Entry) error {
	if e.Name == "fail" {
		return errors.New("Set failed")
	}
	return nil
}

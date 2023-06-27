// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type FetchOwnerInput struct {
	ID string `json:"id" bson:"_id"`
}

type FetchProject struct {
	ID string `json:"id" bson:"_id"`
}

type NewOwner struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type NewProject struct {
	OwnerID     string `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type Owner struct {
	ID    string `json:"_id" bson:"_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type Project struct {
	ID          string `json:"_id" bson:"_id"`
	OwnerID     string `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      Status `json:"status"`
}

type Status string

const (
	StatusNotStarted Status = "NOT_STARTED"
	StatusInProgress Status = "IN_PROGRESS"
	StatusCompleted  Status = "COMPLETED"
)

var AllStatus = []Status{
	StatusNotStarted,
	StatusInProgress,
	StatusCompleted,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusNotStarted, StatusInProgress, StatusCompleted:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
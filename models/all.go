package models

import (
	"database/sql/driver"
	"errors"
	"time"

	"github.com/google/uuid"
)

type TestType string

const (
	WAIS_IV     TestType = "WAIS IV"
	MensaNorway TestType = "mensa norway"
)

func (tt *TestType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("type assertion to string failed")
	}
	*tt = TestType(str)
	return nil
}

func (tt TestType) Value() (driver.Value, error) {
	return string(tt), nil
}

type User struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type IQTest struct {
	ID             uuid.UUID `db:"id"`
	UserID         uuid.UUID `db:"user_id"`
	TestType       TestType  `db:"test_type"`
	CorrectAnswers []bool    `db:"correct_answers"`
	Finished       bool      `db:"finished"`
	Result         int       `db:"result"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

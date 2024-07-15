package db

import (
	"iq/models"
	"time"

	"github.com/google/uuid"
)

func CreateUser(user models.User) (models.User, error) {
  db := GetDB()
  sqlStatement := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`
  err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.ID)
  if err != nil {
    return user, err
  }
  return user, nil
}

func UpdateUser(user models.User, id uuid.UUID) (models.User, error) {
  db := GetDB()
  sqlStatement := `
    UPDATE users
    SET name = $2, email = $3, password = $4, updated_at = $5
    WHERE id = $1
    RETURNING id`
  err := db.QueryRow(sqlStatement, id, user.Name, user.Email, user.Password, time.Now()).Scan(&id)
  if err != nil {
    return models.User{}, err
  }
  user.ID = id
  return user, nil
}

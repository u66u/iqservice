package db

import (
	"database/sql"
	"errors"
	"iq/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(user models.User) (models.User, error) {
	db := GetDB()
	sqlStatement := `
      INSERT INTO users (name, email, password)
      VALUES ($1, $2, $3)
      RETURNING id`

	err := db.QueryRow(sqlStatement, user.Name, user.Email, user.Password).Scan(&user.ID)
	if err != nil {
		return models.User{}, err
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

func SelectUser(email, password string) (models.User, error) {
	db := GetDB()
	var user models.User
	sqlStatement := `
      SELECT id, name, email, password
      FROM users
      WHERE email = $1`
	err := db.QueryRow(sqlStatement, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	// Compare the stored hashed password with the provided password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return models.User{}, errors.New("invalid password")
	}

	return user, nil
}

func UserExistsByEmail(email string) (bool, error) {
    db := GetDB()
    var user models.User
    sqlStatement := `
        SELECT id 
        FROM users 
        WHERE email = $1`
    err := db.QueryRow(sqlStatement, email).Scan(&user.ID)
    if err != nil {
        if err == sql.ErrNoRows {
            return false, nil
        }
        return false, err
    }
    return true, nil
}

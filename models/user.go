package models

import (
	"errors"

	"github.com/wtran29/event-booking/db"
	"github.com/wtran29/event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES ($1, $2) RETURNING id`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	var userID int64
	err = stmt.QueryRow(u.Email, hashedPassword).Scan(&userID)
	if err != nil {
		return err
	}
	u.ID = userID
	return err

}

func (u User) ValidatePassword() error {
	query := `SELECT password FROM users WHERE email = $1`
	row := db.DB.QueryRow(query, u.Email)
	var retrievePassword string
	err := row.Scan(&retrievePassword)
	if err != nil {
		return errors.New("invalid credentials!")
	}

	isPasswordValid := utils.ComparePasswordHash(u.Password, retrievePassword)
	if !isPasswordValid {
		return errors.New("invalid credentials!")
	}
	return nil
}

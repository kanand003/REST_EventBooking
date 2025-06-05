package models

import (
	"errors"

	"github.com/rest-api-event/db"
	"github.com/rest-api-event/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u User) Save() error {
	query := "INSERT INTO users (email, password) VALUES (?, ?)"
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	// Hash the password before saving
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := statement.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userid, err := result.LastInsertId()

	u.ID = userid
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}

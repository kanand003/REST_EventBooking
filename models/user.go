package models

import (
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

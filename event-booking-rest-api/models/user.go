package models

import (
	"errors"
	"fmt"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	const query = `
	INSERT INTO users(email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := res.LastInsertId()
	u.ID = id
	return nil
}

func (u *User) ValidateCredentials() error {
	const query = `
	SELECT password FROM users WHERE email = ?
	`
	var hashedPassword string
	err := db.DB.QueryRow(query, u.Email).Scan(&hashedPassword)
	if err != nil {
		return fmt.Errorf("Credentials invalid: %v", err)
	}
	if !utils.ComparePasswords(hashedPassword, u.Password) {
		return errors.New("Credentials invalid")
	}
	u.Password = hashedPassword
	return nil
}

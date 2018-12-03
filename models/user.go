package models

import "time"

//User represent an user struct
type User struct {
	ID        int       `scheme:"id"`
	Name      string    `scheme:"name"`
	Email     string    `scheme:"email"`
	Pwd       string    `scheme:"pwd"`
	CreatedAt time.Time `scheme:"created_at"`
	UpdatedAt time.Time `scheme:"created_at"`
	Valid     bool      `scheme:"valid"`
}

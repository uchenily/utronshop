package models

//Order represent an order struct
type Order struct {
	ID       int    `scheme:"id"`
	PID      int    `scheme:"pid"`
	Username string `scheme:"username"`
	Email    string `scheme:"email"`
	Tel      string `scheme:"tel"`
	Addr     string `scheme:"addr"`
}

package domain

import "time"

type User struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Email         string    `json:"email"`
	Password      string    `json:"password"`
	Registered_at time.Time `json:"registered_at"`
}

var (
	Users = map[int]*User{}
	Seq   = 1
)

package reamaze

import "time"

type Note struct {
	ID        string    `json:"id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Creator   struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	} `json:"creator"`
}

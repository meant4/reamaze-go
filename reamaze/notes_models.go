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

type GetNotesResponse []Note
type DeleteNoteResponse Note
type UpdateNoteResponse Note
type CreateNoteResponse Note
type UpdateNoteRequest CreateNoteRequest

// CreatorEmail is optional and should be the staff email address for the Re:amaze staff user who you want to be attributed to creating the note. Otherwise, the creator will not be updated.
// CreatedAt is optional and will not be updated if not passed in.
type CreateNoteRequest struct {
	Body         string    `json:"body"`
	CreatorEmail string    `json:"creator_email"`
	CreatedAt    time.Time `json:"created_at"`
}

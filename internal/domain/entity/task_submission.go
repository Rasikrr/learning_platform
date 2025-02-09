package entity

import "time"

type TaskSubmission struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	UserID    string    `json:"user_id"`
	Input     string    `json:"input"`
	Error     *string   `json:"error"`
	Passed    bool      `json:"passed"`
	CreatedAt time.Time `json:"created_at"`
}

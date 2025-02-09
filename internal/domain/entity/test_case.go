package entity

import "time"

type TestCase struct {
	ID             string    `json:"id"`
	TaskID         string    `json:"task_id"`
	Input          string    `json:"input"`
	ExpectedOutput string    `json:"expected_output"`
	IsPublic       bool      `json:"is_public"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

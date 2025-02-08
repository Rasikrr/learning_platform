package entity

import (
	"github.com/Rasikrr/learning_platform/internal/domain/enum"
	"github.com/google/uuid"
	"time"
)

//go:generate easyjson -all course.go

type Course struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	ImageURL    *string   `json:"image_url"`
	Category    Category  `json:"category"`
	Description string    `json:"description"`
	Topics      []*Topic  `json:"topics,omitempty"`
	CreatedBy   uuid.UUID `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Category struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy uuid.UUID `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Topic struct {
	ID             uuid.UUID        `json:"id"`
	CourseID       uuid.UUID        `json:"course_id"`
	Title          string           `json:"title"`
	Description    string           `json:"description"`
	Content        *TopicContent    `json:"content,omitempty"`
	Quizzes        []*Quiz          `json:"quizzes,omitempty"`
	PracticalTasks []*PracticalTask `json:"practical_tasks,omitempty"`
	OrderNumber    int              `json:"order_number"`
	CreatedAt      time.Time        `json:"created_at"`
	UpdatedAt      time.Time        `json:"updated_at"`
}

type TopicContent struct {
	ID                  uuid.UUID `json:"id"`
	TopicID             uuid.UUID `json:"topic_id"`
	Content             string    `json:"content"`
	AdditionalResources []string  `json:"additional_resources"`
	VideoURLs           []string  `json:"video_urls"`
	ImageURLs           []string  `json:"image_urls"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type Quiz struct {
	ID             uuid.UUID `json:"id"`
	TopicID        uuid.UUID `json:"topic_id"`
	Question       string    `json:"question"`
	Options        []string  `json:"options"`
	CorrectAnswers []bool    `json:"correct_answers"`
	MultipleChoice bool      `json:"multiple_choice"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type PracticalTask struct {
	ID              uuid.UUID       `json:"id"`
	TopicID         uuid.UUID       `json:"topic_id"`
	Description     string          `json:"description"`
	DifficultyLevel enum.Difficulty `json:"difficulty_level"`
	StarterCode     string          `json:"starter_code"`
	ExpectedOutput  string          `json:"expected_output"`
	OrderNumber     int             `json:"order_number"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
}

type AnswerQuiz struct {
	QuestionID string `json:"question_id"`
	Answer     []bool `json:"answer"`
}

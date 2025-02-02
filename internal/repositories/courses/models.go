package courses

import (
	entities "github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/google/uuid"
	"time"
)

type course struct {
	ID          uuid.UUID
	Title       string
	ImageURL    *string
	TopicID     uuid.UUID
	Description string
	CreatedBy   uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type courses []course

func (c course) convert() (*entities.Course, error) {
	return &entities.Course{
		ID:          c.ID,
		Title:       c.Title,
		ImageURL:    c.ImageURL,
		TopicID:     c.TopicID,
		Description: c.Description,
		CreatedBy:   c.CreatedBy,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}, nil
}

func (c courses) convert() ([]*entities.Course, error) {
	out := make([]*entities.Course, len(c))
	for i, mm := range c {
		res, err := mm.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}

type topic struct {
	ID        uuid.UUID
	Name      string
	CreatedBy uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type topics []topic

func (t topic) convert() (*entities.Topic, error) {
	return &entities.Topic{
		ID:        t.ID,
		Name:      t.Name,
		CreatedBy: t.CreatedBy,
		CreatedAt: t.CreatedAt,
		UpdatedAt: t.UpdatedAt,
	}, nil
}

func (t topics) convert() ([]*entities.Topic, error) {
	out := make([]*entities.Topic, len(t))
	for i, mm := range t {
		res, err := mm.convert()
		if err != nil {
			return nil, err
		}
		out[i] = res
	}
	return out, nil
}

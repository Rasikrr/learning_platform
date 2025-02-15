package courses

import "github.com/Rasikrr/learning_platform/internal/domain/entity"

//go:generate easyjson -all models.go

type createCourseRequest struct {
	Title       string  `json:"title"`
	ImageURL    *string `json:"image"`
	CategoryID  string  `json:"category_id"`
	Description string  `json:"description"`
}

func (r createCourseRequest) ToEntity(session *entity.Session) *entity.CreateCourseParams {
	return &entity.CreateCourseParams{
		Title:       r.Title,
		ImageURL:    r.ImageURL,
		CategoryID:  r.CategoryID,
		Description: r.Description,
		CreatedBy:   session.UserID.String(),
	}
}

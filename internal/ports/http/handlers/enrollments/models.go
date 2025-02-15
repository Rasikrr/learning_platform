package enrollments

import "github.com/Rasikrr/learning_platform/internal/domain/entity"

//go:generate easyjson -all models.go
type enrollToCourseRequest struct {
	CourseID string `json:"course_id"`
}

type getUserEnrollmentsResponse struct {
	Enrollments []enrollment `json:"enrollments"`
}

func convertToGetUserEnrollmentsResponse(enrollments []*entity.Enrollment) getUserEnrollmentsResponse {
	enrollmentResponses := make([]enrollment, len(enrollments))
	for i, enrollment := range enrollments {
		enrollmentResponses[i] = convertToEnrollment(enrollment)
	}
	return getUserEnrollmentsResponse{
		Enrollments: enrollmentResponses,
	}
}

type enrollment struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Course    course `json:"course"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func convertToEnrollment(e *entity.Enrollment) enrollment {
	return enrollment{
		ID:        e.ID.String(),
		UserID:    e.UserID.String(),
		Course:    convertToCourse(e.Course),
		Status:    e.Status.String(),
		CreatedAt: e.CreatedAt.String(),
		UpdatedAt: e.UpdatedAt.String(),
	}
}

type course struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	ImageURL    *string `json:"image_url"`
	Description string  `json:"description"`
	CreatedBy   string  `json:"created_by"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

func convertToCourse(c entity.Course) course {
	return course{
		ID:          c.ID.String(),
		Title:       c.Title,
		ImageURL:    c.ImageURL,
		Description: c.Description,
		CreatedBy:   c.CreatedBy.String(),
		CreatedAt:   c.CreatedAt.String(),
		UpdatedAt:   c.UpdatedAt.String(),
	}
}

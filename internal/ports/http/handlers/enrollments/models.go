package commands

//go:generate easyjson -all models.go
type enrollToCourseRequest struct {
	CourseID string `json:"course_id"`
}

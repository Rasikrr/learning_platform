package queries

import (
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/samber/lo"
	"net/http"
	"strconv"
	"time"
)

//go:generate easyjson -all models.go

var (
	errCourseOrTopicIDEmpty = errors.New("course_id or topic_id is empty")
)

type getTopicContentRequest struct {
	getTopicQuizzesRequest
}

type getTopicQuizzesRequest struct {
	CourseID string
	TopicID  string
}

type getTopicTasksRequest struct {
	CourseID string
	TopicID  string
	Order    int
}

type getCoursesListRequest struct {
	Limit         int      `json:"limit"`
	Offset        int      `json:"offset"`
	CategoriesIDs []string `json:"categories_ids"`
}

type getCourseResponse struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	ImageURL    *string   `json:"image_url"`
	Category    category  `json:"category"`
	Description string    `json:"description"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type getCourseDetailedResponse struct {
	Course getCourseResponse `json:"course"`
	Topics []topic           `json:"topics"`
}

type getCoursesListResponse struct {
	Courses []getCourseResponse `json:"courses"`
}

type category struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type topic struct {
	ID          string    `json:"id"`
	CourseID    string    `json:"course_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	OrderNumber int       `json:"order_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func convertToGetCoursesListResponse(courses []*entity.Course) getCoursesListResponse {
	courseResponses := make([]getCourseResponse, len(courses))
	for i, course := range courses {
		courseResponses[i] = convertCourse(course)
	}
	return getCoursesListResponse{
		Courses: courseResponses,
	}
}

type getCategoriesListResponse struct {
	Categories []category `json:"categories"`
}

func convertToGetCategoriesListResponse(categories []*entity.Category) getCategoriesListResponse {
	categoryResponses := make([]category, len(categories))
	for i, cat := range categories {
		categoryResponses[i] = convertToCategory(cat)
	}
	return getCategoriesListResponse{
		Categories: categoryResponses,
	}
}

func convertToCategory(cat *entity.Category) category {
	return category{
		ID:        cat.ID.String(),
		Name:      cat.Name,
		CreatedBy: cat.CreatedBy.String(),
		CreatedAt: cat.CreatedAt,
		UpdatedAt: cat.UpdatedAt,
	}
}

func convertToTopic(t *entity.Topic) topic {
	return topic{
		ID:          t.ID.String(),
		Title:       t.Title,
		Description: t.Description,
		CourseID:    t.CourseID.String(),
		OrderNumber: t.OrderNumber,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}

func convertToGetCourseDetailedResponse(c *entity.Course) getCourseDetailedResponse {
	return getCourseDetailedResponse{
		Course: convertCourse(c),
		Topics: lo.Map(c.Topics, func(t *entity.Topic, _ int) topic {
			return convertToTopic(t)
		}),
	}
}

func convertCourse(c *entity.Course) getCourseResponse {
	return getCourseResponse{
		ID:          c.ID.String(),
		Title:       c.Title,
		ImageURL:    c.ImageURL,
		Category:    convertToCategory(&c.Category),
		Description: c.Description,
		CreatedBy:   c.CreatedBy.String(),
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (r *getCoursesListRequest) toParams() *entity.GetCoursesParams {
	if r.Limit == 0 {
		r.Limit = 15
	}
	return &entity.GetCoursesParams{
		Limit:         r.Limit,
		Offset:        r.Offset,
		CategoriesIDs: r.CategoriesIDs,
	}
}

func (req *getTopicQuizzesRequest) GetParameters(r *http.Request) error {
	req.CourseID = r.URL.Query().Get("course_id")
	req.TopicID = r.URL.Query().Get("topic_id")
	if req.CourseID == "" || req.TopicID == "" {
		return errCourseOrTopicIDEmpty
	}
	return nil
}

func (req *getTopicTasksRequest) GetParameters(r *http.Request) error {
	var err error
	req.CourseID = r.URL.Query().Get("course_id")
	req.TopicID = r.URL.Query().Get("topic_id")
	req.Order, err = strconv.Atoi(r.URL.Query().Get("order"))
	if err != nil {
		return err
	}
	if req.CourseID == "" || req.TopicID == "" {
		return errCourseOrTopicIDEmpty
	}
	return nil
}

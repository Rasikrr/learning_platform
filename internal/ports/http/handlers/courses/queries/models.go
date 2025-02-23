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

type getTopicQuizzesResponse struct {
	Quizzes []quizz `json:"quizzes"`
	Passed  bool    `json:"passed"`
}

type quizz struct {
	ID             string    `json:"id"`
	TopicID        string    `json:"topic_id"`
	Question       string    `json:"question"`
	Options        []string  `json:"options"`
	Answers        []bool    `json:"answers,omitempty"`
	MultipleChoice bool      `json:"multiple_choice"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type getTaskResponse struct {
	ID              string          `json:"id"`
	TopicID         string          `json:"topic_id"`
	Description     string          `json:"description"`
	DifficultyLevel string          `json:"difficulty_level"`
	StarterCode     string          `json:"starter_code"`
	ExpectedOutput  *string         `json:"expected_output"`
	OrderNumber     int             `json:"order_number"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	TestCases       bool            `json:"test_cases"`
	Language        string          `json:"language"`
	Solved          bool            `json:"solved"`
	Submission      *taskSubmission `json:"submission"`
}

type taskSubmission struct {
	ID        string    `json:"id"`
	TaskID    string    `json:"task_id"`
	UserID    string    `json:"user_id"`
	Input     string    `json:"input"`
	Passed    bool      `json:"passed"`
	Error     *string   `json:"error"`
	CreatedAt time.Time `json:"created_at"`
}

func convertToGetTaskResponse(task *entity.PracticalTask, submission *entity.TaskSubmission) getTaskResponse {
	return getTaskResponse{
		ID:              task.ID.String(),
		TopicID:         task.TopicID.String(),
		Description:     task.Description,
		DifficultyLevel: task.DifficultyLevel.String(),
		StarterCode:     task.StarterCode,
		ExpectedOutput:  task.ExpectedOutput,
		OrderNumber:     task.OrderNumber,
		CreatedAt:       task.CreatedAt,
		UpdatedAt:       task.UpdatedAt,
		TestCases:       task.TestCases,
		Language:        task.Language.String(),
		Solved:          submission != nil,
		Submission:      convertToTaskSubmission(submission),
	}
}

func convertToTaskSubmission(submission *entity.TaskSubmission) *taskSubmission {
	if submission == nil {
		return nil
	}
	return &taskSubmission{
		ID:        submission.ID,
		TaskID:    submission.TaskID,
		UserID:    submission.UserID,
		Input:     submission.Input,
		Passed:    submission.Passed,
		Error:     submission.Error,
		CreatedAt: submission.CreatedAt,
	}
}

func convertToGetTopicQuizzesResponse(quizzes []*entity.Quiz, passed bool) getTopicQuizzesResponse {
	return getTopicQuizzesResponse{
		Passed: passed,
		Quizzes: lo.Map(quizzes, func(e *entity.Quiz, _ int) quizz {
			return convertQuiz(e, passed)
		}),
	}
}

func convertQuiz(e *entity.Quiz, passed bool) quizz {
	q := quizz{
		ID:             e.ID.String(),
		TopicID:        e.TopicID.String(),
		Question:       e.Question,
		Options:        e.Options,
		MultipleChoice: e.MultipleChoice,
		CreatedAt:      e.CreatedAt,
		UpdatedAt:      e.UpdatedAt,
	}
	if passed {
		q.Answers = e.CorrectAnswers
	}
	return q
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
	req.CourseID = r.PathValue("course_id")
	req.TopicID = r.PathValue("topic_id")
	if req.CourseID == "" || req.TopicID == "" {
		return errCourseOrTopicIDEmpty
	}
	return nil
}

func (req *getTopicTasksRequest) GetParameters(r *http.Request) error {
	var err error
	req.CourseID = r.PathValue("course_id")
	req.TopicID = r.PathValue("topic_id")
	req.Order, err = strconv.Atoi(r.PathValue("order"))
	if err != nil {
		return err
	}
	if req.CourseID == "" || req.TopicID == "" {
		return errCourseOrTopicIDEmpty
	}
	return nil
}

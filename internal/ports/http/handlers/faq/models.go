package faq

//go:generate easyjson -all models.go

type postQuestionRequest struct {
	Title      string `json:"title"`
	Body       string `json:"body"`
	CategoryID string `json:"category_id"`
}

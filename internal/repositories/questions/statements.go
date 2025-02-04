package question

const (
	createQuestionStmt = `INSERT INTO faq_questions (id, category_id, user_id, title, body) VALUES ($1, $2, $3, $4, $5)`
)

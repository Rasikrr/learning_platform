package question

const (
	createQuestionStmt = `INSERT INTO faq_questions (id, category_id, user_id, title, body) VALUES ($1, $2, $3, $4, $5)`
	getByIDStmt        = `SELECT id, category_id, user_id, title, body, created_at, updated_at FROM faq_questions WHERE id = $1`
)

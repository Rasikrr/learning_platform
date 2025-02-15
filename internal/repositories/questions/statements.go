package question

const (
	createQuestionStmt = `INSERT INTO faq_questions (id, category_id, user_id, title, body, image_url) VALUES ($1, $2, $3, $4, $5, $6)`
	getByIDStmt        = `SELECT id, category_id, user_id, title, body, created_at, updated_at, image_url FROM faq_questions WHERE id = $1`
)

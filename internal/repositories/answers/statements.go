package answers

const (
	createAnswerStmt = `INSERT INTO faq_answers (id, question_id, user_id, body) VALUES ($1, $2, $3, $4)`
)

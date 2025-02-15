package answers

const (
	createAnswerStmt = `INSERT INTO faq_answers (id, question_id, user_id, body) VALUES ($1, $2, $3, $4)`

	getByQuestionIDStmt = `SELECT id, question_id, user_id, body 
							FROM faq_answers 
							WHERE question_id = $1 
							ORDER BY updated_at DESC 
							LIMIT $2 
							OFFSET $3`
)

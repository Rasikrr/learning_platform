package quizzes

const (
	getByTopicIDStmt = `SELECT id, topic_id, question, options, correct_answers, created_at, updated_at, multiple_choice
						FROM quizzes
						WHERE topic_id = $1`
	getByTopicIDsStmt = `SELECT id, topic_id, question, options, correct_answers, created_at, updated_at, multiple_choice
						FROM quizzes
						WHERE topic_id = ANY ($1)`
)

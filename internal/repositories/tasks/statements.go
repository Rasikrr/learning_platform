package tasks

const (
	getByTopicIDsStmt = `SELECT id, topic_id, description, difficulty_level, starter_code, expected_output, created_at, updated_at 
						FROM practical_tasks 
						WHERE topic_id = ANY ($1)`
	getByTopicIDStmt = `SELECT id, topic_id, description, difficulty_level, starter_code, expected_output, created_at, updated_at 
						FROM practical_tasks 
						WHERE topic_id = $1`
)

package tasks

const (
	getByIDStmt = `SELECT id, topic_id, description, difficulty_level, starter_code, expected_output, 
       					created_at, updated_at, order_number, test_cases, language
						FROM practical_tasks 
						WHERE id = $1`

	getByTopicIDsStmt = `SELECT id, topic_id, description, difficulty_level, starter_code, expected_output, 
       					created_at, updated_at, order_number, test_cases, language
						FROM practical_tasks 
						WHERE topic_id = ANY ($1)`
	getByTopicIDAndOrderNumStmt = `SELECT id, topic_id, description, difficulty_level, starter_code, expected_output,
       					created_at, updated_at, order_number, test_cases, language
						FROM practical_tasks 
						WHERE topic_id = $1 AND order_number = $2`
)

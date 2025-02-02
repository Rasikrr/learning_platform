package topics

const (
	getByCourseIDStmt = `SELECT id, course_id, title, order_number, created_at, updated_at 
						FROM topics 
						WHERE course_id = $1
						ORDER BY order_number`
)

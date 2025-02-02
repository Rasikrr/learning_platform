package courses

const (
	getCoursesStmt = `SELECT id, title, image_url, topic_id, description, created_by, created_at, updated_at 
					FROM courses 
					ORDER BY created_at DESC 
					LIMIT $1 OFFSET $2`

	getCoursesTopicsByIdsStmt = `SELECT id, name, created_by, created_at, updated_at 
							FROM course_topics 
							WHERE id = ANY ($1)`
)

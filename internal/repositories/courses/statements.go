package courses

// nolint
const (
	getCoursesStmt = `SELECT id, title, image_url, category_id, description, created_by, created_at, updated_at 
					FROM courses 
					ORDER BY created_at DESC 
					LIMIT $1 OFFSET $2`

	getCourseByIDStmt = `SELECT id, title, image_url, category_id, description, created_by, created_at, updated_at 
					FROM courses 
					WHERE id = $1`
)

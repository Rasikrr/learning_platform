package categories

const (
	getCategoriesByIDsStmt = `SELECT id, name, created_by, created_at, updated_at 
							FROM course_category 
							WHERE id = ANY ($1)`

	getAllCategoriesStmt = `SELECT id, name, created_by, created_at, updated_at 
					FROM course_category`
)

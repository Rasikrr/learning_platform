package categories

const (
	getCategoriesByIdsStmt = `SELECT id, name, created_by, created_at, updated_at 
							FROM course_category 
							WHERE id = ANY ($1)`

	getAllCategoriesStmt = `SELECT id, name, created_by, created_at, updated_at 
					FROM course_category`
)

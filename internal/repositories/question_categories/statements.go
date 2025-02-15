// nolint: revive, stylecheck
package question_categories

const (
	getByIDStmt  = `SELECT id, name, created_at, updated_at FROM faq_categories WHERE id = $1`
	getAllStmt   = `SELECT id, name, created_at, updated_at FROM faq_categories`
	getByIDsStmt = `SELECT id, name, created_at, updated_at FROM faq_categories WHERE id = ANY($1)`
)

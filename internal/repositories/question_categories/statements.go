// nolint: revive, stylecheck
package question_categories

const (
	getByIDStmt = `SELECT id, name, created_at, updated_at FROM faq_categories WHERE id = $1`
)

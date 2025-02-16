package users

// nolint: gosec
const (
	getByEmailStmt = `SELECT id, name, last_name, email, password, account_role, created_at, updated_at, deleted_at
						FROM users WHERE email = $1`
	createStmt        = `INSERT INTO users (id, name, last_name, email, password, account_role) VALUES ($1, $2, $3, $4, $5, $6)`
	deleteStmt        = `UPDATE users SET deleted_at = NOW() WHERE id = $1`
	resetPasswordStmt = `UPDATE users SET password = $2 WHERE email = $1`
	getByIDStmt       = `SELECT id, name, last_name, email, password, account_role, created_at, updated_at, deleted_at
						FROM users WHERE id = $1 AND deleted_at IS NULL`
	getByIDsStmt = `SELECT id, name, last_name, email, password, account_role, created_at, updated_at, deleted_at
						FROM users WHERE id = ANY($1)`
	updateStmt = `UPDATE users SET name = $1, last_name = $2 WHERE id = $3`
)

package users

// nolint: gosec
const (
	getByEmailStmt = `SELECT id, name, last_name, email, password, account_role, created_at, updated_at, deleted_at
						FROM users WHERE email = $1`
	createStmt        = `INSERT INTO users (id, name, last_name, email, password, account_role) VALUES ($1, $2, $3, $4, $5, $6)`
	deleteStmt        = `UPDATE users SET deleted_at = NOW() WHERE email = $1`
	resetPasswordStmt = `UPDATE users SET password = $2 WHERE email = $1`
)

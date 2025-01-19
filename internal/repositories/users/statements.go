package users

const (
	getByEmailStmt = `SELECT id, name, last_name, email, password, created_at, updated_at, deleted_at 
						FROM users WHERE email = $1`
	createStmt = `INSERT INTO users (id, name, last_name, email, password) VALUES ($1, $2, $3, $4, $5)`
	deleteStmt = `UPDATE users SET deleted_at = NOW() WHERE email = $1`
)

// nolint: revive, stylecheck, unused, gosec
package tasks_submissions

const (
	insertStmt = `INSERT INTO practical_tasks_submissions (task_id, user_id, input, passed, error, created_at)
					VALUES ($1, $2, $3, $4, $5, now())`
	checkPassesStmt = `SELECT EXISTS(SELECT id 
								FROM practical_tasks_submissions 
								WHERE user_id = $1 AND task_id = $2 AND passed = true)`
)

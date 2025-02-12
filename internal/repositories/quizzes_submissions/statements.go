// nolint: revive, stylecheck
package quizzes_submissions

// nolint: gosec
const (
	checkPassesStmt = `SELECT EXISTS(SELECT id 
								FROM quizzes_submissions 
								WHERE user_id = $1 AND topic_id = $2 AND passed = true)`

	upsertStmt = `INSERT INTO quizzes_submissions (user_id, course_id, topic_id, passed, created_at, updated_at)
					VALUES ($1, $2, $3, $4, now(), now())
    				ON CONFLICT (user_id, topic_id) DO UPDATE SET passed = excluded.passed, updated_at = now()`
)

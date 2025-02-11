// nolint: revive, stylecheck, unused
package test_cases

const (
	getByTaskIDStmt = `SELECT id, task_id, input, expected_output, is_public, created_at, updated_at
						FROM test_cases
						WHERE task_id = $1`
)

package enrollments

const (
	enrollUserStmt             = `INSERT INTO course_enrollments (user_id, course_id) VALUES ($1, $2)`
	getByUserIDAndCourseIDStmt = `SELECT EXISTS(SELECT id 
								FROM course_enrollments 
								WHERE user_id = $1 AND course_id = $2)`

	getUserEnrollmentsStmt = `SELECT * FROM course_enrollments WHERE user_id = $1`
)

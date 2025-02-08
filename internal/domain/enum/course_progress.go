package enum

//go:generate enumer -type=CourseProgress -json -trimprefix CourseProgress -transform=snake -output course_progress_enumer.go
type CourseProgress uint8

const (
	CourseProgressInProgress CourseProgress = iota
	CourseProgressCompleted
)

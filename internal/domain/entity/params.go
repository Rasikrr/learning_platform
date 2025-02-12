package entity

type GetCoursesParams struct {
	Limit         int
	Offset        int
	CategoriesIDs []string
}

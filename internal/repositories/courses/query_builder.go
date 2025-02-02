package courses

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
)

var (
	coursesTable = "courses c"
	coursesCols  = []string{
		"c.id",
		"c.title",
		"c.image_url",
		"c.category_id",
		"c.description",
		"c.created_by",
		"c.created_at",
		"c.updated_at",
	}

	// nolint
	categoriesCols = []string{
		"t.id",
		"t.name",
		"t.created_by",
		"t.created_at",
		"t.updated_at",
	}

	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
)

func generateQuery(params *entity.GetCoursesParams) (string, []interface{}, error) {
	b := psql.Select(coursesCols...).From(coursesTable)
	if params.Category != nil {
		b = b.Join("course_category cc on cc.id = c.category_id")
		b = b.Where("cc.name = ?", *params.Category)
	}
	// nolint
	b = b.Limit(uint64(params.Limit)).Offset(uint64(params.Offset))
	return b.ToSql()
}

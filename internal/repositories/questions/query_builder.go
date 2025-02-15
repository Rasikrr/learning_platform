package question

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/lib/pq"
)

var (
	psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	cols = []string{
		"id",
		"category_id",
		"user_id",
		"title",
		"body",
		"image_url",
		"created_at",
		"updated_at",
	}
)

func generateQuery(params *entity.GetQuestionsParams) (string, []interface{}, error) {
	q := psql.Select(cols...).
		From("faq_questions q")

	if len(params.CategoryIDs) > 0 {
		q = q.Where("q.category_id = ANY($1)", pq.Array(params.CategoryIDs))
	}
	if params.Limit > 0 {
		q = q.Limit(uint64(params.Limit))
	}
	if params.Offset > 0 {
		q = q.Offset(uint64(params.Offset))
	}
	return q.ToSql()
}

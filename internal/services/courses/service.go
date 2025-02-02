package courses

import (
	"context"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
	"github.com/Rasikrr/learning_platform/internal/repositories/courses"
	"github.com/google/uuid"
	"github.com/samber/lo"
)

type Service interface {
	GetCourses(ctx context.Context, limit, offset int) ([]*entity.Course, error)
}

type service struct {
	coursesRepository courses.Repository
}

func NewService(coursesRepository courses.Repository) Service {
	return &service{
		coursesRepository: coursesRepository,
	}
}

func (s *service) GetCourses(ctx context.Context, limit, offset int) ([]*entity.Course, error) {
	courses, err := s.coursesRepository.GetCourses(ctx, limit, offset)
	if err != nil {
		return nil, err
	}
	if err = s.mergeCourses(ctx, courses...); err != nil {
		return nil, err
	}
	return courses, nil
}

func (s *service) mergeCourses(ctx context.Context, course ...*entity.Course) error {
	topicsIDs := lo.Uniq(lo.Map(course, func(c *entity.Course, _ int) uuid.UUID {
		return c.TopicID
	}))
	topics, err := s.coursesRepository.GetTopicsByIds(ctx, topicsIDs)
	if err != nil {
		return err
	}
	topicsMap := lo.SliceToMap(topics, func(t *entity.Topic) (uuid.UUID, *entity.Topic) {
		return t.ID, t
	})
	for _, c := range course {
		if t, ok := topicsMap[c.TopicID]; ok && t != nil {
			c.Topic = *t
		}
	}
	return nil
}

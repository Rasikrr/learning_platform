package courses

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	cache2 "github.com/Rasikrr/learning_platform/internal/cache"
	"github.com/Rasikrr/learning_platform/internal/domain/entity"
)

var (
	contentKey       = "content:%s"
	practiceTasksKey = "practical_tasks:%s:%d"
	quizzesKey       = "quizzes:%s"
)

type Cache interface {
	SetContentByTopicID(ctx context.Context, topicID string, content *entity.TopicContent) error
	GetContentByTopicID(ctx context.Context, topicID string) (*entity.TopicContent, error)

	GetPracticalTaskByTopicIDAndOrder(ctx context.Context, topicID string, order int) (*entity.PracticalTask, error)
	SetPracticalTasksByTopicIDAndOrder(ctx context.Context, topicID string, order int, tasks *entity.PracticalTask) error

	SetQuizzesByTopicID(ctx context.Context, topicID string, quizzes []*entity.Quiz) error
	GetQuizzesByTopicID(ctx context.Context, topicID string) ([]*entity.Quiz, error)
}

type cache struct {
	client cache2.Cache
}

func NewCache(client cache2.Cache) Cache {
	return &cache{
		client: client,
	}
}

func (c *cache) GetContentByTopicID(ctx context.Context, topicID string) (*entity.TopicContent, error) {
	key := genContentKey(topicID)
	content, err := c.client.Get(ctx, key)
	if err != nil {
		if errors.Is(err, cache2.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	val, ok := content.(string)
	if !ok {
		return nil, errors.New("content is not string")
	}
	var contentEntity entity.TopicContent
	if err = json.Unmarshal([]byte(val), &contentEntity); err != nil {
		return nil, err
	}
	return &contentEntity, nil
}

func (c *cache) SetContentByTopicID(ctx context.Context, topicID string, content *entity.TopicContent) error {
	key := genContentKey(topicID)
	bb, err := json.Marshal(content)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, bb)
}

func (c *cache) GetPracticalTaskByTopicIDAndOrder(ctx context.Context, topicID string, order int) (*entity.PracticalTask, error) {
	key := genPracticalTasksKey(topicID, order)
	tasks, err := c.client.Get(ctx, key)
	if err != nil {
		if errors.Is(err, cache2.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	val, ok := tasks.(string)
	if !ok {
		return nil, errors.New("tasks is not string")
	}
	var tasksEntity entity.PracticalTask
	if err = json.Unmarshal([]byte(val), &tasksEntity); err != nil {
		return nil, err
	}
	return &tasksEntity, nil
}

func (c *cache) SetPracticalTasksByTopicIDAndOrder(ctx context.Context, topicID string, order int, tasks *entity.PracticalTask) error {
	key := genPracticalTasksKey(topicID, order)
	bb, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, bb)
}

func (c *cache) SetQuizzesByTopicID(ctx context.Context, topicID string, quizzes []*entity.Quiz) error {
	key := genQuizzesKey(topicID)
	bb, err := json.Marshal(quizzes)
	if err != nil {
		return err
	}
	return c.client.Set(ctx, key, bb)
}

func (c *cache) GetQuizzesByTopicID(ctx context.Context, topicID string) ([]*entity.Quiz, error) {
	key := genQuizzesKey(topicID)
	quizzes, err := c.client.Get(ctx, key)
	if err != nil {
		if errors.Is(err, cache2.ErrNotFound) {
			return nil, nil
		}
		return nil, err
	}
	val, ok := quizzes.(string)
	if !ok {
		return nil, errors.New("quizzes is not string")
	}
	var quizzesEntity []*entity.Quiz
	if err = json.Unmarshal([]byte(val), &quizzesEntity); err != nil {
		return nil, err
	}
	return quizzesEntity, nil
}

func genContentKey(topicID string) string {
	return fmt.Sprintf(contentKey, topicID)
}

func genPracticalTasksKey(topicID string, order int) string {
	return fmt.Sprintf(practiceTasksKey, topicID, order)
}

func genQuizzesKey(topicID string) string {
	return fmt.Sprintf(quizzesKey, topicID)
}

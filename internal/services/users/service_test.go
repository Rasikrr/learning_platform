package users

import (
	"context"
	"errors"
	"github.com/Rasikrr/learning_platform/internal/domain/entities"
	userRepoMock "github.com/Rasikrr/learning_platform/internal/repositories/mocks/users"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_UsersService(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("GetUser", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)
		testUser := &entities.User{
			ID:        uuid.New(),
			Name:      "Rasik",
			Email:     "rasikrr@gmail.com",
			Password:  "123456",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockRepo.EXPECT().
			GetByEmail(ctx, gomock.Any()).
			Return(testUser, nil)

		user, err := usersService.GetByEmail(ctx, testUser.Email)
		require.NoError(t, err)
		require.Equal(t, testUser.Email, user.Email)
		t.Log(user)
	})

	t.Run("UpdateUserNotExists", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)

		mockRepo.EXPECT().
			GetByEmail(ctx, gomock.Any()).
			Return(nil, errors.New("user with this email not found"))

		err := usersService.UpdateName(ctx, "rs.t.95@mail.", "Rasik", "1234")
		require.EqualError(t, err, "user with this email not found")
	})

	t.Run("UpdateUserWrongPassword", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)
		testUser := entities.NewUser("Rassul", "rs.t.95@mail.ru", "rasik1234")

		mockRepo.EXPECT().
			GetByEmail(ctx, gomock.Any()).
			Return(testUser, nil)

		err := usersService.UpdateName(ctx, "rs.t.95@mail.ru", "Rasik", "Rasik1234")
		require.EqualError(t, err, "incorrect password")
	})
	t.Run("UpdateUser", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)
		testUser := entities.NewUser("Rassul", "rs.t.95@mail.ru", "rasik1234")

		gomock.InOrder(
			mockRepo.EXPECT().
				GetByEmail(ctx, gomock.Any()).
				Return(testUser, nil),
			mockRepo.EXPECT().
				UpdateName(ctx, gomock.Any(), "Rasik").
				Return(nil),
		)

		err := usersService.UpdateName(ctx, "rs.t.95@mail.ru", "Rasik", "rasik1234")
		require.NoError(t, err)
	})

	t.Run("DeleteUserNotExists", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)

		mockRepo.EXPECT().
			Delete(ctx, gomock.Any()).
			Return(errors.New("user with this email not found"))

		err := usersService.Delete(ctx, "rs.t.95@mail.ru")
		require.EqualError(t, err, "user with this email not found")
	})

	t.Run("DeleteUserSuccess", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)

		mockRepo.EXPECT().
			Delete(ctx, gomock.Any()).
			Return(nil)

		err := usersService.Delete(ctx, "rs.t.95@mail.ru")
		require.NoError(t, err)
	})
	t.Run("GetAll", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)
		testUsers := []*entities.User{
			{
				ID:        uuid.New(),
				Email:     "rs.t.95@mail.ru",
				Name:      "Test",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			{
				ID:        uuid.New(),
				Email:     "rs.t.96@mail.ru",
				Name:      "Test",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
		}
		mockRepo.EXPECT().
			GetAll(ctx).
			Return(testUsers, nil)
		users, err := usersService.GetAll(ctx)
		require.NoError(t, err)
		require.Equal(t, testUsers, users)
	})
	t.Run("CreateUserExists", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)

		testUser := entities.NewUser("Rassul", "rs.t.95@mail.ru", "rasik1234")

		mockRepo.EXPECT().
			GetByEmail(ctx, gomock.Any()).
			Return(testUser, nil)

		err := usersService.Create(ctx, "Ali", "rs.t.95@mail.ru", "rasik3790")
		require.EqualError(t, err, "user with this email already exists")
	})
	t.Run("CreateUserSuccess", func(t *testing.T) {
		mockRepo := userRepoMock.NewMockRepository(ctrl)
		usersService := NewService(mockRepo)

		gomock.InOrder(
			mockRepo.EXPECT().
				GetByEmail(ctx, gomock.Any()).
				Return(nil, errors.New("user with this email not found")),
			mockRepo.EXPECT().
				Create(ctx, gomock.Any()).
				Return(nil),
		)

		err := usersService.Create(ctx, "Ali", "rs.t.95@mail.ru", "rasik3790")
		require.NoError(t, err)
	})

}

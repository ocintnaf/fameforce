package usecases

import (
	"testing"
	"time"

	"github.com/ocintnaf/fameforce/entities"
	"github.com/ocintnaf/fameforce/mocks"
	"github.com/ocintnaf/fameforce/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewUserUsecase(t *testing.T) {
	t.Run("should return an instance of userUsecase with the given userRepository", func(t *testing.T) {
		userRepository := mocks.NewUserRepositoryMock()
		userUsecase := NewUserUsecase(userRepository)

		assert.Equal(t, userUsecase.userRepository, userRepository)
	})
}

func TestUserUsecase_GetByID(t *testing.T) {
	t.Run("should return an error if the User Repository returns an error", func(t *testing.T) {
		userRepository := mocks.NewUserRepositoryMock()
		userRepository.On("FindByID", mock.Anything).Return(&entities.UserEntity{}, assert.AnError)

		userUsecase := NewUserUsecase(userRepository)

		userDTO, err := userUsecase.GetByID("unknown-id")

		assert.Error(t, err)
		assert.Nil(t, userDTO)
		userRepository.AssertCalled(t, "FindByID", "unknown-id")
	})

	t.Run("should return a User DTO", func(t *testing.T) {
		now := time.Now()

		userEntity := entities.NewUserEntity("cf", "cf@gmail.com", types.UserTypeInfluencer, now, now)

		expectedUserDTO := userEntity.ToDTO()

		userRepository := mocks.NewUserRepositoryMock()
		userRepository.On("FindByID", mock.Anything).Return(userEntity, nil)

		userUsecase := NewUserUsecase(userRepository)

		actualUserDTO, err := userUsecase.GetByID("cf")

		assert.NoError(t, err)
		assert.NotNil(t, actualUserDTO)
		assert.Equal(t, expectedUserDTO, actualUserDTO)
		userRepository.AssertCalled(t, "FindByID", "cf")
	})
}

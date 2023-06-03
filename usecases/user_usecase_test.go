package usecases

import (
	"testing"
	"time"

	"github.com/ocintnaf/fameforce/entities"
	"github.com/ocintnaf/fameforce/mocks"
	"github.com/ocintnaf/fameforce/types"
	"github.com/stretchr/testify/assert"
)

func TestUserUsecase_NewUserUsecase(t *testing.T) {
	t.Run("should return an instance of userUsecase with the given userRepository", func(t *testing.T) {
		userRepository := mocks.NewUserRepositoryMock()
		userUsecase := NewUserUsecase(userRepository)

		assert.Equal(t, userUsecase.userRepository, userRepository)
	})
}

func TestUserUsecase_GetAll(t *testing.T) {
	t.Run("should return an error if the userRepository returns an error", func(t *testing.T) {
		userRepository := mocks.NewUserRepositoryMock()
		userRepository.On("FindAll").Return([]entities.UserEntity{}, assert.AnError)

		userUsecase := NewUserUsecase(userRepository)

		userDTOs, err := userUsecase.GetAll()

		assert.Error(t, err)
		assert.Nil(t, userDTOs)
	})

	t.Run("should return a list of userDTOs", func(t *testing.T) {
		now := time.Now()

		userEntityOne := *entities.NewUserEntity("cf", "cf@gmail.com", types.UserTypeInfluencer, now, now)
		userEntityTwo := *entities.NewUserEntity("cr", "cr@gmail.com", types.UserTypeInfluencer, now, now)
		userEntityThree := *entities.NewUserEntity("ar", "ag@gmail.com", types.UserTypeInfluencer, now, now)

		userDTOOne := *userEntityOne.ToDTO()
		userDTOTwo := *userEntityTwo.ToDTO()
		userDTOThree := *userEntityThree.ToDTO()

		userRepository := mocks.NewUserRepositoryMock()
		userRepository.On("FindAll").Return([]entities.UserEntity{
			userEntityOne,
			userEntityTwo,
			userEntityThree,
		}, nil)

		userUsecase := NewUserUsecase(userRepository)

		userDTOs, err := userUsecase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, userDTOs)
		assert.Len(t, userDTOs, 3)
		assert.Equal(t, userDTOOne, userDTOs[0])
		assert.Equal(t, userDTOTwo, userDTOs[1])
		assert.Equal(t, userDTOThree, userDTOs[2])
		userRepository.AssertNumberOfCalls(t, "FindAll", 1)
	})
}

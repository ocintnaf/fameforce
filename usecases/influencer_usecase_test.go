package usecases

import (
	"testing"
	"time"

	"github.com/ocintnaf/fameforce/entities"
	"github.com/ocintnaf/fameforce/mocks"
	"github.com/stretchr/testify/assert"
)

func TestInfluencerUsecase_NewInfluencerUsecase(t *testing.T) {
	t.Run("should return an instance of influencerUsecase with the given influencerRepository", func(t *testing.T) {
		influencerRepository := mocks.NewInfluencerRepositoryMock()
		influencerUsecase := NewInfluencerUsecase(influencerRepository)

		assert.Equal(t, influencerUsecase.influencerRepository, influencerRepository)
	})
}

func TestInfluencerUsecase_GetAll(t *testing.T) {
	t.Run("should return an error if the influencerRepository returns an error", func(t *testing.T) {
		influencerRepository := mocks.NewInfluencerRepositoryMock()
		influencerRepository.On("FindAll").Return([]entities.InfluencerEntity{}, assert.AnError)

		influencerUsecase := NewInfluencerUsecase(influencerRepository)

		influencerDTOs, err := influencerUsecase.GetAll()

		assert.Error(t, err)
		assert.Nil(t, influencerDTOs)
	})

	t.Run("should return a list of influencerDTOs", func(t *testing.T) {
		now := time.Now()

		influencerEntityOne := *entities.NewInfluencerEntity(1, "Chiara Ferragni", "cf@gmail.com", now, now)
		influencerEntityTwo := *entities.NewInfluencerEntity(2, "Cristiano Ronaldo", "cr@gmail.com", now, now)
		influencerEntityThree := *entities.NewInfluencerEntity(3, "Ariana Grande", "ag@gmail.com", now, now)

		influencerDTOOne := *influencerEntityOne.ToDTO()
		influencerDTOTwo := *influencerEntityTwo.ToDTO()
		influencerDTOThree := *influencerEntityThree.ToDTO()

		influencerRepository := mocks.NewInfluencerRepositoryMock()
		influencerRepository.On("FindAll").Return([]entities.InfluencerEntity{
			influencerEntityOne,
			influencerEntityTwo,
			influencerEntityThree,
		}, nil)

		influencerUsecase := NewInfluencerUsecase(influencerRepository)

		influencerDTOs, err := influencerUsecase.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, influencerDTOs)
		assert.Len(t, influencerDTOs, 3)
		assert.Equal(t, influencerDTOOne, influencerDTOs[0])
		assert.Equal(t, influencerDTOTwo, influencerDTOs[1])
		assert.Equal(t, influencerDTOThree, influencerDTOs[2])
		influencerRepository.AssertNumberOfCalls(t, "FindAll", 1)
	})
}

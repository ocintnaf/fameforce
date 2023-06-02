package usecases

import (
	"fmt"

	"github.com/ocintnaf/fameforce/dtos"
	"github.com/ocintnaf/fameforce/repositories"
)

type influencerUsecase struct {
	influencerRepository repositories.InfluencerRepository
}

type InfluencerUsecase interface {
	GetAll() ([]dtos.InfluencerDTO, error)
}

func NewInfluencerUsecase(influencerRepository repositories.InfluencerRepository) *influencerUsecase {
	return &influencerUsecase{influencerRepository: influencerRepository}
}

func (iu *influencerUsecase) GetAll() ([]dtos.InfluencerDTO, error) {
	var influencerDTOs []dtos.InfluencerDTO

	influencerEntities, err := iu.influencerRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error getting all influencers: %w", err)
	}

	for _, influencerEntity := range influencerEntities {
		influencerDTOs = append(influencerDTOs, *influencerEntity.ToDTO())
	}

	return influencerDTOs, nil
}

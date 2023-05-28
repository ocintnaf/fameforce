package mocks

import (
	"github.com/ocintnaf/fameforce/entities"
	"github.com/stretchr/testify/mock"
)

type influencerRepositoryMock struct {
	mock.Mock
}

func NewInfluencerRepositoryMock() *influencerRepositoryMock {
	return &influencerRepositoryMock{}
}

func (m *influencerRepositoryMock) FindAll() ([]entities.InfluencerEntity, error) {
	args := m.Called()
	return args.Get(0).([]entities.InfluencerEntity), args.Error(1)
}

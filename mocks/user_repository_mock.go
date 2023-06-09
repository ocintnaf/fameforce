package mocks

import (
	"github.com/ocintnaf/fameforce/entities"
	"github.com/stretchr/testify/mock"
)

type userRepositoryMock struct {
	mock.Mock
}

func NewUserRepositoryMock() *userRepositoryMock {
	return &userRepositoryMock{}
}

func (m *userRepositoryMock) FindByID(id string) (*entities.UserEntity, error) {
	args := m.Called(id)

	return args.Get(0).(*entities.UserEntity), args.Error(1)
}

func (m *userRepositoryMock) Create(e *entities.UserEntity) (*entities.UserEntity, error) {
	args := m.Called(e)
	return args.Get(0).(*entities.UserEntity), args.Error(1)
}

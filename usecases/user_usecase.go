package usecases

import (
	"fmt"

	"github.com/ocintnaf/fameforce/dtos"
	"github.com/ocintnaf/fameforce/entities"
	"github.com/ocintnaf/fameforce/repositories"
)

type userUsecase struct {
	userRepository repositories.UserRepository
}

type UserUsecase interface {
	GetByID(id string) (*dtos.UserDTO, error)
	Create(dtos.UserDTO) (*dtos.UserDTO, error)
}

func NewUserUsecase(userRepository repositories.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) GetByID(id string) (*dtos.UserDTO, error) {
	userEntity, err := uu.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	return userEntity.ToDTO(), nil
}

func (uu *userUsecase) Create(dto dtos.UserDTO) (*dtos.UserDTO, error) {
	userEntity := &entities.UserEntity{}
	userEntity.FromDTO(dto)

	savedUser, err := uu.userRepository.Create(userEntity)
	if err != nil {
		return nil, fmt.Errorf("[UserUsecase.Create] Error creating user: %w", err)
	}

	return savedUser.ToDTO(), nil
}

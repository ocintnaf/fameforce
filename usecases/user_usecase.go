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
	GetAll() ([]dtos.UserDTO, error)
	Create(dtos.UserDTO) (*dtos.UserDTO, error)
}

func NewUserUsecase(userRepository repositories.UserRepository) *userUsecase {
	return &userUsecase{userRepository: userRepository}
}

func (uu *userUsecase) GetAll() ([]dtos.UserDTO, error) {
	var userDTOs []dtos.UserDTO

	userEntities, err := uu.userRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("[UserUsecase.GetAll] Error getting all users: %w", err)
	}

	for _, userEntity := range userEntities {
		userDTOs = append(userDTOs, *userEntity.ToDTO())
	}

	return userDTOs, nil
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

package dtos

import (
	"time"

	"github.com/ocintnaf/fameforce/types"
)

type UserDTO struct {
	ID        string         `json:"id" validate:"required"`
	Email     string         `json:"email"`
	Type      types.UserType `json:"type"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
}

func NewUserDTO(
	id string,
	email string,
	userType types.UserType,
	createdAt time.Time,
	updatedAt time.Time,
) *UserDTO {
	return &UserDTO{
		ID:        id,
		Email:     email,
		Type:      userType,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

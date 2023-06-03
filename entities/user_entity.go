package entities

import (
	"time"

	"github.com/ocintnaf/fameforce/dtos"
	"github.com/ocintnaf/fameforce/types"
)

type UserEntity struct {
	BaseEntity[string]
	Email string
	Type  types.UserType
}

func NewUserEntity(
	id string,
	email string,
	userType types.UserType,
	createdAt time.Time,
	updatedAt time.Time,
) *UserEntity {
	return &UserEntity{
		BaseEntity: BaseEntity[string]{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Email: email,
		Type:  userType,
	}
}

func (ue *UserEntity) TableName() string {
	return "users"
}

func (ue *UserEntity) ToDTO() *dtos.UserDTO {
	return dtos.NewUserDTO(ue.ID, ue.Email, ue.Type, ue.CreatedAt, ue.UpdatedAt)
}

func (ue *UserEntity) FromDTO(dto dtos.UserDTO) {
	ue.ID = dto.ID
	ue.Email = dto.Email
	ue.Type = dto.Type
	ue.CreatedAt = dto.CreatedAt
	ue.UpdatedAt = dto.UpdatedAt
}

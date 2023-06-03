package entities

import (
	"time"

	"github.com/ocintnaf/fameforce/dtos"
)

type InfluencerEntity struct {
	BaseEntity
	Name  string
	Email string
}

func NewInfluencerEntity(
	id uint,
	name string,
	email string,
	createdAt time.Time,
	updatedAt time.Time,
) *InfluencerEntity {
	return &InfluencerEntity{
		BaseEntity: BaseEntity{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Name:  name,
		Email: email,
	}
}

func (e *InfluencerEntity) TableName() string {
	return "influencers"
}

func (e *InfluencerEntity) ToDTO() *dtos.InfluencerDTO {
	return dtos.NewInfluencerDTO(e.ID, e.Name, e.Email, e.CreatedAt, e.UpdatedAt)
}

func (e *InfluencerEntity) FromDTO(dto dtos.InfluencerDTO) {
	e.ID = dto.ID
	e.Name = dto.Name
	e.Email = dto.Email
	e.CreatedAt = dto.CreatedAt
	e.UpdatedAt = dto.UpdatedAt
}

package entities

import (
	"time"

	"github.com/ocintnaf/fameforce/dtos"
)

type InfluencerEntity struct {
	BaseEntity[int]
	UserID string
	User   UserEntity
}

func NewInfluencerEntity(
	id int,
	user UserEntity,
	createdAt time.Time,
	updatedAt time.Time,
) *InfluencerEntity {
	return &InfluencerEntity{
		BaseEntity: BaseEntity[int]{
			ID:        id,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		User: user,
	}
}

func (ie *InfluencerEntity) TableName() string {
	return "influencers"
}

func (ie *InfluencerEntity) ToDTO() *dtos.InfluencerDTO {
	return dtos.NewInfluencerDTO(ie.ID, *ie.User.ToDTO(), ie.CreatedAt, ie.UpdatedAt)
}

func (ie *InfluencerEntity) FromDTO(dto dtos.InfluencerDTO) {
	user := &UserEntity{}
	user.FromDTO(dto.User)

	ie.ID = dto.ID
	ie.UserID = dto.User.ID
	ie.User = *user
	ie.CreatedAt = dto.CreatedAt
	ie.UpdatedAt = dto.UpdatedAt
}

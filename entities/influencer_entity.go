package entities

import "time"

type InfluencerEntity struct {
	BaseEntity[int]
	User UserEntity
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

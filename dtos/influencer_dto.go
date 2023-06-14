package dtos

import (
	"time"
)

type InfluencerDTO struct {
	ID        int       `json:"id"`
	User      UserDTO   `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewInfluencerDTO(
	id int,
	user UserDTO,
	createdAt time.Time,
	updatedAt time.Time,
) *InfluencerDTO {
	return &InfluencerDTO{
		ID:        id,
		User:      user,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

package dtos

import "time"

type InfluencerDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewInfluencerDTO(id uint, name string, createdAt time.Time, updatedAt time.Time) *InfluencerDTO {
	return &InfluencerDTO{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

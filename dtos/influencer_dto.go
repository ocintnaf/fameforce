package dtos

import "time"

type InfluencerDTO struct {
	ID        uint      `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

func NewInfluencerDTO(id uint, name string, createdAt time.Time, updatedAt time.Time) *InfluencerDTO {
	return &InfluencerDTO{
		ID:        id,
		Name:      name,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

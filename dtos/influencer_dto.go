package dtos

import "time"

type InfluencerDTO struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewInfluencerDTO(
	id uint,
	name string,
	email string,
	createdAt time.Time,
	updatedAt time.Time,
) *InfluencerDTO {
	return &InfluencerDTO{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

package entities

import "github.com/ocintnaf/fameforce/dtos"

type InfluencerEntity struct {
	ID   uint
	Name string
}

func NewInfluencerEntity(id uint, name string) *InfluencerEntity {
	return &InfluencerEntity{
		ID:   id,
		Name: name,
	}
}

func (e *InfluencerEntity) TableName() string {
	return "influencers"
}

func (e *InfluencerEntity) ToDTO() *dtos.InfluencerDTO {
	return dtos.NewInfluencerDTO(e.ID, e.Name)
}

func (e *InfluencerEntity) FromDTO(dto *dtos.InfluencerDTO) {
	e.ID = dto.ID
	e.Name = dto.Name
}

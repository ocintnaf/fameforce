package dtos

type InfluencerDTO struct {
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func NewInfluencerDTO(id uint, name string) *InfluencerDTO {
	return &InfluencerDTO{
		ID:   id,
		Name: name,
	}
}

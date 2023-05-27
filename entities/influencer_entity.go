package entities

type Influencer struct {
	ID   uint
	Name string
}

func NewInfluencerEntity(id uint, name string) *Influencer {
	return &Influencer{
		ID:   id,
		Name: name,
	}
}

func (e *Influencer) TableName() string {
	return "influencers"
}

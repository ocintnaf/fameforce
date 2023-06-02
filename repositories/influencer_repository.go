package repositories

import (
	"github.com/ocintnaf/fameforce/entities"
	"gorm.io/gorm"
)

type influencerRepository struct {
	db *gorm.DB
}

type InfluencerRepository interface {
	FindAll() ([]entities.InfluencerEntity, error)
	Save(e entities.InfluencerEntity) (entities.InfluencerEntity, error)
}

func NewInfluencerRepository(db *gorm.DB) *influencerRepository {
	return &influencerRepository{db: db}
}

func (ir *influencerRepository) FindAll() ([]entities.InfluencerEntity, error) {
	var influencers []entities.InfluencerEntity

	ir.db.Find(&influencers)

	return influencers, nil
}

func (ir *influencerRepository) Save(e entities.InfluencerEntity) (entities.InfluencerEntity, error) {
	result := ir.db.Save(&e)
	return e, result.Error
}

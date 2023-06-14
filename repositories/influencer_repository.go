package repositories

import (
	"github.com/ocintnaf/fameforce/entities"
	"gorm.io/gorm"
)

type influencerRepository struct {
	db *gorm.DB
}

type InfluencerRepository interface {
	Create(e *entities.InfluencerEntity) (*entities.InfluencerEntity, error)
}

func NewInfluencerRepository(db *gorm.DB) *influencerRepository {
	return &influencerRepository{db: db}
}

func (ir *influencerRepository) Create(e *entities.InfluencerEntity) (*entities.InfluencerEntity, error) {
	result := ir.db.Create(e)
	return e, result.Error
}

package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/ocintnaf/fameforce/entities"
)

const (
	findAllSqlStatement = "SELECT * FROM influencers"
)

type influencerRepository struct {
	db *sqlx.DB
}

type InfluencerRepository interface {
	FindAll() ([]entities.InfluencerEntity, error)
}

func NewInfluencerRepository(db *sqlx.DB) *influencerRepository {
	return &influencerRepository{db: db}
}

func (r *influencerRepository) FindAll() ([]entities.InfluencerEntity, error) {
	var influencers []entities.InfluencerEntity

	rows, err := r.db.Query(findAllSqlStatement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var influencer entities.InfluencerEntity

		if err := rows.Scan(
			&influencer.ID,
			&influencer.Name,
			&influencer.CreatedAt,
			&influencer.UpdatedAt,
		); err != nil {
			return nil, err
		}

		influencers = append(influencers, influencer)
	}

	return influencers, nil
}

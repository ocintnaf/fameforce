package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewInfluencerEntity(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	user := *NewUserEntity("em", "influencer", "elon.musk@twitter.com", createdAt, updatedAt)

	expected := &InfluencerEntity{
		BaseEntity: BaseEntity[int]{
			ID:        1,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		User: user,
	}
	actual := NewInfluencerEntity(
		1,
		user,
		createdAt,
		updatedAt,
	)

	assert.Equal(t, expected, actual)
}

func TestInfluencerEntity_TableName(t *testing.T) {
	expected := "influencers"
	actual := NewInfluencerEntity(
		1,
		UserEntity{},
		time.Now(),
		time.Now(),
	).TableName()

	assert.Equal(t, expected, actual)
}

package entities

import (
	"testing"
	"time"

	"github.com/ocintnaf/fameforce/dtos"
	"github.com/stretchr/testify/assert"
)

func TestInfluencerEntity_NewInfluencerEntity(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := &InfluencerEntity{
		BaseEntity: BaseEntity{
			ID:        1,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Name: "Elon Musk",
	}
	actual := NewInfluencerEntity(1, "Elon Musk", createdAt, updatedAt)

	assert.Equal(t, expected, actual)
}

func TestInfluencerEntity_TableName(t *testing.T) {
	expected := "influencers"
	actual := NewInfluencerEntity(1, "Elon Musk", time.Now(), time.Now()).TableName()

	assert.Equal(t, expected, actual)
}

func TestInfluencerEntity_ToDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := dtos.NewInfluencerDTO(1, "Elon Musk", createdAt, updatedAt)
	actual := NewInfluencerEntity(1, "Elon Musk", createdAt, updatedAt).ToDTO()

	assert.Equal(t, expected, actual)
}

func TestInfluencerEntity_FromDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := NewInfluencerEntity(1, "Elon Musk", createdAt, updatedAt)
	actual := NewInfluencerEntity(0, "", time.Now(), time.Now())
	actual.FromDTO(*dtos.NewInfluencerDTO(1, "Elon Musk", createdAt, updatedAt))

	assert.Equal(t, expected, actual)
}

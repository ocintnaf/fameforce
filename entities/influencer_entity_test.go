package entities

import (
	"testing"
	"time"

	"github.com/ocintnaf/fameforce/dtos"
)

func TestNewInfluencerEntity(t *testing.T) {
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

	if actual.ID != expected.ID {
		t.Errorf("NewInfluencerEntity() ID = %d, want %d", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("NewInfluencerEntity() Name = %s, want %s", actual.Name, expected.Name)
	}

	if actual.CreatedAt != expected.CreatedAt {
		t.Errorf("NewInfluencerEntity() CreatedAt = %s, want %s", actual.CreatedAt, expected.CreatedAt)
	}

	if actual.UpdatedAt != expected.UpdatedAt {
		t.Errorf("NewInfluencerEntity() UpdatedAt = %s, want %s", actual.UpdatedAt, expected.UpdatedAt)
	}
}

func TestTableName(t *testing.T) {
	expected := "influencers"
	actual := NewInfluencerEntity(1, "Elon Musk", time.Now(), time.Now()).TableName()

	if actual != expected {
		t.Errorf("TableName() = %s, want %s", actual, expected)
	}
}

func TestToDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := dtos.NewInfluencerDTO(1, "Elon Musk", createdAt, updatedAt)
	actual := NewInfluencerEntity(1, "Elon Musk", createdAt, updatedAt).ToDTO()

	if actual.ID != expected.ID {
		t.Errorf("ToDTO() ID = %d, want %d", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("ToDTO() Name = %s, want %s", actual.Name, expected.Name)
	}
}

func TestFromDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := NewInfluencerEntity(1, "Elon Musk", createdAt, updatedAt)
	actual := NewInfluencerEntity(0, "", time.Now(), time.Now())
	actual.FromDTO(*dtos.NewInfluencerDTO(1, "Elon Musk", createdAt, updatedAt))

	if actual.ID != expected.ID {
		t.Errorf("FromDTO() ID = %d, want %d", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("FromDTO() Name = %s, want %s", actual.Name, expected.Name)
	}
}

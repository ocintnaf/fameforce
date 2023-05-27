package entities

import (
	"testing"

	"github.com/ocintnaf/fameforce/dtos"
)

func TestNewInfluencerEntity(t *testing.T) {
	expected := &InfluencerEntity{
		ID:   1,
		Name: "Elon Musk",
	}
	actual := NewInfluencerEntity(1, "Elon Musk")

	if actual.ID != expected.ID {
		t.Errorf("NewInfluencerEntity() ID = %d, want %d", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("NewInfluencerEntity() Name = %s, want %s", actual.Name, expected.Name)
	}
}

func TestTableName(t *testing.T) {
	expected := "influencers"
	actual := NewInfluencerEntity(1, "Elon Musk").TableName()

	if actual != expected {
		t.Errorf("TableName() = %s, want %s", actual, expected)
	}
}

func TestToDTO(t *testing.T) {
	expected := dtos.NewInfluencerDTO(1, "Elon Musk")
	actual := NewInfluencerEntity(1, "Elon Musk").ToDTO()

	if actual.ID != expected.ID {
		t.Errorf("ToDTO() ID = %d, want %d", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("ToDTO() Name = %s, want %s", actual.Name, expected.Name)
	}
}

func TestFromDTO(t *testing.T) {
	expected := NewInfluencerEntity(1, "Elon Musk")
	actual := NewInfluencerEntity(0, "")
	actual.FromDTO(dtos.NewInfluencerDTO(1, "Elon Musk"))

	if actual.ID != expected.ID {
		t.Errorf("FromDTO() ID = %d, want %d", actual.ID, expected.ID)
	}

	if actual.Name != expected.Name {
		t.Errorf("FromDTO() Name = %s, want %s", actual.Name, expected.Name)
	}
}

package entities

import (
	"testing"
	"time"

	"github.com/ocintnaf/fameforce/dtos"
	"github.com/ocintnaf/fameforce/types"
	"github.com/stretchr/testify/assert"
)

func TestUserEntity_NewUserEntity(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := &UserEntity{
		BaseEntity: BaseEntity[string]{
			ID:        "user-id",
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		},
		Email: "elon.musk@twitter.com",
		Type:  types.UserTypeInfluencer,
	}
	actual := NewUserEntity(
		"user-id",
		"elon.musk@twitter.com",
		types.UserTypeInfluencer,
		createdAt,
		updatedAt,
	)

	assert.Equal(t, expected, actual)
}

func TestUserEntity_TableName(t *testing.T) {
	expected := "users"
	actual := NewUserEntity(
		"user-id",
		"elon.musk@twitter.com",
		types.UserTypeInfluencer,
		time.Now(),
		time.Now(),
	).TableName()

	assert.Equal(t, expected, actual)
}

func TestUserEntity_ToDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := dtos.NewUserDTO("user-id", "elon.musk@twitter.com", types.UserTypeInfluencer, createdAt, updatedAt)
	actual := NewUserEntity("user-id", "elon.musk@twitter.com", types.UserTypeInfluencer, createdAt, updatedAt).ToDTO()

	assert.Equal(t, expected, actual)
}

func TestUserEntity_FromDTO(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour * 24)

	expected := NewUserEntity("user-id", "elon.musk@twitter.com", types.UserTypeInfluencer, createdAt, updatedAt)
	actual := NewUserEntity("", "", "", time.Now(), time.Now())
	actual.FromDTO(*dtos.NewUserDTO("user-id", "elon.musk@twitter.com", types.UserTypeInfluencer, createdAt, updatedAt))

	assert.Equal(t, expected, actual)
}

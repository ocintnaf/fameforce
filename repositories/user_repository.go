package repositories

import (
	"github.com/ocintnaf/fameforce/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindByID(id string) (*entities.UserEntity, error)
	Create(e *entities.UserEntity) (*entities.UserEntity, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindByID(id string) (*entities.UserEntity, error) {
	user := &entities.UserEntity{}

	err := ur.db.Limit(1).Find(user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ir *userRepository) Create(e *entities.UserEntity) (*entities.UserEntity, error) {
	result := ir.db.Create(e)
	return e, result.Error
}

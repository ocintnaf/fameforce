package repositories

import (
	"github.com/ocintnaf/fameforce/entities"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll() ([]entities.UserEntity, error)
	Save(e entities.UserEntity) (entities.UserEntity, error)
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindAll() ([]entities.UserEntity, error) {
	var users []entities.UserEntity

	ur.db.Find(&users)

	return users, nil
}

func (ir *userRepository) Save(e entities.UserEntity) (entities.UserEntity, error) {
	result := ir.db.Save(&e)
	return e, result.Error
}

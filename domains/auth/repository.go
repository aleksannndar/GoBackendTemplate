package auth

import (
	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(db *gorm.DB, user *User) error
	FindUserByEmail(db *gorm.DB, email string) *User
}

type DBAuthRepository struct {
}

func NewDBAuthRepository(db *gorm.DB) AuthRepository {
	return &DBAuthRepository{}
}

func (r *DBAuthRepository) Save(db *gorm.DB, user *User) error {
	userEntity := user.ToEntity()
	return db.Create(userEntity).Error
}

func (r *DBAuthRepository) FindUserByEmail(db *gorm.DB, email string) *User {
	var userEntity UserEntity
	result := db.First(&userEntity, UserEntity{Email: email})

	if result.Error != nil {
		return nil
	}

	return userEntity.ToDomain()
}

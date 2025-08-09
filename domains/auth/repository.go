package auth

import (
	"GoBackendTemplate/database"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Save(user *User) error
	SaveTransactional(tx *gorm.DB, user *User) error
	FindUserByEmail(email string) *User
	FindUserByEmailTransactional(tx *gorm.DB, email string) *User
}

type DBAuthRepository struct {
	db *gorm.DB
}

func NewDBAuthRepository(db *gorm.DB) AuthRepository {
	return &DBAuthRepository{
		db: db,
	}
}

func (r *DBAuthRepository) Save(user *User) error {
	return r.SaveTransactional(nil, user)
}

func (r *DBAuthRepository) SaveTransactional(tx *gorm.DB, user *User) error {
	db := database.TxOrDB(tx, r.db)
	userEntity := user.ToEntity()
	return db.Create(userEntity).Error
}

func (r *DBAuthRepository) FindUserByEmail(email string) *User {
	return r.FindUserByEmailTransactional(nil, email)
}

func (r *DBAuthRepository) FindUserByEmailTransactional(tx *gorm.DB, email string) *User {
	db := database.TxOrDB(tx, r.db)
	var userEntity UserEntity
	result := db.First(&userEntity, UserEntity{Email: email})

	if result.Error != nil {
		return nil
	}

	return userEntity.ToDomain()
}

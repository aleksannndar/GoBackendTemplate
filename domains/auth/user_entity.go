package auth

import "time"

type UserEntity struct {
	Id        string     `gorm:"column:id;primaryKey"`
	Email     string     `gorm:"column:email;not null;unique"`
	Password  *string    `gorm:"column:password"`
	CreatedAt *time.Time `gorm:"column:created_at;not null"`
}

func (UserEntity) TableName() string {
	return "users"
}

func (u *UserEntity) ToDomain() *User {
	return &User{
		Id:       u.Id,
		Email:    u.Email,
		Password: u.Password,
	}
}

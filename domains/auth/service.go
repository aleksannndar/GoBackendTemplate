package auth

import (
	"GoBackendTemplate/domains/authtoken"
	"fmt"

	"gorm.io/gorm"
)

type IAuthService interface {
	LoginWithPassword(email string, password string) (*LoginResponse, error)
	RegisterWithPassword(email string, password string) (*RegisterResponse, error)
}

type AuthService struct {
	db             *gorm.DB
	authRepository AuthRepository
	jwtService     authtoken.IJWTService
}

func NewAuthService(db *gorm.DB, jwtService authtoken.IJWTService) IAuthService {
	authRepository := NewDBAuthRepository(db)

	return &AuthService{
		db:             db,
		authRepository: authRepository,
		jwtService:     jwtService,
	}
}

type LoginResponse struct {
	UserId string
	Email  string
	Token  string
}

func (s *AuthService) LoginWithPassword(email string, password string) (*LoginResponse, error) {
	user := s.authRepository.FindUserByEmail(email)

	if user == nil {
		return nil, fmt.Errorf("failed to login: no user found for email")
	}

	if user.Password == nil || !CheckHashPassword(*user.Password, password) {
		return nil, fmt.Errorf("failed to login: wrong password")
	}

	token, err := s.jwtService.GenerateJWT(user.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %v", err)
	}

	return &LoginResponse{UserId: user.Id, Email: user.Email, Token: token}, nil
}

type RegisterResponse struct {
	UserId string
	Token  string
}

func (s *AuthService) RegisterWithPassword(email string, password string) (*RegisterResponse, error) {
	var createdUser *User
	err := s.db.Transaction(func(tx *gorm.DB) error {
		user := s.authRepository.FindUserByEmailTransactional(tx, email)
		if user != nil {
			return fmt.Errorf("email already exists")
		}

		user, err := CreateNewuser(email, password)
		if err != nil {
			return err
		}

		err = s.authRepository.SaveTransactional(tx, user)
		if err != nil {
			return err
		}

		createdUser = s.authRepository.FindUserByEmailTransactional(tx, email)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to register: %v", err)
	}

	token, err := s.jwtService.GenerateJWT(createdUser.Id)

	if err != nil {
		return nil, fmt.Errorf("failed to register: %v", err)
	}

	return &RegisterResponse{
		UserId: createdUser.Id,
		Token:  token,
	}, nil
}

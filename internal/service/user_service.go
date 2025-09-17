package service

import (
	"errors"
	"food-delivery-api/constants/user_role"
	"food-delivery-api/internal/models"
	"food-delivery-api/internal/repository"
	"food-delivery-api/pkg/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) Register(name, email, password string, userRole user_role.UserRole) error {
	hash, _ := utils.HashPassword(password)

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hash,
		UserRole: userRole,
	}
	return s.repo.Create(user)
}

func (s *UserService) Login(email, password string) (string, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, _ := utils.GenerateToken(user.UserID, string(user.UserRole))

	return token, nil
}

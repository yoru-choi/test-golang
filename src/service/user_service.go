package service

import (
	"context"

	"test-golang/src/models"
	"test-golang/src/repository"
)

// UserService defines methods for user operations
type UserService interface {
	GetUsers(ctx context.Context) ([]*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) error
	UpdateUser(ctx context.Context, id string, user *models.User) error
	DeleteUser(ctx context.Context, id string) error
}

// userService implements UserService
type userService struct {
	userRepo repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}


func (s *userService) GetUsers(ctx context.Context) ([]*models.User, error) {
	return s.userRepo.GetUsers(ctx)
}

func (s *userService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}
func (s *userService) CreateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

func (s *userService) UpdateUser(ctx context.Context, id string, user *models.User) error {
	return s.userRepo.UpdateUser(ctx, id, user)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}

// service/user_service.go
package service

import (
	"context"
	"project/models"
	"project/repository"
)

// UserService 사용자 서비스 구조체
type UserService struct {
	userRepo *repository.UserRepository
}

// NewUserService 생성자
func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// CreateUser 새 사용자 생성
func (s *UserService) CreateUser(ctx context.Context, user *models.User) error {
	return s.userRepo.CreateUser(ctx, user)
}

// GetUserByID ID로 사용자 조회
func (s *UserService) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	return s.userRepo.GetUserByID(ctx, id)
}

// UpdateUser 사용자 정보 업데이트
func (s *UserService) UpdateUser(ctx context.Context, id string, user *models.User) error {
	return s.userRepo.UpdateUser(ctx, id, user)
}

// DeleteUser 사용자 삭제
func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	return s.userRepo.DeleteUser(ctx, id)
}

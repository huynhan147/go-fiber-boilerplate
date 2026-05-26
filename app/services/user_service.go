package services

import (
	"context"
	"errors"
	"fmt"
	"myapp/app/repositories"
	"myapp/models"
	"myapp/pkg/cache"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const userCacheTTL = 5 * time.Minute

type UserService interface {
	GetAll(page, limit int) ([]models.User, int64, error)
	GetByID(id uint) (*models.User, error)
	Create(name, email, password string) (*models.User, error)
	Update(id uint, name, email string) (*models.User, error)
	Delete(id uint) error
}

type userService struct {
	repo  repositories.UserRepository
	cache *cache.Cache
}

func NewUserService(repo repositories.UserRepository, cache *cache.Cache) UserService {
	return &userService{repo: repo, cache: cache}
}

func (s *userService) GetAll(page, limit int) ([]models.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 15
	}
	return s.repo.FindAll(page, limit)
}

func (s *userService) GetByID(id uint) (*models.User, error) {
	ctx := context.Background()
	key := fmt.Sprintf("user:%d", id)

	// Thử lấy từ cache trước
	var user models.User
	if err := s.cache.Get(ctx, key, &user); err == nil {
		return &user, nil
	}

	// Cache miss → lấy từ DB
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}

	// Lưu vào cache
	_ = s.cache.Set(ctx, key, u, userCacheTTL)
	return u, nil
}

func (s *userService) Create(name, email, password string) (*models.User, error) {
	existing, _ := s.repo.FindByEmail(email)
	if existing != nil && existing.ID > 0 {
		return nil, errors.New("email already exists")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &models.User{Name: name, Email: email, Password: string(hashed)}
	if err := s.repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Update(id uint, name, email string) (*models.User, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("user not found")
	}

	user.Name = name
	user.Email = email
	if err := s.repo.Update(user); err != nil {
		return nil, err
	}

	// Xóa cache sau khi update
	_ = s.cache.Delete(context.Background(), fmt.Sprintf("user:%d", id))
	return user, nil
}

func (s *userService) Delete(id uint) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}
	// Xóa cache sau khi delete
	_ = s.cache.Delete(context.Background(), fmt.Sprintf("user:%d", id))
	return nil
}

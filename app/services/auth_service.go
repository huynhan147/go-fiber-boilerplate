package services

import (
	"context"
	"errors"
	"fmt"
	"myapp/app/repositories"
	"myapp/models"
	"myapp/pkg/cache"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Login(email, password string) (string, *models.User, error)
	Logout(tokenStr string) error
	IsTokenBlacklisted(tokenStr string) bool
}

type authService struct {
	repo  repositories.UserRepository
	cache *cache.Cache
	cfg   *viper.Viper
}

func NewAuthService(repo repositories.UserRepository, cache *cache.Cache, cfg *viper.Viper) AuthService {
	return &authService{repo: repo, cache: cache, cfg: cfg}
}

func (s *authService) Login(email, password string) (string, *models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := s.generateToken(user)
	if err != nil {
		return "", nil, err
	}

	return token, user, nil
}

// Logout blacklist token trên Redis cho đến khi hết hạn
func (s *authService) Logout(tokenStr string) error {
	expireHours := s.cfg.GetInt("JWT_EXPIRE_HOURS")
	ttl := time.Duration(expireHours) * time.Hour
	key := fmt.Sprintf("blacklist:%s", tokenStr)
	return s.cache.SetString(context.Background(), key, "1", ttl)
}

func (s *authService) IsTokenBlacklisted(tokenStr string) bool {
	key := fmt.Sprintf("blacklist:%s", tokenStr)
	return s.cache.Exists(context.Background(), key)
}

func (s *authService) generateToken(user *models.User) (string, error) {
	claims := jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * time.Duration(s.cfg.GetInt("JWT_EXPIRE_HOURS"))).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.cfg.GetString("JWT_SECRET")))
}

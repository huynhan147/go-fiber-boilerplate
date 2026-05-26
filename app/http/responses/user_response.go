package responses

import (
	"myapp/models"
	"time"
)

// UserResponse — tương đương Laravel API Resource
// Chỉ expose field cần thiết, ẩn password
type UserResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserResponse(user *models.User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format(time.RFC3339),
		UpdatedAt: user.UpdatedAt.Format(time.RFC3339),
	}
}

// NewUserCollection — tương đương Laravel UserCollection
func NewUserCollection(users []models.User) []UserResponse {
	result := make([]UserResponse, len(users))
	for i, u := range users {
		result[i] = NewUserResponse(&u)
	}
	return result
}

// AuthResponse — trả về sau login thành công
type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

func NewAuthResponse(token string, user *models.User) AuthResponse {
	return AuthResponse{
		Token: token,
		User:  NewUserResponse(user),
	}
}

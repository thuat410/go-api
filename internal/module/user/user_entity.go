package user

import (
	"time"
)

// UserEntity ánh xạ 1-1 với bảng public.user_profiles
type UserEntity struct {
	// Dùng string cho UUID vì pgx hỗ trợ map tự động từ kiểu uuid của Postgres sang string của Go
	ID       string `json:"id"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`

	// Dùng con trỏ *string cho các trường cho phép NULL
	FullName  *string `json:"full_name"`
	AvatarURL *string `json:"avatar_url"`

	Source    string    `json:"source"`
	CreatedAt time.Time `json:"created_at"`
}

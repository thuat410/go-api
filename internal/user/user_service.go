package user

import (
	"errors"
)

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserService struct {
	// Sau này bạn bơm Repository hoặc Supabase Client vào đây
}

func CreateUserService() *UserService {
	return &UserService{}
}

func (s *UserService) FindUserById(id string) (*UserDTO, error) {
	if id == "" {
		return nil, errors.New("ID không được để trống")
	}
	if id == "404" {
		return nil, errors.New("không tìm thấy người dùng")
	}
	user := &UserDTO{
		ID:    id,
		Name:  "Thuat Nguyen",
		Email: "thuat.nguyen@example.com",
	}

	return user, nil
}

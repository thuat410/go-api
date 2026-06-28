package user

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

// UserDTO là dữ liệu sạch sẽ, không chứa con trỏ để trả về cho Frontend
type UserDTO struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	UserName  string `json:"user_name"`
	FullName  string `json:"full_name"`
	AvatarURL string `json:"avatar_url"`
}

type UserService struct {
	userRepo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{userRepo: repo}
}

// ID đầu vào giờ là string
func (s *UserService) GetUserByID(ctx context.Context, id string) (*UserDTO, error) {
	if id == "" {
		return nil, errors.New("ID can not be blank")
	}
	entity, err := s.userRepo.FindByID(ctx, id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errors.New("Can not find user profile")
		}
		return nil, errors.New("Internal server error")
	}

	// Xử lý an toàn cho các trường NULL từ Database
	// Nếu giá trị trong DB là NULL, con trỏ sẽ là nil -> gán chuỗi rỗng ""
	fullName := ""
	if entity.FullName != nil {
		fullName = *entity.FullName
	}

	avatarURL := ""
	if entity.AvatarURL != nil {
		avatarURL = *entity.AvatarURL
	}

	// Map Entity sang DTO
	dto := &UserDTO{
		ID:        entity.ID,
		Email:     entity.Email,
		UserName:  entity.UserName,
		FullName:  fullName,
		AvatarURL: avatarURL,
	}

	return dto, nil
}

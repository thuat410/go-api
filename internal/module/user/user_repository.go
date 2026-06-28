package user

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

// Chú ý: ID bây giờ là kiểu string (chứa chuỗi UUID)
func (repo *UserRepository) FindByID(ctx context.Context, id string) (*UserEntity, error) {
	var u UserEntity

	// Cập nhật câu query quét đúng bảng user_profiles
	query := `
		SELECT
			id, email, user_name, full_name, avatar_url, source, created_at
		FROM user_profiles
		WHERE id = $1
	`

	// Hứng dữ liệu vào các trường mới
	err := repo.pool.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.Email,
		&u.UserName,
		&u.FullName,
		&u.AvatarURL,
		&u.Source,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}

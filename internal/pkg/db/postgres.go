package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

// InitDB trả về một Connection Pool của pgx
func InitDB() *pgxpool.Pool {
	dsn := os.Getenv("BD_URL")
	if dsn == "" {
		log.Fatal("❌ DB_URL is not defined in file .env")
	}

	// pgx yêu cầu Context cho hầu hết các thao tác
	ctx := context.Background()

	// Khởi tạo Connection Pool
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("❌ Failed to create connection pool to Supabase: %v", err)
	}

	// Ping thử để kiểm tra kết nối mạng
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("❌ Failed to ping Supabase: %v", err)
	}

	log.Println("🔌 Connected to Supabase via PGX successfully!")
	return pool
}

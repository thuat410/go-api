package main

import (
	"log"
	"os"

	"go-api/internal/product"
	"go-api/internal/user"
	"go-api/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type HealthResponse struct {
	Message string `json:"message"`
	Author  string `json:"author"`
	Health  string `json:"health"`
}

func checkHealth(c *gin.Context) {
	data := HealthResponse{
		Message: "Check health",
		Author:  "Thuat Nguyen",
		Health:  "APIs is running.",
	}
	utils.OKResponse(c, data)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Can not find .env, system will use environment variables from OS")
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port
	r := gin.Default()
	r.GET("/health", checkHealth)
	v1 := r.Group("/api/v1")
	user.RegisterRoutes(v1)
	product.RegisterRoutes(v1)

	utils.PrintSuccessBanner(port)
	err = r.Run(addr)
	if err != nil {
		log.Fatalf("❌ Không thể khởi động Server: %v", err)
	}
}

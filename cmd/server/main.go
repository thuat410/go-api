package main

import (
	"encoding/json"
	"fmt"
	"net/http" // Thư viện chuẩn để làm web server (giống Servlet/Tomcat)
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Thiết lập Header giống như trong Spring
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"message": "Check Health",
		"author":  "Thuat Nguyen",
		"tech":    "Golang",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	// Định nghĩa route: Khi vào đường dẫn "/" thì gọi hàm helloHandler
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server đang chạy tại http://localhost:8080...")

	// Bắt đầu lắng nghe tại cổng 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Lỗi server: %v\n", err)
	}
}

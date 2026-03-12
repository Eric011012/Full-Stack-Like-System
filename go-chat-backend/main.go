package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"

	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	// 1. 连接数据库 (请确保密码正确)
	connStr := "user=postgres password=mysecretpassword dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// 2. 自动初始化数据库结构
	initSQL, err := os.ReadFile("db/init.sql")
	if err != nil {
		log.Printf("警告: 未能读取 init.sql 文件: %v", err)
	} else {
		_, err = db.Exec(string(initSQL))
		if err != nil {
			log.Fatalf("初始化数据库失败: %v", err)
		}
		fmt.Println("数据库初始化成功")
	}

	// 3. API 接口逻辑
	http.HandleFunc("/api/like", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		// 1. 如果是 POST，插入一条新纪录
		if r.Method == http.MethodPost {
			_, err := db.Exec("INSERT INTO like_history DEFAULT VALUES")
			if err != nil {
				// 这里如果报错，通常是因为 like_history 表还没建
				log.Printf("写入失败: %v", err)
				http.Error(w, err.Error(), 500)
				return
			}
		}

		// 2. 无论 GET 还是 POST，最后都统计总行数返回给前端
		var totalLikes int
		err := db.QueryRow("SELECT COUNT(*) FROM like_history").Scan(&totalLikes)
		if err != nil {
			log.Printf("查询失败: %v", err)
			totalLikes = 0
		}

		// 3. 返回统计后的数字
		json.NewEncoder(w).Encode(map[string]int{"count": totalLikes})
	})

	http.HandleFunc("/api/reset", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")

		if r.Method == http.MethodPost {
			// 使用 TRUNCATE 快速清空表，并重置自增 ID
			_, err := db.Exec("TRUNCATE TABLE like_history RESTART IDENTITY")
			if err != nil {
				http.Error(w, err.Error(), 500)
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Reset successful"})
		}
	})
	fmt.Println("Server starting at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

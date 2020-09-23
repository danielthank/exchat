package infra

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlHandler struct {
	*gorm.DB
}

func NewSqlHandler() *SqlHandler {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to init db")
	}

	return &SqlHandler{db}
}

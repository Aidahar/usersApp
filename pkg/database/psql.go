package database

import (
	"fmt"

	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// type ConnectionInfo struct {
// 	Host     string
// 	Port     int
// 	Username string
// 	DBName   string
// 	SSLMode  string
// 	Password string
// }

func ConnectionDB(dsn string) (*gorm.DB, error) {

	fmt.Println("conname is\t\t", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatal(err)
	}

	db.AutoMigrate(&domain.User{})

	return db, nil
}

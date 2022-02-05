package database

import (
	"fmt"

	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectionDBSlite(dsn string) (*gorm.DB, error) {

	fmt.Println("conname is\t\t", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		logrus.Fatal(err)
	}

	db.AutoMigrate(&domain.User{})

	return db, nil
}

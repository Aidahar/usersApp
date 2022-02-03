package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Aidahar/filmsApi/internal/domain"
	"github.com/Aidahar/filmsApi/internal/repository/psql"
	"github.com/Aidahar/filmsApi/internal/service"
	transport "github.com/Aidahar/filmsApi/internal/transport/echo"
	"github.com/Aidahar/filmsApi/pkg/database"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	err := godotenv.Load(".env")
	db_driver := os.Getenv("POSTGRES_DRIVER")
	db_host := os.Getenv("POSTGRES_HOST")
	db_port := os.Getenv("POSTGRES_PORT")
	db_name := os.Getenv("POSTGRES_NAME")
	db_user := os.Getenv("POSTGRES_USER")
	db_password := os.Getenv("POSTGRES_PASSWORD")

	if err != nil {
		fmt.Println("cannot connetc to database", db_driver)
		log.Fatal("connect error:", err)
	}

	dsn := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		db_host, db_port, db_user, db_name, db_password)

	db, err := database.ConnectionDB(dsn)

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s\n", err.Error())
	}

	db.AutoMigrate(&domain.User{})

	repos := psql.NewUserRepository(db)
	service := service.NewUsersService(repos)
	handlers := transport.NewHandler(*service)
	e := handlers.InitRoutes()
	e.Logger.Fatal(e.Start(":8000"))
}

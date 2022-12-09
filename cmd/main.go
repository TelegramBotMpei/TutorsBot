package main

import (
	"context"
	"fmt"
	"github.com/TelegramBotMPEI/internal/models"
	"github.com/TelegramBotMPEI/internal/repository"
	"log"
	"os"
)

func main() {
	postgresUSER := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=%s", "localhost", 5432, postgresUSER, postgresPassword, postgresDB, "disable")
	connDB, err := repository.New(context.Background(), connString)
	if err != nil {
		log.Fatal("cannot connect to DB")
	}

	student := models.Student{
		ID:                 1,
		Name:               "Anton",
		Surname:            "Zotov",
		Course:             3,
		ScientificDirector: "smth smth",
		CourseWorkSubject:  "something",
		DiplomaWorkSubject: "something",
	}

	_, err = connDB.CreateStudent(context.TODO(), &student)
	if err != nil {
		log.Fatalf("cannot create user: %v", err.Error())
	}
}

package main

import (
	db "forumv2/db/SQLite3"
	"forumv2/internal/handler"
	"forumv2/internal/repository"
	"forumv2/internal/service"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// statment, _ := db.Database()
	db, err := db.Database()
	if err != nil {
		log.Println(err)
	}
	defer db.Close()

	repo := repository.NewRepository(db)
	service := service.NewService(repo)
	handler := handler.NewHandler(service)

	handler.InitRoutes()
}

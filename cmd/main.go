package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/user"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/pkg/boostrap"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/pkg/handler"
	"github.com/joho/godotenv"
)

var PORT = fmt.Sprintf(":%s", os.Getenv("PORT"))

func main() {
	godotenv.Load("../.env")

	server := http.NewServeMux()

	db, err := boostrap.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	logger := boostrap.NewLogger()
	userRepository := user.NewRepository(db, logger)
	userService := user.NewService(logger, userRepository)

	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, userService))

	fmt.Println("Server listening on 8080")

	log.Fatal(http.ListenAndServe(PORT, server))
}

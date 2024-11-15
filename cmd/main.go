package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/user"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/pkg/boostrap"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/pkg/handler"
)

func main() {
	server := http.NewServeMux()

	db := boostrap.NewDB()
	logger := boostrap.NewLogger()

	userRepository := user.NewRepository(db, logger)
	userService := user.NewService(logger, userRepository)

	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, userService))

	fmt.Println("Server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}

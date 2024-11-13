package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/user"
)

func main() {
	server := http.NewServeMux()

	db := user.DB{
		Users: []domain.User{
			{
				ID:        1,
				FirstName: "lucho",
				LastName:  "nicolosi",
				Email:     "nicolosi@gmail.com",
			},
			{
				ID:        3,
				FirstName: "lucho2",
				LastName:  "nicolosi",
				Email:     "nicolosi2@gmail.com",
			},
			{
				ID:        4,
				FirstName: "lucho3",
				LastName:  "nicolosi",
				Email:     "nicolosi3@gmail.com",
			},
		},
		MaxUserID: 3,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	userRepository := user.NewRepository(db, logger)
	userService := user.NewService(logger, userRepository)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, userService))

	fmt.Println("Server listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}

package boostrap

import (
	"log"
	"os"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
func NewDB() user.DB {
	return user.DB{
		Users: []domain.User{
			{
				ID:        1,
				FirstName: "lucho",
				LastName:  "nicolosi",
				Email:     "nicolosi@gmail.com",
			},
			{
				ID:        2,
				FirstName: "lucho2",
				LastName:  "nicolosi",
				Email:     "nicolosi2@gmail.com",
			},
			{
				ID:        3,
				FirstName: "lucho3",
				LastName:  "nicolosi",
				Email:     "nicolosi3@gmail.com",
			},
		},
		MaxUserID: 3,
	}
}

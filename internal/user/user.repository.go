package user

import (
	"context"
	"log"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
)

type DB struct {
	Users     []domain.User `json:"users"`
	MaxUserID uint64        `json:"max_user_id"`
}

type (
	UserRepository interface {
		Create(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	userRepository struct {
		db  DB
		log *log.Logger
	}
)

func NewRepository(db DB, logger *log.Logger) UserRepository {
	return &userRepository{
		db:  db,
		log: logger,
	}
}

func (r *userRepository) Create(ctx context.Context, user *domain.User) error {
	r.db.MaxUserID++
	user.ID = r.db.MaxUserID
	r.db.Users = append(r.db.Users, *user)
	r.log.Println("Repository create")
	return nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("Repository get all")
	return r.db.Users, nil
}

package user

import (
	"context"
	"errors"
	"log"
	"slices"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
)

type DB struct {
	Users     []domain.User
	MaxUserID uint64
}

type (
	UserRepository interface {
		Create(ctx context.Context, user *domain.User) error
		Update(ctx context.Context, user *domain.User) error
		GetAll(ctx context.Context) ([]domain.User, error)
		GetById(ctx context.Context, id uint64) (*domain.User, error)
		Delete(ctx context.Context, id uint64) (string, error)
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
func (r *userRepository) Update(ctx context.Context, user *domain.User) error {
	index := slices.IndexFunc(r.db.Users, func(u domain.User) bool {
		return u.ID == user.ID
	})

	r.db.Users[index] = *user

	r.log.Println("Repository update")
	return nil
}

func (r *userRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	r.log.Println("Repository get all")
	return r.db.Users, nil
}

func (r *userRepository) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	index := slices.IndexFunc(r.db.Users, func(u domain.User) bool {
		return u.ID == id
	})
	if index == -1 {
		return nil, errors.New("user not found")
	}
	return &r.db.Users[index], nil
}
func (r *userRepository) Delete(ctx context.Context, id uint64) (string, error) {
	users := slices.DeleteFunc(r.db.Users, func(u domain.User) bool {
		return u.ID == id
	})
	r.db.Users = users
	return "User deleted", nil
}

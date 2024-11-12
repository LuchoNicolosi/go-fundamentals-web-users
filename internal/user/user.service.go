package user

import (
	"context"
	"log"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
)

type (
	UserService interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
	}

	userService struct {
		log            *log.Logger
		userRepository UserRepository
	}
)

func NewService(log *log.Logger, repo UserRepository) UserService {
	return &userService{
		log:            log,
		userRepository: repo,
	}
}

func (s userService) Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error) {
	user := &domain.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
	}
	if err := s.userRepository.Create(ctx, user); err != nil {
		return nil, err
	}
	s.log.Println("service create")
	return user, nil
}

func (s userService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	s.log.Println("service get all")
	return users, nil
}

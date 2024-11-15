package user

import (
	"context"
	"log"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
)

type (
	UserService interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		Update(ctx context.Context, id uint64, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		GetById(ctx context.Context, id uint64) (*domain.User, error)
		Delete(ctx context.Context, id uint64) (string, error)
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
func (s userService) Update(ctx context.Context, id uint64, firstName, lastName, email string) (*domain.User, error) {

	user, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if email != "" {
		user.Email = email
	}
	if firstName != "" {
		user.FirstName = firstName
	}
	if lastName != "" {
		user.LastName = lastName
	}

	err = s.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}
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

func (s userService) GetById(ctx context.Context, id uint64) (*domain.User, error) {
	user, err := s.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	s.log.Println("service get by id")
	return user, nil
}
func (s userService) Delete(ctx context.Context, id uint64) (string, error) {
	v, err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return "", err
	}
	s.log.Println("service delete")
	return v, nil
}

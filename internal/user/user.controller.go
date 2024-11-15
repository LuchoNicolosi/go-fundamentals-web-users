package user

import (
	"context"
	"errors"
	"fmt"
)

type (
	UserController func(ctx context.Context, data interface{}) (interface{}, error)
	Endpoints      struct {
		Create  UserController
		GetAll  UserController
		GetById UserController
		Update  UserController
		Delete  UserController
	}

	GetReq struct {
		UserID uint64
	}
	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	UpdateRequest struct {
		UserID    uint64
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
	DeleteReq struct {
		UserID uint64
	}
)

func MakeEndpoints(ctx context.Context, service UserService) Endpoints {
	return Endpoints{
		Create:  makeCreateEndpoint(service),
		GetAll:  makeGetAllEndpoint(service),
		GetById: makeGetByIdEndpoint(service),
		Update:  makeUpdateEndpoint(service),
		Delete:  makeDeleteEndpoint(service),
	}
}

func makeGetAllEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		users, err := service.GetAll(ctx)
		if err != nil {
			return nil, err
		}
		return users, nil
	}
}
func makeGetByIdEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		result := data.(GetReq)
		fmt.Println(result.UserID)
		user, err := service.GetById(ctx, result.UserID)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
}
func makeDeleteEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		result := data.(DeleteReq)
		v, err := service.Delete(ctx, result.UserID)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func makeCreateEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		reqData := data.(CreateRequest)

		if reqData.FirstName == "" {
			return nil, errors.New("first name is required")
		}
		if reqData.LastName == "" {
			return nil, errors.New("last name is required")
		}
		if reqData.Email == "" {
			return nil, errors.New("email is required")
		}

		user, err := service.Create(ctx, reqData.FirstName, reqData.LastName, reqData.Email)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
func makeUpdateEndpoint(service UserService) UserController {
	return func(ctx context.Context, data interface{}) (interface{}, error) {
		reqData := data.(UpdateRequest)

		user, err := service.Update(ctx, reqData.UserID, reqData.FirstName, reqData.LastName, reqData.Email)
		if err != nil {
			return nil, err
		}

		return user, nil
	}
}

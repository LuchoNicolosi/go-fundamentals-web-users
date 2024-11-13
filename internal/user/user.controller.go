package user

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
)

type (
	UserController func(res http.ResponseWriter, req *http.Request)
	Endpoints      struct {
		Create UserController
		GetAll UserController
	}

	CreateRequest struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoints(ctx context.Context, service UserService) UserController {
	return func(res http.ResponseWriter, req *http.Request) {
		switch req.Method {
		case http.MethodGet:
			GetAllUsers(ctx, service, res)
		case http.MethodPost:
			var data CreateRequest
			err := json.NewDecoder(req.Body).Decode(&data)
			if err != nil {
				domain.MsgResponse(res, http.StatusBadRequest, err.Error())
				return
			}
			CreateUser(ctx, service, res, data)
		default:
			domain.InvalidMethodResponse(res)
			return
		}
	}
}

func GetAllUsers(ctx context.Context, service UserService, res http.ResponseWriter) {
	users, err := service.GetAll(ctx)
	if err != nil {
		domain.MsgResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(users)
	if err != nil {
		domain.MsgResponse(res, http.StatusInternalServerError, err.Error())
		return
	}
	domain.DataResponse(res, http.StatusOK, result)
}

func CreateUser(ctx context.Context, service UserService, res http.ResponseWriter, data interface{}) {
	reqData := data.(CreateRequest)

	if reqData.FirstName == "" {
		domain.MsgResponse(res, http.StatusBadRequest, "first name is required")
		return
	}
	if reqData.LastName == "" {
		domain.MsgResponse(res, http.StatusBadRequest, "last name is required")
		return
	}
	if reqData.Email == "" {
		domain.MsgResponse(res, http.StatusBadRequest, "email is required")
		return
	}

	user, err := service.Create(ctx, reqData.FirstName, reqData.LastName, reqData.Email)
	if err != nil {
		domain.MsgResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	result, err := json.Marshal(user)
	if err != nil {
		domain.MsgResponse(res, http.StatusInternalServerError, err.Error())
		return
	}

	domain.DataResponse(res, http.StatusCreated, result)
}

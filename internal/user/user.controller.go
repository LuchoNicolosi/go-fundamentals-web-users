package user

import (
	"context"
	"encoding/json"
	"net/http"
	"os/user"

	message "github.com/LuchoNicolosi/go-fundamentals-web-users/messages"
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
			var user user.User
			err := json.NewDecoder(req.Body).Decode(&user)
			if err != nil {
				message.MsgResponse(res, http.StatusBadRequest, err.Error())
				return
			}
			CreateUser(res, user)
		default:
			message.InvalidMethodResponse(res)
		}
	}
}

func GetAllUsers(res http.ResponseWriter, req *http.Request) {
	result, err := json.Marshal(users)
	if err != nil {
		message.MsgResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	message.DataResponse(res, http.StatusOK, result)
}

func CreateUser(res http.ResponseWriter, data interface{}) {
	userData := data.(user.User)

	if userData.FirstName == "" {
		message.MsgResponse(res, http.StatusBadRequest, "first name is required")
		return
	}
	if userData.LastName == "" {
		message.MsgResponse(res, http.StatusBadRequest, "last name is required")
		return
	}
	if userData.Email == "" {
		message.MsgResponse(res, http.StatusBadRequest, "email is required")
		return
	}

	maxId++
	userData.ID = maxId
	users = append(users, userData)

	result, err := json.Marshal(userData)
	if err != nil {
		message.MsgResponse(res, http.StatusBadRequest, err.Error())
		return
	}

	message.DataResponse(res, http.StatusCreated, result)
}

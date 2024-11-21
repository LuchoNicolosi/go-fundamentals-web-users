package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/user"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/pkg/transport"
	"github.com/LuchoNicolosi/go-web-response/response"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(res http.ResponseWriter, req *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {
		url := req.URL.Path

		path, pathSize := transport.Clean(url)

		params := make(map[string]string)

		if pathSize == 4 && path[2] != "" {

			params["userId"] = path[2]
		}

		params["token"] = req.Header.Get("Authorization")

		t := transport.New(res, req, context.WithValue(ctx, "params", params))

		var endpoint user.UserController
		var decode func(ctx context.Context, req *http.Request) (interface{}, error)

		log.Println(req.Method, ": ", req.URL)

		switch req.Method {
		case http.MethodGet:
			switch pathSize {
			case 3:
				endpoint = endpoints.GetAll
				decode = decodeGetAllUser
			case 4:
				endpoint = endpoints.GetById
				decode = decodeGetUser
			}
		case http.MethodPost:
			switch pathSize {
			case 3:
				endpoint = endpoints.Create
				decode = decodeCreateUser
			}
		case http.MethodPut:
			switch pathSize {
			case 4:
				endpoint = endpoints.Update
				decode = decodeUpdateUser
			}
		case http.MethodDelete:
			switch pathSize {
			case 4:
				endpoint = endpoints.Delete
				decode = decodeDeleteUser
			}
		}

		if endpoint != nil && decode != nil {
			t.Server(
				transport.Endpoint(endpoint),
				decode,
				encodeResponse,
				encodeError,
			)
		} else {
			domain.InvalidMethodResponse(res)
		}

	}
}

func decodeGetUser(ctx context.Context, req *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)

	id, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}
	return user.GetReq{
		UserID: id,
	}, nil
}

func decodeGetAllUser(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}
func decodeCreateUser(ctx context.Context, req *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	if err := tokenVerify(params["token"]); err != nil {
		return nil, response.Unauthorized(err.Error())
	}

	var data user.CreateRequest
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}
	return data, nil
}
func decodeUpdateUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var data user.UpdateRequest
	params := ctx.Value("params").(map[string]string)
	if err := tokenVerify(params["token"]); err != nil {
		return nil, response.Unauthorized(err.Error())
	}
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}

	data.UserID = userId

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: %v", err.Error()))
	}

	return data, nil
}

func decodeDeleteUser(ctx context.Context, req *http.Request) (interface{}, error) {
	params := ctx.Value("params").(map[string]string)
	if err := tokenVerify(params["token"]); err != nil {
		return nil, response.Unauthorized(err.Error())
	}
	userId, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		return nil, response.BadRequest(err.Error())
	}
	return user.DeleteReq{
		UserID: userId,
	}, nil
}
func encodeResponse(ctx context.Context, res http.ResponseWriter, data interface{}) error {

	resData := data.(response.Response)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	res.WriteHeader(resData.StatusCode())
	return json.NewEncoder(res).Encode(data)
}

func encodeError(ctx context.Context, err error, res http.ResponseWriter) {
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	errData := err.(response.Response)

	res.WriteHeader(errData.StatusCode())
	_ = json.NewEncoder(res).Encode(errData)
}

func tokenVerify(token string) error {
	if os.Getenv("TOKEN") != token {
		return errors.New("invalid token")
	}
	return nil
}

package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/domain"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/internal/user"
	"github.com/LuchoNicolosi/go-fundamentals-web-users/pkg/transport"
)

func NewUserHTTPServer(ctx context.Context, router *http.ServeMux, endpoints user.Endpoints) {
	router.HandleFunc("/users/", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(res http.ResponseWriter, req *http.Request) {

	return func(res http.ResponseWriter, req *http.Request) {
		url := req.URL.Path

		path, pathSize := transport.Clean(url)

		// params := make(map[string]string)
		var id string
		if pathSize == 4 && path[2] != "" {
			id = path[2]
			// params["userId"] = path[2]
		}

		t := transport.New(res, req, context.WithValue(ctx, "user_id", id))

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
	userId, err := strconv.ParseUint(ctx.Value("user_id").(string), 10, 64)
	if err != nil {
		return nil, err
	}
	return user.GetReq{
		UserID: userId,
	}, nil
}

func decodeGetAllUser(ctx context.Context, req *http.Request) (interface{}, error) {
	return nil, nil
}
func decodeCreateUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var data user.CreateRequest
	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}
	return data, nil
}
func decodeUpdateUser(ctx context.Context, req *http.Request) (interface{}, error) {
	var data user.UpdateRequest
	userId, err := strconv.ParseUint(ctx.Value("user_id").(string), 10, 64)
	if err != nil {
		return nil, err
	}

	data.UserID = userId

	if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("invalid request format: %v", err.Error())
	}

	return data, nil
}

func decodeDeleteUser(ctx context.Context, req *http.Request) (interface{}, error) {
	userId, err := strconv.ParseUint(ctx.Value("user_id").(string), 10, 64)
	if err != nil {
		return nil, err
	}
	return user.DeleteReq{
		UserID: userId,
	}, nil
}
func encodeResponse(ctx context.Context, res http.ResponseWriter, data interface{}) error {
	resultData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	status := http.StatusOK
	res.WriteHeader(status)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(res, `{"status": %d,"data":%s`, status, resultData)
	return nil
}

func encodeError(ctx context.Context, err error, res http.ResponseWriter) {
	status := http.StatusInternalServerError
	res.WriteHeader(status)
	res.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(res, `{"status": %d,"message":"%s"`, status, err.Error())
}

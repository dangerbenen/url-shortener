package auth

import (
	"fmt"
	"net/http"
	"olymp/configs"
	"olymp/pkg/req"
	"olymp/pkg/res"
)

type AuthHandlerDeps struct {
	*configs.Cfg
}

type AuthHandler struct {
	*configs.Cfg
}

func NewAuthHandler(router *http.ServeMux, deps AuthHandlerDeps) {
	handler := &AuthHandler{
		Cfg: deps.Cfg,
	}
	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (handler *AuthHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// читаем body
		body, err := req.HandleBody[LoginRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		//sendResp
		resp := LoginResponse{
			Token: "123",
		}
		res.Json(w, resp, 200)

	}
}

func (handler *AuthHandler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			return
		}
		fmt.Println(body)
		//sendResp
		resp := RegisterResponse{
			Token: "123",
		}
		res.Json(w, resp, 200)
	}
}

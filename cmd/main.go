package main

import (
	"fmt"
	"net/http"
	"olymp/configs"
	"olymp/internal/auth"
	"olymp/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()
	_ = db.NewDb(cfg)
	router := http.NewServeMux()

	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Cfg: cfg,
	})
	fmt.Println(cfg.Auth.Secret)
	server := &http.Server{
		Addr:    ":52",
		Handler: router,
	}

	fmt.Println("http://localhost:52")
	server.ListenAndServe()

}

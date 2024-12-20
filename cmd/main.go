package main

import (
	"fmt"
	"net/http"
	"olymp/configs"
	"olymp/internal/auth"
	"olymp/internal/link"
	"olymp/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()
	db := db.NewDb(cfg)
	router := http.NewServeMux()
	//Repos
	linkRepository := link.NewLinkRepository(db)

	//Handler
	auth.NewAuthHandler(router, auth.AuthHandlerDeps{
		Cfg: cfg,
	})
	link.NewLinkHandler(router, link.LinkHandlerDeps{
		LinkRepository: linkRepository,
	})

	server := &http.Server{
		Addr:    ":52",
		Handler: router,
	}

	fmt.Println("http://localhost:52")
	server.ListenAndServe()

}

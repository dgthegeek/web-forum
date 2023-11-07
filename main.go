package main

import (

	// internals "golang-rest-api-starter/internals/config/database"
	server "golang-rest-api-starter/internals/config/server"
	"golang-rest-api-starter/router"
	"golang-rest-api-starter/service/middleware"
)

func main() {
	router := &router.NewRouter{
		Middlewares: []router.Middleware{
			middleware.AuthMiddleware,
			middleware.LoggerMiddleware,
			middleware.DBMiddleware,
		},
	}

	// instantiate a new server
	server := server.Config{
		PORT:     ":8000",
		Hostname: "http://localhost",
	}
	server.Init(router)
}

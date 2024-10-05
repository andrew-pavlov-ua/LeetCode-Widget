package main

import (
	"cmd/internal/controllers"
	"cmd/internal/db"
	"cmd/internal/env"
	"cmd/internal/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Creating variables app need to connect to db with correct dsn and port
	var (
		dsn  = env.Must("POSTGRES_DSN")
		port = env.Must("PORT")
	)

	// Setting up connection with db
	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	// creating service that manipulates with users
	var userService = services.NewUserService(repository)

	r := gin.Default()
	r.LoadHTMLGlob("public/view/*.html")

	webController := controllers.NewWebController(userService)

	r.GET("/", webController.ReturnIndex)                                            // Returning main index.html
	r.GET("/api/slug/:leetcode_user_slug/badge.svg", webController.StatsBadgeBySlug) // Starting with badge creation

	r.
		//Icons
		Static("/favicon.ico", "./internal/templates/v1/img/LeetCodeLogo.png").
		Static("/LeetCodeLogo.png", "./internal/templates/v1/img/LeetCodeLogo.png")

	var serverErr = r.Run(port) // Running app on the port from .env
	if serverErr != nil {
		log.Fatalln(serverErr)

		return
	}
	fmt.Println("Server started on http://localhost:8080")
}

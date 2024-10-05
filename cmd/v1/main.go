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
	var (
		dsn  = env.Must("POSTGRES_DSN")
		port = env.Must("PORT")
	)

	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	var userService = services.NewUserService(repository)

	r := gin.Default()
	r.LoadHTMLGlob("public/view/*.html")

	webController := controllers.NewWebController(userService)

	r.GET("/", webController.ReturnIndex)

	r.GET("/redirect", webController.RedirectToLc)
	r.GET("/api/slug/:leetcode_user_slug/badge.svg", webController.StatsBadgeBySlug)

	r.
		//Icons
		Static("/favicon.ico", "./internal/templates/v1/img/LeetCodeLogo.png").
		Static("/LeetCodeLogo.png", "./internal/templates/v1/img/LeetCodeLogo.png")

	var serverErr = r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if serverErr != nil {
		log.Fatalln(serverErr)

		return
	}
	fmt.Println("Server started on http://localhost:8080")
}

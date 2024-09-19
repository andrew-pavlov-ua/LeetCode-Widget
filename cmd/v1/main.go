package main

import (
	"cmd/internal/controllers"
	"cmd/internal/db"
	"cmd/internal/env"
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

	r := gin.Default()
	r.LoadHTMLGlob("public/view/*.html")

	webController := controllers.NewWebController()

	r.GET("/", webController.ReturnIndex)

	r.GET("/redirect", webController.RedirectToLc)
	r.GET("/api/slug/:leetcode_user_slug/badge.svg", webController.StatsBadgeBySlug)

	var serverErr = r.Run(port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if serverErr != nil {
		log.Fatalln(serverErr)

		return
	}
	fmt.Println("Server started on http://localhost:8080")
}

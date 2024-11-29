package main

import (
	"fmt"
	"log"
	"os"

	"cmd/internal/controllers"
	"cmd/internal/db"
	"cmd/internal/env"
	"cmd/internal/server"
	"cmd/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Started program ")
	// Creating variables app need to connect to db with correct dsn and port
	var (
		dsn = env.Must("POSTGRES_DSN")
	)

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error while log file opening: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	log.Println("Starting app")

	// Setting up connection with db
	var pgConnection = db.MustConnection(dsn)
	defer pgConnection.Close()

	var repository = db.MustRepository(pgConnection)
	defer repository.Close()

	// creating services
	var userService = services.NewLcUserService(repository)
	var visitsService = services.NewVisistsStatsService(repository)

	r := gin.Default()
	r.LoadHTMLGlob("public/view/*.html")
	r.Static("/assets/images", "./public/assets/images")
	r.Static("/style", "./public/view/style")
	// r.Static("/.well-known/acme-challenge", "./.well-known/acme-challenge")

	webController := controllers.NewWebController(userService, visitsService)

	r.GET("/", webController.ReturnIndex)                                            // Returning main index.html
	r.GET("/api/slug/:leetcode_user_slug/badge.svg", webController.StatsBadgeBySlug) // Starting with badge creation
	r.GET("/:leetcode_user_slug/redirect", webController.VisitsCountRedirect)        // Processing profile view
	// r.GET("/lcb/api/cw/id/:cw_user_id/stats", webController.ReturnCWStatsById)

	server.Run(r.Handler())
}

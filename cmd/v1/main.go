package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cmd/internal/controllers"
	"cmd/internal/db"
	"cmd/internal/env"
	"cmd/internal/server"
	"cmd/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
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

	// redis
	opt, err := redis.ParseURL(env.Must("REDIS_OPT"))
	if err != nil {
		panic(err)
	}
	client := redis.NewClient(opt)

	// Ping to test the connection
	ctx := context.Background()
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Could not connect to Redis:", err)
	} else {
		fmt.Println("Redis connection successful:", pong)
	}

	// Example: Set and Get a key
	err = client.Set(ctx, "mykey", "myvalue", 0).Err() // 0 means no expiration
	if err != nil {
		fmt.Println("Error setting key:", err)
	}

	val, err := client.Get(ctx, "mykey").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
	} else {
		fmt.Println("Value of mykey:", val)
	}

	// creating services
	var userService = services.NewLcUserService(repository)
	var visitsService = services.NewVisistsStatsService(repository)

	r := gin.Default()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("public/view/*.html")
	// r.Static("/.well-known/acme-challenge", "./.well-known/acme-challenge")

	webController := controllers.NewWebController(userService, visitsService)

	r.GET("/", webController.ReturnIndex)                                         // Returning main index.html
	r.GET("/redirect-page/:leetcode_user_slug", webController.ReturnRedirectPage) // Redirect page

	r.GET("/api/slug/:leetcode_user_slug/badge.svg", webController.StatsBadgeBySlug) // Starting with badge creation
	r.GET("/:leetcode_user_slug/redirect", webController.VisitsCountRedirect)        // Processing profile view
	r.GET("/favicon.ico", func(c *gin.Context) { c.Status(http.StatusNoContent) })

	// CSS
	r.Static("/style", "./public/view/style")
	r.Static("/redirect-page/style", "./public/view/style")
	// Assets
	r.Static("/assets/images", "./public/assets/images")
	r.Static("/redirect-page/assets/images", "./public/assets/images")
	r.Static("/assets/js", "./public/assets/js")
	r.Static("/redirect-page/assets/js", "./public/assets/js")
	//Icons
	r.Static("/site_ico.ico", "./public/assets/images/site_ico.ico")
	r.Static("/redirect-page/site_ico.ico", "./public/assets/images/site_ico.ico")

	server.Run(r.Handler())

}

package main

import (
	"cmd/internal/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("public/view/*.html")

	webController := controllers.NewWebController()

	r.GET("/", webController.ReturnSimpleView)

	r.GET("/redirect", webController.RedirectToLc)
	r.GET("/api/:leetcode_user_slug/badge.svg", webController.StatsBadge)
	r.GET("/LeetCodeLogo", webController.ReturnSimpleView)

	err := r.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		return
	}
	fmt.Println("Server started on http://localhost:8080")
}

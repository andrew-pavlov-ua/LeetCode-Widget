package controllers

import (
	_ "cmd/internal/db"
	"cmd/internal/services"
	v1 "cmd/internal/templates/v1"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type WebController struct {
	userService *services.UserService
}

func NewWebController(userService *services.UserService) *WebController {
	return &WebController{
		userService: userService,
	}
}

func (c *WebController) ReturnIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func (c *WebController) StatsBadgeBySlug(ctx *gin.Context) {
	// Init userNotFound badge and getting userSlug (Leetcode id) from the url
	var badge []byte
	userSlug := ctx.Param("leetcode_user_slug")

	userData, err := c.userService.Upsert(ctx.Request.Context(), userSlug)

	if err != nil {
		fmt.Println("err 38: ", err)
	} else if userData == nil || userData.Rank == 0 {
		badge = []byte(v1.BadgeNoUserFound())
	} else {
		// Calculating bars width (px) in the badge
		barsWidth := v1.BarsWidth{
			EasyWidth:   c.CalculateWidth(userData.EasyCount, v1.EasyMaxValue),
			MediumWidth: c.CalculateWidth(userData.MediumCount, v1.MediumMaxValue),
			HardWidth:   c.CalculateWidth(userData.HardCount, v1.HardMaxValue)}

		badge = []byte(v1.Badge(*userData, barsWidth))
	}

	c.renderImage(ctx, badge)
}

func (c *WebController) renderImage(ctx *gin.Context, data []byte) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, "image/svg+xml", data)
}

func (c *WebController) CalculateWidth(count int64, max int64) float64 {
	countPerPixel := float64(v1.BarWidthValue) / float64(max)
	result := countPerPixel * float64(count)
	return result
}

func (c *WebController) GetLeetCodeLogo(ctx *gin.Context) {
	imgData, err := os.ReadFile("images/LeetCodeLogo.png")
	if err != nil {
		log.Println("Error reading image file: ", err)
		ctx.String(http.StatusInternalServerError, "Error reading image file")
		return
	}
	ctx.Data(http.StatusOK, "image/png", imgData)
}

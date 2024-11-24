package controllers

import (
	_ "cmd/internal/db"
	"cmd/internal/services"
	v1 "cmd/internal/templates/v1"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebController struct {
	userService   *services.UserService
	visitsService *services.VisitsStatsService
}

func NewWebController(userService *services.UserService, visitsService *services.VisitsStatsService) *WebController {
	return &WebController{
		userService:   userService,
		visitsService: visitsService}
}

func (c *WebController) ReturnIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func (c *WebController) StatsBadgeBySlug(ctx *gin.Context) {
	// Init userNotFound badge and getting userSlug (Leetcode id) from the url
	var badge []byte
	var visitStats v1.VisitsStats
	logo_path := "public/assets/images/logo_base64.txt"
	userSlug := ctx.Param("leetcode_user_slug")

	userData, err := c.userService.GetOrCreate(ctx.Request.Context(), userSlug)

	if err != nil {
		fmt.Println("StatsBadgeBySlug: error upserting stats ", err)
	} else if userData == nil || userData.Rank == 0 {
		badge = []byte(v1.BadgeNoUserFound())
	} else {
		// Calculating bars width (px) in the badge
		barsWidth := v1.BarsWidth{
			EasyWidth:   c.CalculateWidth(userData.EasyCount, v1.EasyMaxValue),
			MediumWidth: c.CalculateWidth(userData.MediumCount, v1.MediumMaxValue),
			HardWidth:   c.CalculateWidth(userData.HardCount, v1.HardMaxValue)}

		// Founding user's lc vivsits count
		visitStats, err = c.visitsService.GetFullStatsCount(ctx, userSlug)
		if err != nil {
			fmt.Println("StatsBadgeBySlug: error getting full count stats", err)
		}

		logo_base64 := services.ReadFile(logo_path)

		badge = []byte(v1.Badge(*userData, barsWidth, visitStats, logo_base64))
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

func (c *WebController) VisitsCountRedirect(ctx *gin.Context) {
	userSlug := ctx.Param("leetcode_user_slug")
	redirectUrl := fmt.Sprintf("https://leetcode.com/u/%s/", userSlug)

	// Adding 1 to current hour's profile visits
	err := c.visitsService.InsertCount(ctx, userSlug)
	if err != nil {
		fmt.Println("VisitsCountRedirect: error redirecting user", err)
	}
	ctx.Redirect(http.StatusFound, redirectUrl)
}

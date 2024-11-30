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
	userService   *services.LcUserService
	visitsService *services.VisitsStatsService
}

func NewWebController(userService *services.LcUserService, visitsService *services.VisitsStatsService) *WebController {
	return &WebController{
		userService:   userService,
		visitsService: visitsService}
}

func (c *WebController) ReturnIndex(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func (c *WebController) ReturnRedirectPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "redirect_page.html", gin.H{})
}

func (c *WebController) StatsBadgeBySlug(ctx *gin.Context) {
	// Init userNotFound badge and getting userSlug (Leetcode id) from the url
	var badge []byte
	// var visitStats v1.VisitsStats
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

		// // Founding user's lc vivsits count
		// user_id, err := c.userService.GetUserIdBySlug(ctx, userSlug)
		// if err != nil {
		// 	fmt.Println("StatsBadgeBySlug: error getting userId", err)
		// }

		// // visitStats, err = c.visitsService.GetFullStatsCount(ctx, user_id)
		// if err != nil {
		// 	fmt.Println("StatsBadgeBySlug: error getting full count stats", err)
		// }

		logo_base64 := services.ReadFile(logo_path)

		badge = []byte(v1.Badge(*userData, barsWidth, logo_base64))
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
	userId, err := c.userService.GetUserIdBySlug(ctx, userSlug)
	if err != nil {
		fmt.Println("VisitsCountRedirect: error getting userId", err)
	}

	redirectUrl := fmt.Sprintf("https://leetcode.com/u/%s/", userSlug)

	// Adding 1 to current hour's profile visits
	err = c.visitsService.InsertCount(ctx, userId)
	if err != nil {
		fmt.Println("VisitsCountRedirect: error redirecting user", err)
	}
	ctx.Redirect(http.StatusFound, redirectUrl)
}

// In development
// func (c *WebController) ReturnCWStatsById(ctx *gin.Context) {
// 	userId := ctx.Param("cw_user_id")
// 	data, err := codewars_api.GetUserProfile(userId)
// 	if err != nil {
// 		fmt.Printf("ReturnCWStatsById: err: %v", err)
// 	}

// 	ctx.Data(http.StatusOK, "application/json", []byte(services.FormatJSON(data)))
// }

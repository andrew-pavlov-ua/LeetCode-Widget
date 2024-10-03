package controllers

import (
	_ "cmd/internal/db"
	"cmd/internal/services"
	v1 "cmd/internal/templates/v1"
	"fmt"
	"github.com/gin-gonic/gin"
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

func (c *WebController) RedirectToLc(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "https://leetcode.com")
}

func (c *WebController) StatsBadgeBySlug(ctx *gin.Context) {
	userSlug := ctx.Param("leetcode_user_slug")

	userId := c.userService.Upsert(ctx.Request.Context(), userSlug)
	fmt.Printf("userId:%v\n", userId)

	userData, err := c.userService.GetByStatsById(ctx, userId)
	if err != nil || userData == nil {
		fmt.Println("err 38")
		fmt.Println(err)
		ctx.HTML(http.StatusInternalServerError, "error.html", gin.H{})
	}

	fmt.Printf("________userData:%v\n", userData)

	barsWidth := v1.BarsWidth{
		EasyWidth:   c.CalculateWidth(userData.EasyCount, v1.EasyMaxValue),
		MediumWidth: c.CalculateWidth(userData.MediumCount, v1.MediumMaxValue),
		HardWidth:   c.CalculateWidth(userData.HardCount, v1.HardMaxValue)}

	c.renderImage(ctx, []byte(v1.Badge(*userData, barsWidth)))
}

//func (c *WebController) StatsBadgeBySlug(ctx *gin.Context) {
//	userSlug := ctx.Param("leetcode_user_slug")
//
//	userData := leetcode_api.MatchedUserMapToUserProfile(userSlug)
//
//	lcUserData := v1.NewLcUserDataFromReq(*userData)
//
//	barsWidth := v1.BarsWidth{
//		EasyWidth:   c.CalculateWidth(lcUserData.EasyCount, v1.EasyMaxValue),
//		MediumWidth: c.CalculateWidth(lcUserData.MediumCount, v1.MediumMaxValue),
//		HardWidth:   c.CalculateWidth(lcUserData.HardCount, v1.HardMaxValue)}
//
//	c.renderImage(ctx, []byte(v1.Badge(*lcUserData, barsWidth)))
//}

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
		ctx.String(http.StatusInternalServerError, "Error reading image file")
		return
	}
	ctx.Data(http.StatusOK, "image/png", imgData)
}

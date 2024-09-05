package controllers

import (
	"cmd/internal/leetcode_api"
	v1 "cmd/internal/templates/v1"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type WebController struct {
}

func NewWebController() *WebController {
	return &WebController{}
}

func (c *WebController) ReturnSimpleView(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{})
}

func (c *WebController) RedirectToLc(ctx *gin.Context) {
	ctx.Redirect(http.StatusTemporaryRedirect, "https://leetcode.com")
}

func (c *WebController) StatsBadge(ctx *gin.Context) {

}

func (c *WebController) StatsBadgeBySLug(ctx *gin.Context) {
	userSlug := ctx.Param("leetcode_user_slug")

	userData := leetcode_api.MatchedUserMapToUserProfile(userSlug)

	lcUserData := v1.NewLcUserDataFromReq(*userData)

	barsWidth := v1.BarsWidth{
		EasyWidth:   c.CalculateWidth(lcUserData.EasyCount, v1.EasyMaxValue),
		MediumWidth: c.CalculateWidth(lcUserData.MediumCount, v1.MediumMaxValue),
		HardWidth:   c.CalculateWidth(lcUserData.HardCount, v1.HardMaxValue)}

	c.renderImage(ctx, []byte(v1.Badge(*lcUserData, barsWidth)))
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
		ctx.String(http.StatusInternalServerError, "Error reading image file")
		return
	}
	ctx.Data(http.StatusOK, "image/png", imgData)
}

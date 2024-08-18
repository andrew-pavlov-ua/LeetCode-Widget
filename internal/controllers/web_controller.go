package controllers

import (
	v1 "cmd/internal/templates/v1"
	"github.com/gin-gonic/gin"
	"net/http"
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
	lcStats := v1.LcStats{Username: "a", Rank: 1, Lvl: 2, Experience: 100, EasyCount: 300, MediumCount: 1000, HardCount: 5, TotalCount: 6}
	barsWidth := v1.BarsWidth{
		EasyWidth:   c.CalculateWidth(lcStats.EasyCount, v1.EasyMaxValue),
		MediumWidth: c.CalculateWidth(lcStats.MediumCount, v1.MediumMaxValue),
		HardWidth:   c.CalculateWidth(lcStats.HardCount, v1.HardMaxValue)}

	c.renderImage(ctx, []byte(v1.Badge(lcStats, barsWidth)))
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

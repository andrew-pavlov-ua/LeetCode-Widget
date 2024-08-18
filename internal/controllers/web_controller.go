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
	lcStats := v1.LcStats{Username: "a", Rank: 1, Lvl: 2, Experience: 100, EasyCount: 300, MediumCount: 4, HardCount: 5, TotalCount: 6}

	c.renderImage(ctx, []byte(v1.Badge(lcStats)))
}

func (c *WebController) renderImage(ctx *gin.Context, data []byte) {
	ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Header("Pragma", "no-cache")
	ctx.Header("Expires", "0")
	ctx.Data(http.StatusOK, "image/svg+xml", data)
}

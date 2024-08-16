package controllers

import (
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

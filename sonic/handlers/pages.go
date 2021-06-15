package handlers

import (
	"github.com/openware/pkg/sonic/models"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomarkdown/markdown"
)

// SetPageRoutes configure module HTTP routes
func SetPageRoutes(router *gin.Engine, ptr models.IPage) error {
	for _, p := range ptr.List() {
		router.GET(p.GetPath(), pageGet(p))
	}
	return nil
}

func pageGet(p models.IPage) func(c *gin.Context) {
	return func(c *gin.Context) {
		body := string(markdown.ToHTML([]byte(p.GetBody()), nil, nil))

		c.HTML(http.StatusOK, "page.html", gin.H{
			"title":       p.GetTitle(),
			"description": p.GetDescription(),
			"body":        template.HTML(body),
		})
	}
}

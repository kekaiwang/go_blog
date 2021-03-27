package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaiwang/go-blog/app/api"
	"github.com/kekaiwang/go-blog/router/middleware"
)

func SetupRouter(g *gin.Engine) {
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.OrderAccess)

	g.LoadHTMLGlob("web/*")

	g.HEAD("/ping", api.HealthPing)
	// 404 redirect
	g.NoRoute(func(c *gin.Context) {
		c.Header("Content-type", "text/html; charset=utf-8")

		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"Title": "Kekai Wang's blog",
		})
	})

	g.GET("/", api.GetIndexArticle)                 // index page
	g.GET("/article/:slug", api.GetArticleDetail)   // article detail
	g.GET("/page/:slug", api.PageInfo)              //page info
	g.GET("/categories/:link", api.GetCategoryList) //category list
	g.GET("/tags/:link", api.GetTagList)            //tag list

	admin := g.Group("/admin")
	{
		admin.POST("/login", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
		admin.POST("/create/article", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "success",
			})
		})
	}
}

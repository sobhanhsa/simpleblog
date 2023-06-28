package main

import (
	// "fmt"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sobhanhsa/simpleblog/controllers"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/initializers"
	"github.com/sobhanhsa/simpleblog/middlewares"
	// "github.com/sobhanhsa/simpleblog/rssdecoders"
	// "github.com/sobhanhsa/simpleblog/models"
)

func init() {
	initializers.LoudVars()
	db.ConnectDB()

	// db.PrementDelete()
	// db.ModelMigrate(models.Article{})
}
func main() {

	//backend router
	apiEngine := gin.New()

	apiG := apiEngine.Group("/api")

	// apiG.Use(middlewares.UserAuth)

	{
		apiG.GET("/main", controllers.MainPage)
		apiG.GET("/category/:category/", controllers.ArticleCategory)
		apiG.GET("/profile/:username", controllers.ShowProfile)
		apiG.GET("/result", controllers.SearchArticle)
		apiG.GET("/validate", controllers.UserValidate)
		apiG.POST("/signup", controllers.SignUpPage)
		apiG.POST("/login", controllers.LoginPage)
		apiG.GET("/articles", controllers.ShowArticles)
		apiG.POST("/article", controllers.PublishArticle)
		apiG.GET("/article/:id", controllers.ShowArticle)
		apiG.PUT("/article/:id", controllers.UpdateArticle)
		apiG.DELETE("/article/:id", controllers.DeleteArticle)
	}

	//front router

	fr := gin.Default()

	fr.Static("/", "./public")

	//main router
	r := gin.New()

	r.Use(middlewares.UserAuth)

	r.GET("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiEngine.HandleContext(c)
		} else {
			fr.HandleContext(c)
		}
	})

	r.POST("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiEngine.HandleContext(c)
		} else {
			fr.HandleContext(c)
		}
	})

	r.PUT("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiEngine.HandleContext(c)
		} else {
			fr.HandleContext(c)
		}
	})

	r.DELETE("/*any", func(c *gin.Context) {
		path := c.Param("any")
		if strings.HasPrefix(path, "/api") {
			apiEngine.HandleContext(c)
		} else {
			fr.HandleContext(c)
		}
	})

	r.Run()

}

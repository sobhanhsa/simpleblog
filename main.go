package main

import (
	// "fmt"

	"time"

	"github.com/gin-gonic/gin"
	"github.com/sobhanhsa/simpleblog/controllers"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/initializers"
	"github.com/sobhanhsa/simpleblog/middlewares"

	// "github.com/sobhanhsa/simpleblog/rssdecoders"
	// "github.com/sobhanhsa/simpleblog/models"
	"github.com/gin-contrib/cors"
)

func init() {
	initializers.LoudVars()
	db.ConnectDB()

	// db.PrementDelete()
	// db.ModelMigrate(models.Article{})
}
func main() {

	r := gin.Default()

	r.Use(middlewares.UserAuth)

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:8080"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.GET("/", controllers.MainPage)
	r.GET("/category/:category/", controllers.ArticleCategory)
	r.GET("/profile/:username", controllers.ShowProfile)
	r.GET("/result", controllers.SearchArticle)
	r.GET("/validate", controllers.UserValidate)
	r.POST("/signup", controllers.SignUpPage)
	r.POST("/login", controllers.LoginPage)
	r.GET("/articles", controllers.ShowArticles)
	r.POST("/article", controllers.PublishArticle)
	r.GET("/article/:id", controllers.ShowArticle)
	r.PUT("/article/:id", controllers.UpdateArticle)
	r.DELETE("/article/:id", controllers.DeleteArticle)

	r.Run()

}

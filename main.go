package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/sobhanhsa/simpleblog/controllers"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/initializers"
	"github.com/sobhanhsa/simpleblog/middlewares"
	// "github.com/sobhanhsa/simpleblog/models"
)

func init() {
	initializers.LoudVars()
	db.ConnectDB()
	// db.PrementDelete()
	// db.ModelMigrate(models.User{})
}
func main() {

	r := gin.Default()

	r.Use(middlewares.UserAuth)

	r.GET("/", controllers.MainPage)
	r.GET("/validate", controllers.UserValidate)
	r.POST("/signup", controllers.SignUpPage)
	r.POST("/login", controllers.LoginPage)
	r.GET("/articles", controllers.ShowArticles)
	r.POST("/article", controllers.PublishArticle)
	r.GET("/article/:title", controllers.ShowArticle)
	r.PUT("/article/:title", controllers.UpdateArticle)
	r.DELETE("/article/:title", controllers.DeleteArticle)

	r.Run()
}

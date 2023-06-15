package main

import (
	// "fmt"

	"github.com/gin-gonic/gin"
	"github.com/sobhanhsa/simpleblog/controllers"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/initializers"
	"github.com/sobhanhsa/simpleblog/middlewares"
	"github.com/sobhanhsa/simpleblog/models"
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

	r.GET("/ss", func(c *gin.Context) {
		var arts models.Article

		// for i := 0; i > len(arts); i++ {
		// 	var deletedTime time.Time = arts[i].DeletedAt.Time
		// 	var now time.Time = time.Now()

		// 	deletedTime.Add(24 * time.Hour)

		// 	if !now.Before(deletedTime) {
		// 		db.DB.Unscoped().Delete(&arts[i])
		// 	}
		// }

		c.JSON(200, gin.H{"data": arts})
	})
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

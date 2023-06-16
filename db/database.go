package db

import (
	"log"
	"os"
	"time"

	"github.com/sobhanhsa/simpleblog/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faled connect to database \n", err)
	}

}

func CreateArticle(auther string, title string, body string, hashtag string) (*gorm.DB, models.Article) {
	article := models.Article{Auther: auther, Title: title, Body: body, HashTag: hashtag}

	result := DB.Create(&article)

	return result, article
}

func CreateUser(email string, username string, name string, password string) models.User {
	user := models.User{Email: email, Username: username, Name: name, Password: password}

	DB.Create(&user)

	return user
}

func ModelMigrate(model interface{}) {
	DB.AutoMigrate(model)
}

func PrementDelete() {
	var deletedArticles []models.Article

	DB.Unscoped().Find(&deletedArticles)

	// log.Fatalln(deletedArticles)

	for i := 0; i > len(deletedArticles); i++ {
		var deletedTime time.Time = deletedArticles[i].DeletedAt.Time
		var now time.Time = time.Now()

		deletedTime = deletedTime.Add(24 * time.Hour)

		if !now.Before(deletedTime) {
			DB.Unscoped().Delete(&deletedArticles[i])
		}
	}
}

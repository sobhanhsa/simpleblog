package controllers

import (
	// "net/http"

	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/models"
	"github.com/sobhanhsa/simpleblog/utils"
)

func MainPage(c *gin.Context) {
	// var articles models.Article
	var articles []models.Article

	var dateOrder string = c.Query("dateorder")

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 4
	}

	var orderQeury string

	if orderQeury = "created_at DESC"; dateOrder == "oldest" {
		orderQeury = "created_at ASC"
	}

	db.DB.Order(orderQeury).Limit(limit).Find(&articles)

	c.JSON(200, gin.H{"message": articles})
}

func PublishArticle(c *gin.Context) {
	// get authentication status
	userauthstatus, _ := c.Get("userAuthStatus")

	//check the user authentication status
	if userauthstatus == nil || userauthstatus == false {
		c.JSON(401, gin.H{"message": "your not logged in", "user": userauthstatus})
		return
	}

	//convert any type variable to models.User variable
	User, hasdone := utils.UserAdjust(userauthstatus)

	//if type converting fails
	if !hasdone {
		c.JSON(500, gin.H{"message": "some thing went wrong (please infrom server staff)", "user": User})
		return
	}

	var reqBody struct {
		Title   string
		Body    string
		Hashtag string
	}

	//get data from req body and store that in reqBody
	c.Bind(&reqBody)

	if (reqBody.Body == "") || (reqBody.Title == "") {
		c.JSON(400, gin.H{"message": "please input required fields such title and body"})
		return
	}

	//publish article
	result, insertedArticle := db.CreateArticle(User.Username, reqBody.Title, reqBody.Body, reqBody.Hashtag)

	if result.Error != nil {
		c.JSON(400, gin.H{"message": "an error occurred", "error": result.Error})
		return
	}
	//return response
	c.JSON(200, gin.H{"message": "your acticle succesfully published", "article": insertedArticle})
}

func ShowArticles(c *gin.Context) {
	var articles []models.Article
	db.DB.Find(&articles)
	c.JSON(200, gin.H{"articles": articles})
}

func ShowArticle(c *gin.Context) {
	var article models.Article

	title := c.Param("title")

	if title == "" {
		c.JSON(400, gin.H{"message": "please input the desired title name"})
		return
	}

	db.DB.Where("title=?", title).Find(&article)

	if article.ID == 0 {
		c.JSON(400, gin.H{"message": "no article founded with " + title + " title"})
		return
	}

	c.JSON(200, gin.H{"articles": article})
}

func UpdateArticle(c *gin.Context) {
	// get authentication status
	userauthstatus, _ := c.Get("userAuthStatus")

	//check the user authentication status
	if userauthstatus == nil || userauthstatus == false {
		c.JSON(401, gin.H{"message": "your not logged in", "user": userauthstatus})
		return
	}

	//convert any type variable to models.User variable
	User, hasdone := utils.UserAdjust(userauthstatus)

	if !hasdone {
		c.JSON(500, gin.H{"message": "some thing went wrong (please infrom server staff)", "user": User})
		return
	}

	title := c.Param("title")

	var article models.Article

	var body struct {
		Title   string
		Body    string
		Hashtag string
	}

	c.Bind(&body)

	if result := db.DB.Where("title=? and auther=?", title, User.Username).Find(&article); article.ID == 0 {
		c.JSON(400, gin.H{"message": "you dont have any article with this title", "erorr": result.Error})
		return
	}

	if result := db.DB.Model(&article).Updates(models.Article{Auther: "",
		Title: body.Title, Body: body.Body, HashTag: body.Hashtag}); result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(400, gin.H{"message": "update process faild", "erorr": result.Error, "title": title})
		return
	}

	c.JSON(200, gin.H{"article": article})
}
func DeleteArticle(c *gin.Context) {
	// get authentication status
	userauthstatus, _ := c.Get("userAuthStatus")

	//check the user authentication status
	if userauthstatus == nil || userauthstatus == false {
		c.JSON(401, gin.H{"message": "your not logged in", "user": userauthstatus})
		return
	}

	//convert any type variable to models.User variable
	User, hasdone := utils.UserAdjust(userauthstatus)

	if !hasdone {
		c.JSON(500, gin.H{"message": "some thing went wrong (please infrom server staff)", "user": User})
		return
	}

	title := c.Param("title")

	var article models.Article

	if result := db.DB.Where("title=? and auther=?", title, User.Username).Find(&article); article.ID == 0 {
		c.JSON(400, gin.H{"message": "you dont have any article with this title", "erorr": result.Error})
		return
	}

	db.DB.Unscoped().Delete(&article)

	c.JSON(200, gin.H{"message": "article succesfully deleted"})
}

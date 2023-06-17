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

	var articles []models.Article

	var dateOrder string = c.Query("dateorder")

	//convert strin to int
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

func ArticleCategory(c *gin.Context) {

	var category string = c.Param("category")

	var articles []models.Article

	if !utils.CheckCat(category) {
		c.JSON(400, gin.H{"message": "undefined category"})
		return
	}

	db.DB.Where("category = ?", category).Find(&articles)

	c.JSON(200, gin.H{"articles": articles})

}

func SearchArticle(c *gin.Context) {

	var searchValue string = c.Query("search_query")

	var articles []models.Article

	var findedArticle models.Article

	db.DB.Where("title = ?", searchValue).First(&findedArticle)

	db.DB.Model(&models.Article{}).Where("hash_tag LIKE ?", "%"+searchValue+"%").
		Order("created_at DESC").Find(&articles)

	c.JSON(200, gin.H{"finded article": findedArticle, "related articles": articles})

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
		Title    string
		Category string
		Body     string
		Hashtag  string
	}

	//get data from req body and store that in reqBody
	c.Bind(&reqBody)

	if (reqBody.Body == "") || (reqBody.Title == "") || (reqBody.Category == "") {
		c.JSON(400, gin.H{"message": "please input required fields such title , body and category"})
		return
	}

	//check the article category
	if !(utils.CheckCat(reqBody.Category)) {
		c.JSON(400, gin.H{"message": "your desired caterogy is not available",
			"available categories": utils.Categoryies})
		return
	}

	//publish article
	result, insertedArticle := db.CreateArticle(User.Username,
		reqBody.Category, reqBody.Title, reqBody.Body, reqBody.Hashtag)

	if result.Error != nil {
		c.JSON(400, gin.H{"message": "an error occurred", "error": result.Error})
		return
	}
	//return response
	c.JSON(200, gin.H{"message": "your acticle succesfully published", "article": insertedArticle})
}

func ShowProfile(c *gin.Context) {

	var username string = c.Param("username")

	var User models.User

	var articles []models.Article

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		limit = 4
	}

	db.DB.Where("username=?", username).Find(&User)

	if User.ID == 0 {
		c.JSON(400, gin.H{"message": "no user with this username exists", "user": username})
		return
	}

	db.DB.Where("auther=?", username).Order("created_at DESC").Limit(limit).Find(&articles)

	c.JSON(200, gin.H{"username": username, "email": User.Email, "name": User.Name, "articles": articles})

}

func ShowArticles(c *gin.Context) {
	var articles []models.Article
	db.DB.Find(&articles)
	c.JSON(200, gin.H{"articles": articles})
}

func ShowArticle(c *gin.Context) {
	var article models.Article

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(400, gin.H{"message": "invalid id", "id": id})
	}

	db.DB.First(&article, id)

	if article.ID == 0 {
		c.JSON(400, gin.H{"message": "no article founded with given id", "id": id})
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

	id := c.Param("id")

	var article models.Article

	var body struct {
		Title    string
		Category string
		Body     string
		Hashtag  string
	}

	c.Bind(&body)

	//check the article category
	if !(utils.CheckCat(body.Category)) {
		c.JSON(400, gin.H{"message": "your desired caterogy is not available",
			"available categories": utils.Categoryies})
		return
	}

	if result := db.DB.Where("id=? and auther=?", id, User.Username).Find(&article); article.ID == 0 {
		c.JSON(400, gin.H{"message": "you dont have any article with this id", "erorr": result.Error})
		return
	}

	if result := db.DB.Model(&article).Updates(models.Article{Auther: "", Category: body.Category,
		Title: body.Title, Body: body.Body, HashTag: body.Hashtag}); result.Error != nil {
		fmt.Println(result.Error)
		c.JSON(400, gin.H{"message": "update process faild", "erorr": result.Error, "id": id})
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

	id := c.Param("id")

	var article models.Article

	if result := db.DB.Where("id=? and auther=?", id, User.Username).Find(&article); article.ID == 0 {
		c.JSON(400, gin.H{"message": "you dont have any article with this id", "erorr": result.Error})
		return
	}

	db.DB.Unscoped().Delete(&article)

	c.JSON(200, gin.H{"message": "article succesfully deleted"})
}

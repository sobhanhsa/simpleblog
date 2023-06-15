package controllers

import (
	// "log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/models"
	"github.com/sobhanhsa/simpleblog/validators"
	"golang.org/x/crypto/bcrypt"
)

func LoginPage(c *gin.Context) {
	//check the user authentication status
	if userauthstatus, _ := c.Get("userAuthStatus"); (userauthstatus != false) && (userauthstatus != nil) {
		c.JSON(200,
			gin.H{"message": "you are already logged in;", "user": userauthstatus})
		return
	}

	var User models.User

	var userinfo struct {
		Email    string
		Username string
		Password string
	}

	c.Bind(&userinfo)

	if (userinfo.Email == "" && userinfo.Username == "") || userinfo.Password == "" {
		c.JSON(400, gin.H{"message": "please input required fields (email or username and password)"})
		return
	}
	if userinfo.Email == "" {
		db.DB.Where("username=?", userinfo.Username).First(&User)

	} else {
		db.DB.Where("email=?", userinfo.Email).First(&User)
	}
	if userinfo.Username == "" {
		c.JSON(400, gin.H{"message": "no account was founded with given username/email"})
		return
	}
	var hashPassErr error = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(userinfo.Password))
	if hashPassErr != nil {
		c.JSON(401, gin.H{"message": "incorrect password"})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": User.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	// log.Fatal(err)
	if err != nil {
		c.JSON(500, gin.H{"message": "some thing wrong happened", "error": err,
			"token": tokenString})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.String(200, tokenString)
}
func SignUpPage(c *gin.Context) {

	var userinfo struct {
		Email    string
		Username string
		Name     string
		Password string
	}

	c.Bind(&userinfo)

	if (userinfo.Username == "") || (userinfo.Name == "") || (userinfo.Password == "") {
		c.JSON(400, gin.H{"message": "please fill the required fields such as password , email and username"})
		return
	}

	if !(validators.IsEmailValid(userinfo.Email)) {
		c.JSON(400, gin.H{"message": "invalid email please emendate it"})
		return
	}

	var User models.User

	db.DB.Where("email=?", userinfo.Email).Find(&User)
	if User.Email != "" {
		c.JSON(400, gin.H{"message": "taken email"})
		return
	}
	db.DB.Where("username=?", userinfo.Username).Find(&User)
	if User.Username != "" {
		c.JSON(400, gin.H{"message": "taken username"})
		return
	}

	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(userinfo.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{"messgae": "something unusual happened; please try later"})
		return
	}

	// c.JSON(200, gin.H{"hashed password": hashedpassword})
	// return

	var result = db.CreateUser(userinfo.Email, userinfo.Username, userinfo.Name, string(hashedpassword))

	if result.Error != nil {
		c.JSON(500, gin.H{"messgae": "something unusual happened; please try later", "data": result})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": User.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"message": "some thing wrong happened", "error": err,
			"token": tokenString})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.String(200, tokenString)

	c.JSON(200, gin.H{"message": "congratulation your account succesfully created"})
}
func UserValidate(c *gin.Context) {
	var resMessage string = "authrurization was succesfull"
	//get the user from costumized request
	user, _ := c.Get("userAuthStatus")
	if user == false {
		resMessage = "authrurization was unsuccesfull"
	}
	c.JSON(200, gin.H{"message": resMessage, "user": user})
}

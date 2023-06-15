package middlewares

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sobhanhsa/simpleblog/db"
	"github.com/sobhanhsa/simpleblog/models"
)

func UserAuth(c *gin.Context) {
	// fmt.Println("in the middleware")

	//get signedtoken from the cookie
	signedToken, err := c.Cookie("Authorization")

	if err != nil {
		c.Set("userAuthStatus", false)
		return
	}

	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["sub"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.Set("userAuthStatus", false)
		return
	}

	// decode-validate the signedtoken
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.Set("userAuthStatus", "token expired")
			return
		}

		var User models.User

		//find matched user with token sub
		db.DB.Where("id =?", claims["sub"]).First(&User)

		if User.ID == 0 {
			c.Set("userAuthStatus", false)
			return
		}

		//set user password to empty string
		User.Password = "0000"

		//attach to request
		c.Set("userAuthStatus", User)

	} else {
		c.Set("userAuthStatus", false)
	}

	c.Next()
}

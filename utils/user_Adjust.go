package utils

import (
	"github.com/sobhanhsa/simpleblog/models"
)

func UserAdjust(userstruct interface{}) (models.User, bool) {

	user, ok := userstruct.(models.User)

	return user, ok
}

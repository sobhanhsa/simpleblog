package utils

import (
	"github.com/sobhanhsa/simpleblog/models"
)

func UserAdjust(userstruct interface{}) (models.User, bool) {

	user, ok := userstruct.(models.User)

	return user, ok

	// input := fmt.Sprintf(`%+v`, userstruct)

	// // input = strings.ReplaceAll(input, "Model:", "")

	// fmt.Println(input)

	// var user models.User

	// err := json.Unmarshal([]byte(input), &user)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return user, false
	// }

	// return user, true

}

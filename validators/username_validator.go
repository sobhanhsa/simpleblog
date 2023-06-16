package validators

import (
	"regexp"
)

var usernameRegex string = "^[A-Za-z][A-Za-z0-9]{3,14}$"

func UsernameValidator(username string) bool {

	r := regexp.MustCompile(usernameRegex)

	return r.MatchString(username)
}

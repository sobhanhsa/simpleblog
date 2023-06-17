package utils

var Categoryies []string = []string{
	"science",
	"business",
	"lifestyle",
	"sport",
}

func CheckCat(caterogy string) bool {
	var status bool

	for _, c := range Categoryies {
		status = false
		if c == caterogy {
			status = true
			break
		}
	}

	return status
}

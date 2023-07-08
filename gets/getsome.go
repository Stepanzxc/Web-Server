package gets

import (
	"strconv"
)

func GetId(id string) int {
	n, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return n
}

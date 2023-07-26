package faker

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
)

func RandomWordFromFile(fileName string) string {
	file, err := os.Open(fileName)
	if err != nil {
		return ""
	}
	defer file.Close()
	var line []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = append(line, scanner.Text())
	}
	num := rand.Intn(len(line))
	arr := strings.Split(line[num], ",")

	result := arr[rand.Intn(len(arr)-1)]
	return result
}

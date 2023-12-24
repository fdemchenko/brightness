package brightness

import (
	"os"
	"strconv"
	"strings"
)

func getNumberFromFile(filename string) (int, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return 0, nil
	}

	value, err := strconv.Atoi(strings.TrimSpace(string(content)))
	if err != nil {
		return 0, nil
	}
	return value, nil
}

func checkBounds(value, minValue, maxValue int) int {
	if value < minValue {
		return minValue
	}
	if value > maxValue {
		return maxValue
	}
	return value
}

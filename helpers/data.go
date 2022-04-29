package helpers

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateData() []int {
	var max_num = 100
	var max = &max_num

	water := rand.Intn(*max + 1)
	wind := rand.Intn(*max + 1)

	payload := []int{water, wind}

	return payload
}

func GetStatus() ([]int, string) {
	var status string

	data := GenerateData()
	switch {
	case data[0] < 5 && data[1] < 6:
		status = "Aman"
		return data, status
	case data[0] > 8 || data[1] > 15:
		status = "Bahaya"
		return data, status
	default:
		status = "Siaga"
		return data, status
	}

}

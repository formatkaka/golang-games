package eggtoss

import "fmt"

func contains(arr [3]float64, str float64) bool {
	fmt.Println("check : ", str)
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

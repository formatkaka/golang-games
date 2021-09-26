package eggtoss

func contains(arr [3]float64, str float64) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

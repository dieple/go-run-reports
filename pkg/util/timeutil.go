package util

import "time"

func CurrentMonth() string {
	return time.Now().Format("2006-01")
}

func PlanToLimit(plan string) int {
	switch plan {
	case "Ultimate":
		return 1000
	case "Enterprise":
		return 500
	case "Basic":
		return 100
	case "Lite":
		return 20
	case "Trial":
		return 10
	default:
		return 0
	}
}
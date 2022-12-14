package common

import (
	"strings"
)

func ContainsIgnoreCase(arr []string, target string) bool {
	target = strings.ToUpper(target)
	for _, v := range arr {
		if strings.ToUpper(v) == target {
			return true
		}
	}
	return false
}

// Min returns the minimum value in a list of values.
func Min(values ...interface{}) float64 {
	minValue := values[0].(float64)
	for _, value := range values {
		switch value := value.(type) {
		case int64:
			if value < int64(minValue) {
				minValue = float64(value)
			}
		case float64:
			if value < minValue {
				minValue = value
			}
		case int:
			// Handle float64 values
			if value < int(minValue) {
				minValue = float64(value)
			}
		default:
			// ignore other types
		}
	}
	return minValue
}

// Max returns the maximum value in a list of values.
func Max(values ...interface{}) float64 {
	maxValue := values[0].(float64)
	for _, value := range values {
		switch value := value.(type) {
		case int64:
			if value > int64(maxValue) {
				maxValue = float64(value)
			}
		case float64:
			if value > maxValue {
				maxValue = value
			}
		case int:
			if value > int(maxValue) {
				maxValue = float64(value)
			}
		default:
			// ignore other types
		}
	}
	return maxValue
}

// RemoveElementByIndex removes an element from a slice by index.
func RemoveElementByIndex[T any](slice []T, index int) []T {
	// check if index is out of range
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// InsertElementByIndex inserts an element to a slice by index.
func InsertElementByIndex[T any](slice []T, index int, element T) []T {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], append([]T{element}, slice[index:]...)...)
}

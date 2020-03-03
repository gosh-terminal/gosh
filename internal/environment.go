package internal

import (
	"reflect"
)

func validateBooleanVariable(b string) bool {
	yesOpts := [3]string{"yes", "1", "true"}

	return ItemExists(yesOpts, b)
}

// ItemExists Check if item exists in array
func ItemExists(arrayType interface{}, item interface{}) bool {
	arr := reflect.ValueOf(arrayType)

	for i := 0; i < arr.Len(); i++ {
		if arr.Index(i).Interface() == item {
			return true
		}
	}

	return false
}

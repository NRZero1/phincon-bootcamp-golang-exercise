package utils

import (
	"fmt"
	"reflect"
)

func PrintToConsole(objectToPrint interface{}) {
	rv := reflect.ValueOf(objectToPrint)

	// var out []interface{}
	if rv.Kind() == reflect.Slice {
		for i := 0; i < rv.Len(); i ++ {
			fmt.Println(rv.Index(i))
		}
	}
}
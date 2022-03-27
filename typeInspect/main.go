package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x1 float32 = 5.7

	inspectIfTypeFloat(x1)
}

func inspectIfTypeFloat(i interface{}) {
	v := reflect.ValueOf(i)
	fmt.Println("Type:", v.Type()) //same as reflect.TypeOf(i)
	fmt.Println("Is type float64?", v.Kind() == reflect.Float64)
	fmt.Println("Float Value:", v.Float())
}

package main

import (
	"fmt"
	"reflect"
)

// type Printer interface {
// 	Print(s string)
// }

// type pStruct struct {
// 	s string
// }

// func (p *pStruct) Print(s string) {
// 	p.s = s
// 	fmt.Println(s)
// }

// func main() {
// 	p := new(pStruct)
// 	inspectType(p)
// }

// func inspectType(obj interface{}) {
// 	v := reflect.ValueOf(obj)
// 	t := v.Type()
// 	myInterface := reflect.TypeOf((*Printer)(nil)).Elem()

// 	fmt.Println("obj implements Printer?", t.Implements(myInterface))

// 	if t.Implements(myInterface) {
// 		printFunc := v.MethodByName("Print")
// 		args := []reflect.Value{reflect.ValueOf("Printing Hello from a refleciton object!!!")}
// 		printFunc.Call(args)
// 	}
// }

// func main() {
// 	type myStruct struct {
// 		Field1 int
// 		Field2 string
// 		field3 float64
// 	}
// 	mys := myStruct{2, "Hello", 2.4}

// 	mysRValue := reflect.ValueOf(mys)
// 	mysRType := reflect.TypeOf(mys)

// 	for i := 0; i < mysRType.NumField(); i++ {
// 		fieldRType := mysRType.Field(i) //datatype: StructField
// 		fieldRValue := mysRValue.Field(i)
// 		fmt.Printf("Field Name: '%s', field type: '%s', field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface)
// }

func main() {
	type myStruct struct {
		Field1 int     `alias:"f1" desc:"field number 1"`
		Field2 string  `alias:"f2" desc:"field number 2"`
		Field3 float64 `alias:"f3" desc:"field number 3"`
	}
	mys := myStruct{2, "Hello", 2.4}
	InspectStructType(&mys)
}

func InspectStructType(i interface{}) {
	mysRValue := reflect.ValueOf(i)
	if mysRValue.Kind() != reflect.Ptr {
		return
	}
	mysRValue = mysRValue.Elem()
	if mysRValue.Kind() != reflect.Struct {
		return
	}
	mysRValue.Field(0).SetInt(15)
	mysRType := mysRValue.Type() //reflect.TypeOf(i)

	for i := 0; i < mysRType.NumField(); i++ {
		fieldRType := mysRType.Field(i) //datatype: StructField
		fieldRValue := mysRValue.Field(i)
		fmt.Printf("Field Name: '%s', field type: '%s', field value: '%v' \n", fieldRType.Name, fieldRType.Type, fieldRValue.Interface())
		fmt.Println("Struct tags, alias:", fieldRType.Tag.Get("alias"), "description:", fieldRType.Tag.Get("desc"))
	}
}

// 	InspectStructType(mys)
// func InspectStructType(i interface{}) {
// 	mysRValue := reflect.ValueOf(i)
// 	mysRType := reflect.TypeOf(i)

// 	for i := 0; i < mysRType.NumField(); i++ {
// 		fieldRType := mysRType.Field(i) //datatype: StructField
// 		fieldRValue := mysRValue.Field(i)
// 		fmt.Printf("Field Name: '%s', field type: '%s', field value: '%v', fieldRType.Name, fieldRType.Type, fieldRValue")
// 	}
// }

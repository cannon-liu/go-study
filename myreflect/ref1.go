package main

import (
	"fmt"
	"reflect"
)

type Order struct {
	ordId      int
	customerId int
}

func query(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()
	fmt.Println("Type: ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind ", k)
	if reflect.TypeOf(q).Kind() == reflect.Struct {
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			fmt.Println("value field ", v.Field(i))
		}
	}

}

func main() {
	o := Order{
		ordId:      123,
		customerId: 456,
	}
	query(o)
	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "hello"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)
}

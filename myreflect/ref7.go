package main

import (
	"fmt"
	"reflect"
)

//普通函数
func add(a, b int) int {
	return a + b
}

func main() {
	//将函数包装为反射值对象
	funcValue := reflect.ValueOf(add)

	//构造函数参数，传入两个整形值
	paramList := []reflect.Value{reflect.ValueOf(2), reflect.ValueOf(3)}

	//反射调用函数
	retList := funcValue.Call(paramList)

	fmt.Println(len(retList))
	fmt.Println(retList[0])
	fmt.Println(retList)
}

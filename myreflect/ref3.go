package main

import (
	"fmt"
	"reflect"
)

func main() {

	//声明整型变量a并赋初值
	var a int = 1024

	//获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)

	//获取interface{}类型的值，通过类型断言转换
	var getA int = valueOfA.Interface().(int)

	//获取64位的值，强制类型转换为int类型
	var getB int = int(valueOfA.Int())

	fmt.Println(getA, getB)
}

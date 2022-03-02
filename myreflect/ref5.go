package main

import (
	"fmt"
	"reflect"
)

func main() {

	//声明整形变量a并赋初值
	var a int = 1024

	value := reflect.ValueOf(a)
	fmt.Println("是可选址", value.CanAddr(), value.CanSet())

	//获取变量a的反射值对象
	rValue := reflect.ValueOf(&a)

	fmt.Println("是可选址", rValue.CanAddr(), rValue.CanSet())

	//取出a地址的元素(a的值)
	rValue = rValue.Elem()

	fmt.Println("是可选址", rValue.CanAddr(), rValue.CanSet())

	//尝试将a修改为1
	rValue.SetInt(1)

	//打印a的值
	fmt.Println(rValue.Int())
}

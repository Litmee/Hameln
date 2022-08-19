package main

import (
	"Hameln/pkg/cache"
	"fmt"
	"unsafe"
)

type User struct {
	//Num  int8
	//Id   int32
	//Age  int64
	Name string
	Id   int32
}

func main() {
	align := unsafe.Alignof(User{})
	fmt.Printf("最大内存对齐边界: %v \n", align)
	var user User
	fmt.Printf("内存大小: %v \n", unsafe.Sizeof(user))
	cache.DoRedis()
}

package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id   uint64
	Name string
}

func doing() string {
	return "hello world"
}

func main() {
	println(doing())
	align := unsafe.Alignof(User{})
	fmt.Printf("最大内存对齐边界: %v \n", align)
	var user User
	fmt.Printf("内存大小: %v \n", unsafe.Sizeof(user))
}

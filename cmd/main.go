package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Id    int32
	Name  string
	Money int32
}

func main() {
	align := unsafe.Alignof(User{})
	fmt.Printf("最大内存对齐边界: %v \n", align)
	var user User
	fmt.Printf("内存大小: %v \n", unsafe.Sizeof(user))
}

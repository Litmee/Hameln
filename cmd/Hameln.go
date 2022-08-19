package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Code int8
	Id   int32
	Num  int64
}

func main() {
	align := unsafe.Alignof(User{})
	fmt.Printf("最大内存对齐边界: %v \n", align)
	var user User
	fmt.Printf("内存大小: %v \n", unsafe.Sizeof(user))
}

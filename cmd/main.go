package main

import (
	"Hameln/internal"
	"fmt"
	"unsafe"
)

type User struct {
	Id      int32
	Name    string
	Address string
	Sex     int8
}

func main() {
	align := unsafe.Alignof(User{})
	fmt.Printf("最大内存对齐边界: %v \n", align)
	var user User
	fmt.Printf("内存大小: %v \n", unsafe.Sizeof(user))

	fmt.Println("-----------------------------------------------")

	var myNorth internal.MyNorth

	internal.DoNorth(&myNorth, "Good")
}

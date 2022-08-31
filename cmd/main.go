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
}

func main() {
	align := unsafe.Alignof(User{})
	fmt.Printf("最大内存对齐边界: %v \n", align)
	var user User
	fmt.Printf("内存大小: %v \n", unsafe.Sizeof(user))

	fmt.Println("-----------------------------------")

	user.Id = 100

	var btRoot = new(internal.BinaryTree)
	var btLeft = new(internal.BinaryTree)
	btLeft.SetValue(user)
	btRoot.SetLeft(btLeft)

	value := (btRoot.GetLeft().GetValue()).(User)
	fmt.Println(value.Id)

	fmt.Println(user.Id)
	fmt.Println("-----------------------------------")
	updateId(&user)
	fmt.Println(user.Id)
}

func updateId(user *User) {
	user.Id = 200
}

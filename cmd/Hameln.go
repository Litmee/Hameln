package main

import (
	"Hameln/pkg/cache"
	"fmt"
)

type IsName string

type User struct {
	Id      int32
	Name    IsName
	Address string
	Sex     int8
	Home    string
	Money   int64
}

func (u *User) Apply() string {
	return string(u.Name)
}

func main() {
	user := User{Id: 1, Name: "Tom"}
	fmt.Println(user.Apply())
	cache.DoRedis()
}

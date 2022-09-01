package internal

import (
	"fmt"
	"log"
)

type GoNorth interface {
	GetNorth(s string) (n int, err error)
}

func DoNorth(gn GoNorth, s string) {
	n, err := gn.GetNorth(s)
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("res: %v", n)
}

type MyNorth struct{}

func (mn MyNorth) GetNorth(s string) (n int, err error) {
	n = len(s)
	return n, nil
}

package main

import (
	"fmt"
	//"reflect"
	"github.com/Rbd3178/redBlackTree/tree"
)


func main() {
	var tr tree.Tree[int, int]
	err := tr.Insert(2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
	err = tr.Insert(2, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success!")
	}
}

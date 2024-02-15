package main

import (
	"fmt"
	"reflect"
	"github.com/Rbd3178/redBlackTree/tree"
)


func main() {
	tr := tree.New[int, string]
	tr.Insert(8215, "8215")

	for _, pair := range tr.InOrder() {
		for _, element := range pair {
			// Print each element of the pair
			fmt.Print(element, " ")
		}
		fmt.Println()
	}
	/*for {
		var k int
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if k == -1 {
			break
		}
		err = tr.Insert(k, "test")
		if err != nil {
			fmt.Println("Error during insertion:", err)
		} else {
			tr.Visualize()
		}
		
		fmt.Print("\n\n\n\n")
	}*/
}

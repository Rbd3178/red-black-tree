package main

import (
	"fmt"
	"github.com/Rbd3178/redBlackTree/tree"
)

func main() {
	myTree := tree.New()
	// 387, 7492, 2314, 506, 9127, 6543, 1832, 4421, 769, 890
	for {
		var k int
		_, err := fmt.Scan(&k)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		if k == -1 {
			break
		}
		err = myTree.Insert(k, "test")
		if err != nil {
			fmt.Println("Error during insertion:", err)
		} else {
			myTree.Visualize()
		}
		
		fmt.Print("\n\n\n\n")
	}
}

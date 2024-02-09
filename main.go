package main

import (
	"fmt"
	"github.com/Rbd3178/redBlackTree/tree"
)

func main() {
	fmt.Println("Test")
	myTree := tree.New()
	/* myTree.Insert(387, "387")
	myTree.Insert(7492, "7492")
	myTree.Insert(2314, "2314")
	myTree.Insert(506, "506")
	myTree.Insert(9127, "9127")
	myTree.Insert(6543, "6543")
	myTree.Insert(1832, "1832")
	myTree.Insert(4421, "4421")
	myTree.Insert(769, "769")
	myTree.Visualize()
	myTree.Insert(890, "890")
	fmt.Print("\n\n\n\n\n")
	myTree.Visualize() */
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
		myTree.Insert(k, "test")
		myTree.Visualize()
		fmt.Print("\n\n\n\n")

	}
}

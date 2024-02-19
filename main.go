package main

import (
	"fmt"
	"github.com/Rbd3178/redBlackTree/tree"
)

func main() {
	var myTree tree.Tree[int, string]
loop:
	for {
		var command string
		_, err := fmt.Scanln(&command)
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue loop
		}
		switch command {
		case "insert":
			fmt.Print("Enter key and value: ")
			var key int
			var val string
			_, err = fmt.Scanln(&key, &val)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue loop
			}
			err = myTree.Insert(key, val)
			if err != nil {
				fmt.Println("Error when inserting:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println("Success!")
			fmt.Println()

		case "delete":
			fmt.Print("Enter key: ")
			var key int
			_, err = fmt.Scanln(&key)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue loop
			}
			err = myTree.Delete(key)
			if err != nil {
				fmt.Println("Error when deleting:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println("Success!")
			fmt.Println()

		case "visualize":
			myTree.Visualize()
			fmt.Println()

		case "inorder":
			for _, pair := range myTree.InOrder() {
				for _, element := range pair {
					fmt.Print(element, " ")
				}
				fmt.Println()
			}
			fmt.Println()

		case "verify":
			err := myTree.Verify()
			if err != nil {
				fmt.Println("Not a valid red-black tree:", err)
				continue loop
			}
			fmt.Println("Valid red-black tree")
			fmt.Println()

		case "depth":
			fmt.Println(myTree.BlackDepth())
			fmt.Println()

		case "max":
			key, val, err := myTree.Max()
			if err != nil {
				fmt.Println("Error when getting max element:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println(key, val)
			fmt.Println()

		case "min":
			key, val, err := myTree.Min()
			if err != nil {
				fmt.Println("Error when getting min element:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println(key, val)
			fmt.Println()

		case "size":
			fmt.Println(myTree.Size())
			fmt.Println()

		case "next":
			fmt.Print("Enter key: ")
			var key int
			_, err = fmt.Scanln(&key)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue loop
			}
			var nextKey int
			var nextVal string
			nextKey, nextVal, err = myTree.Next(key)
			if err != nil {
				fmt.Println("Error when getting next element:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println(nextKey, nextVal)
			fmt.Println()

		case "previous":
			fmt.Print("Enter key: ")
			var key int
			_, err = fmt.Scanln(&key)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue loop
			}
			var prevKey int
			var prevVal string
			prevKey, prevVal, err = myTree.Prev(key)
			if err != nil {
				fmt.Println("Error when getting previous element:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println(prevKey, prevVal)
			fmt.Println()

		case "at":
			fmt.Print("Enter key: ")
			var key int
			_, err = fmt.Scanln(&key)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue loop
			}
			var val string
			val, err = myTree.At(key)
			if err != nil {
				fmt.Println("Error when getting vakue at key:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println(val)
			fmt.Println()

		case "assign":
			fmt.Print("Enter key and value: ")
			var key int
			var val string
			_, err = fmt.Scanln(&key, &val)
			if err != nil {
				fmt.Println("Error reading input:", err)
				continue loop
			}
			err = myTree.Assign(key, val)
			if err != nil {
				fmt.Println("Error when assigning:", err)
				fmt.Println()
				continue loop
			}
			fmt.Println("Success!")
			fmt.Println()

		case "exit":
			break loop

		default:
			fmt.Println("Unknown command")
			fmt.Println()
		}
	}
}

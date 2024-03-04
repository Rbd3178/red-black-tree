# Red-Black Tree implementation in Go

Red-Black tree is a self-balancing binary search tree, a container for storing key-value pairs, with fast operations for search, insertion and deletion of elements. Unlike a hash table, it also maintains the order of elements at all times.

## tree.Tree type
Generics were used to allow the keys to be of any type that supports the operators `<`, `<=`, `>=`, `>`. The value can be of any type.

Creating an instance of Tree:
```Go
package main

// importing the package
import "github.com/Rbd3178/redBlackTree/tree"


func main() {
    ...
    // creating a tree with strings for keys and ints for values
    var myTree tree.Tree[string, int]
    ...
}
```


## Methods

### Search

- `Tree.At(key)` Returns a value stored by a passed key in O(log(n)). Returns an error if the passed key is not in the structure.

- `Tree.Next(key)` Returns the first key-value pair with the key bigger than the passed one, even if no values are stored by the passed key, in O(log(n)).

- `Tree.Prev(key)` Returns the first key-value pair with the key smaller than the passed one, even if no values are stored by the passed key, in O(log(n)).

- `Tree.Max()` Returns the key-value pair with the biggest key in O(log(n)). Returns an error if the tree is empty.

- `Tree.Min()` Returns the key-value pair with the smallest key in O(log(n)). Returns an error if the tree is empty.

- `Tree.BlackDepth()` Returns the amount of black nodes on the path from the root to leaf nodes in O(log(n)).

- `Tree.Size()` Returns the amount of pairs stored in the tree in O(1).

### Modification

- `Tree.Insert(key, val)` Adds an element to the tree with key and value equal to `key` and `val` respectfully in O(log(n)). Returns an error if an element with the passed key already exists.

- `Tree.Assign(key, val)` Changes the value by the passed key to `val` in O(log(n)). Returns an error if the passed key is not in the structure.

- `Tree.Delete(key)` Deletes an element with the key equal to `key` from the tree in O(log(n)). Returns an error if there is no element with the passed key.

### Traversal

- `Tree.InOrder()` Returns a slice of all key-value pairs stored in the tree, ordered by key in O(n).

- `Tree.Range(min, max)` Returns a slice of key-value pairs with a key `>= min` and `< max`, ordered by key in O(k), where k is the amount of such pairs.

- `Tree.Verify()` Traverses the tree and checks if all properties of a red-black tree are met in O(n), returning an error if it encounters a problem with the tree structure.

- `Tree.Visualize()` Prints out a visualization of the tree to the terminal, nodes are represented by their key and color.

## Unit tests
Automatic table driven tests are implemented in `tree_test.go` for most of the methods.

## CLI app
In `main.go` there is a small CLI app that allows for interaction with the tree. 
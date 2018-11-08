package tree

import (
	"fmt"
	"reflect"
)

// Item is a generic(ish) field for use in the tree
type Item interface{}

// Node contains fields for binary search tree
type Node struct {
	data  Item
	left  *Node
	right *Node
}

// BST is a binary search tree implementation
type BST struct {
	root *Node
	size int
}

// Min returns the minimum value for the tree
func (bst *BST) Min() Item {
	if bst.root == nil {
		return nil
	}
	node := bst.root
	for node.left != nil {
		node = node.left
	}
	return node.data
}

// Max returns max item in the tree
func (bst *BST) Max() Item {
	if bst.root == nil {
		return nil
	}
	node := bst.root
	for node.right != nil {
		node = node.right
	}
	return node.data
}

// Search navigates the tree to determine if tree contains values provided
func (bst *BST) Search(data Item) bool {

	return search(bst.root, data)

}

func search(n *Node, key Item) bool {
	if n == nil {
		return false
	}

	location := Compare(key, n.data)
	if location == "greater" {
		return search(n.right, key)
	}

	if location == "less" {
		return search(n.left, key)
	}
	return true
}

// Insert places a node inside the tree properly organized
func (bst *BST) Insert(data Item) {

	n := &Node{data: data}
	if bst.root == nil {
		bst.root = n
		bst.size++
		return
	}

	insertNode(bst.root, n)
	bst.size++
}

func insertNode(n *Node, newNode *Node) {
	if n == nil {
		n = newNode
		return
	}
	location := Compare(n.data, newNode.data)
	if location == "less" {
		if n.left == nil {
			n.left = newNode
		} else {
			insertNode(n.left, newNode)
		}
	} else {
		if n.right == nil {
			n.right = newNode
		} else {
			insertNode(n.right, newNode)
		}
	}
}

// Remove finds an item in the bst and removes it (reorganizes using next smallest item to right)
func (bst *BST) Remove(key Item) {
	node := remove(bst.root, key)
	bst.root = node
	bst.size--
}

func remove(node *Node, key Item) *Node {
	if node == nil {
		return nil
	}

	loc := Compare(key, node.data)
	if loc == "left" {
		node.left = remove(node.left, key)
		return node
	}

	if loc == "right" {
		node.right = remove(node.right, key)
		return node
	}

	// Case of equality condition
	if node.left == nil && node.right == nil {
		node = nil
		return nil
	}

	// Having only left child
	if node.left == nil {
		node = node.right
		return node
	}

	// Only right child
	if node.right == nil {
		node = node.left
		return node
	}

	// Finding first field on left side of the right node (next largest field from object being deleted)
	nextNode := node.right
	for {
		if nextNode != nil && nextNode.left != nil {
			nextNode = nextNode.left
		} else {
			break
		}
	}

	node.data = nextNode.data
	node.right = remove(node.right, node.data)
	return node
}

// Root returns root of the tree
func (bst *BST) Root() *Node {
	return bst.root
}

// Size returns current size of the tree
func (bst *BST) Size() int {
	return bst.size
}

// Compare takes in two interfaces and determines which is greater
func Compare(a interface{}, b interface{}) string {
	if reflect.TypeOf(a) != reflect.TypeOf(b) {
		panic("Invalid compare, differing types")
	}
	switch a.(type) {
	case string:
		if a.(string) > b.(string) {
			return "greater"
		}
		if a.(string) < b.(string) {
			return "less"
		}
		if a.(string) == b.(string) {
			return "equals"
		}
	case int:
		if a.(int) > b.(int) {
			return "greater"
		}
		if a.(int) < b.(int) {
			return "less"
		}
		if a.(int) == b.(int) {
			return "equals"
		}
	case int32:
		if a.(int32) > b.(int32) {
			return "greater"
		}
		if a.(int32) < b.(int32) {
			return "less"
		}
		if a.(int32) == b.(int32) {
			return "equals"
		}
	case int64:
		if a.(int64) > b.(int64) {
			return "greater"
		}
		if a.(int64) < b.(int64) {
			return "less"
		}
		if a.(int64) == b.(int64) {
			return "equals"
		}
	case uint:
		if a.(uint) > b.(uint) {
			return "greater"
		}
		if a.(uint) < b.(uint) {
			return "less"
		}
		if a.(uint) == b.(uint) {
			return "equals"
		}
	case uint32:
		if a.(uint32) > b.(uint32) {
			return "greater"
		}
		if a.(uint32) < b.(uint32) {
			return "less"
		}
		if a.(uint32) == b.(uint32) {
			return "equals"
		}
	case uint64:
		if a.(uint64) > b.(uint64) {
			return "greater"
		}
		if a.(uint64) < b.(uint64) {
			return "less"
		}
		if a.(uint64) == b.(uint64) {
			return "equals"
		}
	default:
		panic(fmt.Sprintf("Cannot determine type: %v", reflect.TypeOf(a)))
	}

	panic("could not resolve")

}

// String prints a visual representation of the tree
func (bst *BST) String() {
	fmt.Println("------------------------------------------------")
	stringify(bst.root, 0)
	fmt.Println("------------------------------------------------")
}

// internal recursive function to print a tree
func stringify(n *Node, level int) {
	if n != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "       "
		}
		format += "---[ "
		level++
		stringify(n.left, level)
		fmt.Printf(format+"%v\n", n.data)
		stringify(n.right, level)
	}
}

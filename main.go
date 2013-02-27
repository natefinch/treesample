package main

import (
	"fmt"
	"github.com/natefinch/tree"
)

func main() {

	// Create a StringTree from the given slice data
	st := NewStringTree([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
	fmt.Println("Data:", st.Data)

	walk := func(n *tree.Node) bool {
		fmt.Print(st.Data[n.Val] + " ")
		return true
	}

	// walk over the tree in order, printing each node
	fmt.Print("Tree: [")
	tree.Walk(st.Tree.Head, walk)
	fmt.Println("]")

	// add new a new item
	st.Insert("pineapple")
	fmt.Println("Data:", st.Data)

	// print the tree again
	fmt.Print("Tree: [")
	tree.Walk(st.Tree.Head, walk)
	fmt.Println("]")

	// let's do a binary search for the item we just inserted
	fmt.Printf("Tree contains pineapple: %v\n", st.Find("pineapple"))

	// now search for something that doesn't exist
	fmt.Printf("Tree contains XXXXX: %v\n", st.Find("XXXXX"))
}

// Sample code for making a new type of tree.

// NewStringTree returns a new StringTree populated with the given data
func NewStringTree(s []string) *StringTree {
	st := &StringTree{Data: s}
	st.Tree = tree.New(st)
	for i := range s {
		st.Tree.Insert(i)
	}
	return st
}

// StringTree just packages the data and the tree together
type StringTree struct {
	Data   []string
	Tree   *tree.Tree
	search string
}

// Cmp is just a standard string compare
func (s *StringTree) Compare(i, j int) int8 {
	var target string
	if i == -1 {
		target = s.search
	} else {
		target = s.Data[i]
	}
	other := []rune(s.Data[j])
	for i, r := range target {
		if i > len(other) {
			return 1
		}
		c := r - other[i]
		if c < 0 {
			return -1
		}
		if c > 0 {
			return 1
		}
	}
	if len(other) > len(target) {
		return -1
	}
	return 0
}

// Insert inserts the given string into the tree
func (s *StringTree) Insert(val string) {
	// add data to the backing slice
	s.Data = append(s.Data, val)

	// use the Tree's insert method, passing StringData's comparison funcion
	s.Tree.Insert(len(s.Data) - 1)
}

// Find returns true if the string is in the tree
func (s *StringTree) Find(val string) bool {
	s.search = val
	return s.Tree.Search(-1) != nil
}

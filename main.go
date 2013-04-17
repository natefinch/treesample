package main

import (
	"fmt"
	"github.com/natefinch/tree"
)

func main() {

	// Create a StringTree from the given slice data
	st := NewStringTree([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
	fmt.Println("Data:", st.data)

	walk := func(n *tree.Node) bool {
		fmt.Print(st.data[n.Val] + " ")
		return true
	}

	// walk over the tree in order, printing each node
	fmt.Print("Tree: [")
	tree.Walk(st.tree.Head, walk)
	fmt.Println("]")

	// add new a new item
	st.Insert("pineapple")
	fmt.Println("Data:", st.data)

	// print the tree again
	fmt.Print("Tree: [")
	tree.Walk(st.tree.Head, walk)
	fmt.Println("]")

	// let's do a binary search for the item we just inserted
	fmt.Printf("Tree contains pineapple: %v\n", st.Find("pineapple"))

	// now search for something that doesn't exist
	fmt.Printf("Tree contains XXXXX: %v\n", st.Find("XXXXX"))

}

// NewStringTree returns a new StringTree populated with the given data
func NewStringTree(s []string) *StringTree {
	st := &StringTree{data: s}
	st.tree = &tree.Tree{}
	for i := range s {
		st.tree.Insert(i, st.compare)
	}
	return st
}

// Sample code for making a new type of tree.

// StringTree just packages the data and the tree together
type StringTree struct {
	data   []string
	tree   *tree.Tree
	search string
}

// Compare compares two strings in the slice
func (s *StringTree) compare(i, j int) int8 {
	var target []rune
	if i == -1 {
		target = []rune(s.search)
	} else {
		target = []rune(s.data[i])
	}
	other := []rune(s.data[j])
	return strCmp(target, other)
}

// Insert inserts the given string into the tree
func (s *StringTree) Insert(val string) {
	s.data = append(s.data, val)

	s.tree.Insert(len(s.data)-1, s.compare)
}

// Find returns true if the string is in the tree
func (s *StringTree) Find(val string) bool {
	s.search = val
	n, _ := s.tree.Search(-1, s.compare)
	return n != nil
}

// Delete removes a value from the tree, returning true if the value was in the tree
func (s *StringTree) Delete(val string) bool {
	s.search = val
	if i, err := s.tree.Delete(-1, s.compare); err == nil {
		// ideally we'd clean this up better
		s.data[i] = ""
		return true
	} else {
		return false
	}
}

// strCmp is just a standard string compare
func strCmp(target, other []rune) int8 {
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

package main

import (
	"fmt"
	"github.com/natefinch/tree"
)

func main() {

	// Create a StringTree from the given slice data
	st := Treeify([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
	fmt.Println("Data:", st.Data)

	// walk over the tree in order, printing each node
	fmt.Print("Tree: [")
	tree.Walk(st.Tree, func(t *tree.Tree) { fmt.Print(st.Data[t.Val] + " ") })
	fmt.Println("]")

	// add new a new item
	st.Insert("pineapple")
	fmt.Println("Data:", st.Data)

	// print the tree again
	fmt.Print("Tree: [")
	tree.Walk(st.Tree, func(t *tree.Tree) { fmt.Print(st.Data[t.Val] + " ") })
	fmt.Println("]")

	// let's do a binary search for the item we just inserted
	fmt.Printf("Tree contains pineapple: %v\n", st.Find("pineapple"))

	// now search for something that doesn't exist
	fmt.Printf("Tree contains XXXXX: %v\n", st.Find("XXXXX"))
}

// Sample code for making a new type of tree.

// Treeify returns a new StringTree populated with the given data
func Treeify(s []string) *StringTree {
	st := &StringTree{StringData(s), tree.New()}
	for i, val := range s {
		st.Tree.Insert(i, st.Data.Cmp(val))
	}
	return st
}

// StringData is just a slice that implements a comparison function
type StringData []string

// Cmp returns a closure that will compare the given string with
// the string located at an index of the underlying slice
func (s StringData) Cmp(val string) func(int) int8 {

	// this is just a standard string compare based on the runes
	return func(idx int) int8 {
		other := []rune(s[idx])
		for i, r := range val {
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
		return 0
	}
}

// StringTree just packages the data and the tree together
type StringTree struct {
	Data StringData
	Tree *tree.Tree
}

// Insert inserts the given string into the tree
func (s *StringTree) Insert(val string) {
	// add data to the backing slice
	s.Data = append(s.Data, val)

	// use the Tree's insert method, passing StringData's comparison closure
	s.Tree.Insert(len(s.Data)-1, s.Data.Cmp(val))
}

// Find returns true if the string is in the tree
func (s *StringTree) Find(val string) bool {

	// use the tree's Search function, passing StringData's comparison closure
	// Search returns the index of the item in the backing data if it exists, or -1 if it doesn't exist
	return s.Tree.Search(s.Data.Cmp(val)) != -1
}

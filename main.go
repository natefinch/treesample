package main

import (
	"fmt"
	"github.com/natefinch/tree"
)

type StringData []string

func (s StringData) Less(val string) func(int) bool {
	return func(i int) bool { return val < s[i] }
}

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

func main() {

	// StringData is the backing data for our binary tree (just a slice of strings)
	s := StringData([]string{"banana", "apple", "domino", "cat", "zebra", "monkey", "hippo"})
	fmt.Println("Data:", s)

	t := tree.New()

	// insert using StringData's less function
	for i, v := range s {
		t.Insert(i, s.Less(v))
	}

	// walk over the tree in order, printing each node
	fmt.Print("Tree: [")
	tree.Walk(t, func(t *tree.Tree) { fmt.Print(s[t.Val] + " ") })
	fmt.Println("]")

	// add new a new item to the backing data
	w := "pineapple"
	s = append(s, w)

	fmt.Println("Data:", s)

	// now insert it into the right place in the tree
	t.Insert(len(s)-1, s.Less(w))

	// print the tree again
	fmt.Print("Tree: [")
	tree.Walk(t, func(t *tree.Tree) { fmt.Print(s[t.Val] + " ") })
	fmt.Println("]")

	// let's do a binary search for the item we just inserted
	idx := t.Search(s.Cmp(w))

	fmt.Printf("Tree contains %s: %v\n", w, idx != -1)

	// now search for something that doesn't exist
	idx = t.Search(s.Cmp("XXXXX"))
	fmt.Printf("Tree contains XXXXX: %v\n", idx != -1)

}

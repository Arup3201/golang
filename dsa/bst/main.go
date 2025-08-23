package main

import (
	"fmt"

	bst "example.com/bst/v1"
)

func main() {
	tree := bst.NewTree()
	entries := []int{20, 15, 25, 40, 10, 12, 23, 8, 13, 16, 1, 30}

	var node *bst.Node
	for _, entry := range entries {
		fmt.Printf("Inserting node %d to the tree\n", entry)

		node = bst.NewNode(entry)
		tree.Insert(node)

		fmt.Printf("Printing the tree\n")
		tree.Print()
	}

	searches := []int{1, 8, 30, 50}
	for _, search := range searches {
		fmt.Printf("Searching %d\n", search)
		_, ok := tree.Search(search)

		if !ok {
			fmt.Printf("Search failed\n")
		} else {
			fmt.Printf("Found %d\n", search)
		}
	}


	deletes := []int{40, 10}
	for _, term := range deletes {
		fmt.Printf("Deleting %d\n", term)
		node, ok := tree.Search(term)

		if !ok {
			fmt.Printf("%d not found in the tree\n", term)
			continue
		}

		tree.Delete(node)
		fmt.Printf("After deleting %d, tree: \n", term)
		tree.Print()
	}
}

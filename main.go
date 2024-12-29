package main

import (
	"fmt"
	"dsa-go/search"
)

func main() {
	arr := []int{1, 2, 3, 4, 6, 7, 8, 9}
	fmt.Println(search.BinarySearch(arr, 5))
	// trie := trie.NewTrie()
	// trie.Insert("apple")
	// trie.Insert("banana")
	// trie.Insert("app")
	// trie.Insert("orange")

	// fmt.Println(trie.StartsWith("appl"))   // true
}

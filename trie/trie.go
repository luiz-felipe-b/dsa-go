package trie

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

func (trie *Trie) Insert(word string) {
	currentNode := trie.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			currentNode.children[char] = NewTrieNode()
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEndOfWord = true
}

func (trie *Trie) Search(word string) bool {
	currentNode := trie.root
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return currentNode.isEndOfWord
}

func (trie *Trie) StartsWith(prefix string) bool {
	currentNode := trie.root
	for _, char := range prefix {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return true
}

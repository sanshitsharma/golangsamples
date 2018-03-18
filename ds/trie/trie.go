package trie

// Trie struct defines the Trie node structure
type Trie struct {
	children    []*Trie
	isEndOfWord bool
}

func newNode() *Trie {
	return &Trie{
		children:    make([]*Trie, 256),
		isEndOfWord: false,
	}
}

// NewTrie instanciates a new Trie
func NewTrie() *Trie {
	return newNode()
}

func insertRecurse(word string, indx int, node *Trie) {
	if indx > len(word)-1 {
		return
	}

	char := word[indx]
	//fmt.Println("Inserting...", string(char))
	if node.children[char] == nil {
		//fmt.Println("Char not found at root..")
		node.children[char] = newNode()
	}

	if indx == len(word)-1 {
		//fmt.Println("Marking end of word..")
		node.children[char].isEndOfWord = true
	}

	insertRecurse(word, indx+1, node.children[char])
}

// Insert adds a the word into Trie
func (trie *Trie) Insert(word string) {

	// We start at the root, and check to see if the pointer at the chars
	// numeric index is set or not. This can be done in constant time
	insertRecurse(word, 0, trie)
}

func searchRecurse(word string, indx int, node *Trie) bool {
	if indx > len(word)-1 {
		return false
	}

	c := word[indx]
	if node.children[c] == nil {
		return false
	}

	if indx == len(word)-1 && node.children[c].isEndOfWord {
		return true
	}

	return searchRecurse(word, indx+1, node.children[c])
}

// Search returns true is word exists in Trie
func (trie *Trie) Search(word string) bool {
	if word == "" {
		return false
	}

	return searchRecurse(word, 0, trie)
}

/*
func (trie *Trie) Print() {
	for i, c := range trie.children {
		if c != nil {
			fmt.Println("character:", string(i), " IsEndofWord:", c.isEndOfWord, " child:", c)
		}
	}
}
*/

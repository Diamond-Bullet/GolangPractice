package data_structures

// Trie 前缀树
type Trie struct {
	isEnd    bool
	Children [26]*Trie
}

// Insert /** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	pre := t
	for _, c := range word {
		if pre.Children[c-97] == nil {
			pre.Children[c-97] = &Trie{}
		}
		pre = pre.Children[c-97]
	}
	pre.isEnd = true
}

// Search /** Returns if the word is in the trie. */
func (t *Trie) Search(word string) bool {
	pre := t
	for _, c := range word {
		if pre.Children[c-97] == nil {
			return false
		}
		pre = pre.Children[c-97]
	}
	return pre.isEnd
}

// StartsWith /** Returns if there is any word in the trie that starts with the given prefix. */
func (t *Trie) StartsWith(prefix string) bool {
	pre := t
	for _, c := range prefix {
		if pre.Children[c-97] == nil {
			return false
		}
		pre = pre.Children[c-97]
	}
	return true
}

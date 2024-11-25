package basics

const (
	letterSum = 26
	aASCII    = 'a'
)

// Trie # here we implement a trie based on lowercase.
// you can simply turn it into a more general form by replacing '[letterSum]*Trie' to 'map[string]*Trie'
// and might change some corresponding details in the codes.
type Trie struct {
	isEnd    bool
	Children [letterSum]*Trie
}

// Insert # Inserts a word into the trie.
func (t *Trie) Insert(word string) {
	pre := t
	for _, c := range word {
		if pre.Children[c-aASCII] == nil {
			pre.Children[c-aASCII] = &Trie{}
		}
		pre = pre.Children[c-aASCII]
	}
	pre.isEnd = true
}

// Search # Returns if the word is in the trie.
func (t *Trie) Search(word string) bool {
	pre := t
	for _, c := range word {
		if pre.Children[c-aASCII] == nil {
			return false
		}
		pre = pre.Children[c-aASCII]
	}
	return pre.isEnd
}

// StartsWith # Returns if there is any word in the trie that starts with the given prefix.
func (t *Trie) StartsWith(prefix string) bool {
	pre := t
	for _, c := range prefix {
		if pre.Children[c-aASCII] == nil {
			return false
		}
		pre = pre.Children[c-aASCII]
	}
	return true
}

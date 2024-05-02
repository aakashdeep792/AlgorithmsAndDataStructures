package binaryTree

/*
208. Implement Trie (Prefix Tree)
Solved
Medium
Topics
Companies
A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and retrieve keys in a dataset of strings. There are various applications of this data structure, such as autocomplete and spellchecker.

Implement the Trie class:

Trie() Initializes the trie object.
void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true if there is a previously inserted string word that has the prefix prefix, and false otherwise.


Example 1:

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]

Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

*/

type Trie struct {
	child []*Trie
	end   bool
}

func Constructor() Trie {
	return Trie{
		child: make([]*Trie, 26),
	}
}

func (this *Trie) Insert(word string) {
	root := this
	for i := 0; i < len(word); i++ {
		if root.child[word[i]-'a'] == nil {
			tmp := Constructor()
			root.child[word[i]-'a'] = &tmp
			root = &tmp
		} else {
			root = root.child[word[i]-'a']
		}

	}
	root.end = true
}

func (this *Trie) Search(word string) bool {
	root := this
	for i := 0; i < len(word); i++ {
		if root.child[word[i]-'a'] == nil {
			return false
		}
		root = root.child[word[i]-'a']
	}

	return root.end
}

func (this *Trie) StartsWith(prefix string) bool {
	root := this
	for i := 0; i < len(prefix); i++ {
		if root.child[prefix[i]-'a'] == nil {
			return false
		}
		root = root.child[prefix[i]-'a']
	}

	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

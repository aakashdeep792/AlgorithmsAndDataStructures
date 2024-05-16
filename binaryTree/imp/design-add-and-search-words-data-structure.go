package binaryTree

/*
211. Design Add and Search Words Data Structure

Design a data structure that supports adding new words and finding if a string matches any previously added string.

Implement the WordDictionary class:

WordDictionary() Initializes the object.
void addWord(word) Adds word to the data structure, it can be matched later.
bool search(word) Returns true if there is any string in the data structure that matches word or false otherwise. word may contain dots '.' where dots can be matched with any letter.


Example:

Input
["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
Output
[null,null,null,null,false,true,true,true]

Explanation
WordDictionary wordDictionary = new WordDictionary();
wordDictionary.addWord("bad");
wordDictionary.addWord("dad");
wordDictionary.addWord("mad");
wordDictionary.search("pad"); // return False
wordDictionary.search("bad"); // return True
wordDictionary.search(".ad"); // return True
wordDictionary.search("b.."); // return True


Constraints:

1 <= word.length <= 25
word in addWord consists of lowercase English letters.
word in search consist of '.' or lowercase English letters.
There will be at most 2 dots in word for search queries.
At most 104 calls will be made to addWord and search.

*/

type WordDictionary struct {
	child []*WordDictionary
	end   bool
}

func NewNode() *WordDictionary {
	tmp := Constructor1()
	return &tmp
}

func Constructor1() WordDictionary {
	return WordDictionary{
		child: make([]*WordDictionary, 26),
		end:   false,
	}
}

func (this *WordDictionary) AddWord(word string) {
	ptr := this
	for i := 0; i < len(word); i++ {
		tmp := word[i] - 'a'

		ch := ptr.child[tmp]
		if ch == nil {
			// create the new node and update the child node value in the parent
			ch = NewNode()
			ptr.child[tmp] = ch
		}
		ptr = ch
	}
	ptr.end = true
}

func (this *WordDictionary) Search(word string) bool {
	ptr := this
	for i := 0; i < len(word); i++ {
		// for the . case
		if word[i] == '.' {
			ans := false
			for j := 0; j < len(ptr.child); j++ {
				chTmp := ptr.child[j]
				if chTmp != nil {
					ans = ans || chTmp.Search(word[i+1:])
				}
			}
			return ans
		}
		// for other cases
		tmp := word[i] - 'a'
		ch := ptr.child[tmp]
		if ch == nil {
			return false
		}
		ptr = ch
	}

	if ptr != nil {
		return ptr.end
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

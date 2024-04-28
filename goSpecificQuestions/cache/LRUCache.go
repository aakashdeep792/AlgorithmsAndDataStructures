package cache

import (
	"container/list"
	"fmt"
	"hash/fnv"
)

/*

items map[int64]*list.List ={
	1:&list.List{
		root:list.Element{
			Value:&list.Element{  // this are same
				Value: &Item{
					key : "name"
					value : "aakash"
				}
			}},
		len:9
}}

c.list = &list.List{
	root:list.Element{           // this are same
		Value: &Item{
			key : "name"
			value : "aakash"
		}
	},
	len:12
}

*/

// lru, ttl and most frequently used

type KeyHash uint64
type LRUCache struct {
	items    map[KeyHash]*list.List
	list     *list.List
	capacity uint64
}

func NewLRUCache(capacity, ttl uint64) *LRUCache {

	return &LRUCache{
		items:    make(map[KeyHash]*list.List),
		list:     list.New(),
		capacity: capacity,
	}
}

type Item struct {
	key   any
	value any
}

func hash(key any) KeyHash {
	h := fnv.New64()
	_, err := h.Write([]byte(fmt.Sprintf("#%v", key)))
	if err != nil {
		fmt.Errorf("hash failed for key:%v", key)
	}
	return KeyHash(h.Sum64())

}

func (c *LRUCache) GetVal(key any) any {
	kh := hash(key)
	if _, ok := c.items[kh]; !ok {
		return nil
	}
	l := c.items[kh]
	for e := l.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*list.Element) // point to element in the list
		i := tmp.Value.(*Item)
		if i.key == key {
			c.list.MoveToFront(tmp)
			return i.value
		}

	}

	return nil
}

func (c *LRUCache) SetVal(key, value any) {
	// kh := KeyHash(hash(key))
	// if _, ok := c.items[kh]; !ok {
	// 	c.items[kh] = list.New()
	// }
	// // list start contains the newer  element and end contains the older element
	// // insert the element in the list
	// it := c.list.PushFront(&Item{key: key, value: value})
	// // insert the same element in the same element in hash map
	// c.items[kh].PushFront(it)

	c.update(key, value)

	if c.list.Len() > int(c.capacity) {
		// this value needs to be deleted
		i := c.list.Back()
		k := i.Value.(*Item).key
		keyHash := hash(k)
		// find the element in the hastmap that need to be deleted
		for e := c.items[keyHash].Front(); e != nil; e = e.Next() {
			tmp := e.Value.(*list.Element)
			if tmp == i {
				c.items[keyHash].Remove(e)
				// Remove the element from the list
				c.list.Remove(i)
				break
			}
		}
	}
}

func (c *LRUCache) update(key, value any) {
	kh := hash(key)
	l := c.items[kh]
	if l == nil {
		c.items[kh] = list.New()
		l = c.items[kh]
	}

	// treverse the map list, and find if key already exit
	for e := l.Front(); e != nil; e = e.Next() {
		tmp := e.Value.(*list.Element)
		i := tmp.Value.(*Item)
		if i.key == key { // update the value
			i.value = value
			c.list.MoveToFront(tmp) // move the updated entry to front
			return
		}
	}

	// list start contains the newer  element and end contains the older element
	// insert the element in the list
	it := c.list.PushFront(&Item{key: key, value: value})
	// insert the same element in the same element in hash map
	c.items[kh].PushFront(it)

}

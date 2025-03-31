// Final Commit Done

package advancedawesome

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	head     *list.List
}

type Node struct {
	key, value int
}

var myCache = LRUCache{
	capacity: 2,
	cache:    make(map[int]*list.Element),
	head:     list.New(),
}

func (lc *LRUCache) Get(k int) int {
	if element, present := myCache.cache[k]; present {
		myCache.head.MoveToFront(element)
		return element.Value.(*Node).value
	}
	return -1
}

func (lc *LRUCache) Put(k, v int) {
	if element, present := myCache.cache[k]; present {
		element.Value.(*Node).value = v
		myCache.head.MoveToFront(element)
		return
	}

	if myCache.head.Len() >= myCache.capacity {
		oldestElement := myCache.head.Back()
		delete(myCache.cache, oldestElement.Value.(*Node).key)
		myCache.head.Remove(oldestElement)
	}

	newNode := &Node{key: k, value: v}
	newElement := myCache.head.PushFront(newNode)
	myCache.cache[k] = newElement
}

func LRUCacheStart() {
	fmt.Println("\nImplement LRU Cache Using Doubly Linked List and Hashing")

	myCache.Put(1, 1)
	myCache.Put(2, 2)

	fmt.Print(myCache.Get(1), " ")

	myCache.Put(3, 3)

	fmt.Print(myCache.Get(2), " ")

	myCache.Put(4, 4)

	fmt.Print(myCache.Get(1), " ")
	fmt.Print(myCache.Get(3), " ")
	fmt.Print(myCache.Get(4), " ")

	fmt.Println("\nEnding LRU Cache")
}

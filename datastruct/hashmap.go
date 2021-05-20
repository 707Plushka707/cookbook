package main

import "fmt"

const arraySize  = 100

type HashTable struct {
    array [arraySize]*bucket
}

type bucket struct {
	head *bucketNode
}

// BucketNode is a linked list node that holds a keys
type bucketNode struct {
	key  string
	next *bucketNode
}

// insert will be a key and insert it to HashTable
func (h *HashTable) Insert(key string) {
	index := hash(key)
    fmt.Println("index is : ", index)
	h.array[index].insert(key)
}

// search will be search for element by key.
func (h *HashTable) search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

func (h *HashTable) delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// hash
func hash(key string) int {
	var sum int32 = 0
	for _, v := range key {
		sum += v
	}

	return int(sum) % 100 //arraySize
}

// Init will create a bucket in each slot of the hashtable
func Init() *HashTable {
	result := &HashTable{}
    for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// insert will create a node and insert it in bucket by key
func (b *bucket) insert(key string) {
	if !b.search(key) {
		newNode := &bucketNode{key: key}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(key, " already exists")
	}

}

// search take a key and return true if key is exist.
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete delete key from bucket
func (b *bucket) delete(key string) {

	if b.head.key == key {
		b.head = b.head.next
		return
	}
	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

func main() {

    hashTable := Init()
    list := []string{"eric", "jawad", "hamid", "ghanim", "borita", "ja3tot"}
    for _,  v := range list {
        hashTable.Insert(v)
    }
    //fmt.Println(hashTable)
    fmt.Println(hashTable.search("ja3tot"))
    hashTable.delete("ja3tot")
    fmt.Println(hashTable.search("ja3tot"))


}

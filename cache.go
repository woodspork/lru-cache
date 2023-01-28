package main

import "fmt"

type cache struct {
  linkedList *doublyLinkedList
  hashMap map[string]*node
  capacity int
}

func initCache() *cache {
  linkedList := initDLL()
  hashMap := make(map[string]*node)
  cache := cache{
    linkedList: linkedList,
    hashMap: hashMap,
    capacity: 10,
  }

  return &cache
}

func (c *cache) cacheLength() int {
  return c.linkedList.len
}

func (c *cache) cacheAtCapacity() bool {
  return c.linkedList.len == c.capacity
}

func (c *cache) cacheHit(data string) {
  node := c.hashMap[data]
  c.linkedList.moveNodeToTail(node)

  return
}

func (c *cache) cacheMiss(data string) {
  cacheAtCapacity := c.cacheAtCapacity()
  linkedList := c.linkedList

  if cacheAtCapacity == true {
    linkedList.removeHead()
    linkedList.appendToTail(data)
    c.hashMap[data] = linkedList.tail
  } else {
    linkedList.appendToTail(data)
    c.hashMap[data] = linkedList.tail
  }

  return
}

func (c *cache) cacheAlter(data string) {
  if c.hashMap[data] == nil {
    c.cacheMiss(data)
  } else {
    fmt.Println("cacheHit")
    c.cacheHit(data)
  }

  return
}


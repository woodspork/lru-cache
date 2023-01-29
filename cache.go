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

func (c *cache) cacheHit(n *node) {
  // node := c.hashMap[data
  c.linkedList.moveNodeToTail(n)

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

func (c *cache) set(data string) {
  c.linkedList.appendToTail(data)
  c.hashMap[data] = c.linkedList.tail

  return
}

func (c *cache) get(data string) {
  node := c.hashMap[data]
  nodeExists := node != nil
  if nodeExists == true {
    c.cacheHit(node)
  } else {
    c.cacheMiss(data)
  }


  return
}

func (c *cache) delete(data string) {
  node := c.hashMap[data]
  nodeExists := node != nil

  if nodeExists == false {
    return
  } else {
    c.linkedList.removeNode(node)
    c.hashMap[data] = nil
  }

  return
}










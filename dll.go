package main

import "fmt"

type node struct {
  data string
  prev *node
  next *node
}

type doublyLinkedList struct {
  len int
  tail *node
  head *node
}

func initDLL() *doublyLinkedList {
  return &doublyLinkedList{}
}

func (d *doublyLinkedList) traverseLinkedList() {
  currNode := d.head

  for currNode != nil {
    fmt.Println("data", currNode.data)
    currNode = currNode.next
  }

  return
}

// removing the least recently used item in the cache
func (d *doublyLinkedList) removeHead() {
  head := d.head
  if head.next == nil {
    return
  } else {
    head.next.prev = nil
    head.next = nil
  }

  return
}

func (d *doublyLinkedList) moveNodeToTail (n *node) {

  if n.next == nil {
    return
  } else if n.prev == nil {
    d.head = n.next
    n.next.prev = nil
    n.next = nil
    n.prev = d.tail
    d.tail.next = n
    d.tail = n
  } else {
    n.prev.next = n.next
    n.next.prev = n.prev
    n.next = nil
    n.prev = d.tail
    d.tail.next = n
    d.tail = n
  }


  return
}

func (d *doublyLinkedList) appendToHead(data string) {
  newNode := &node{
    data: data,
  }

  if d.head == nil {
    d.head = newNode
    d.tail= newNode
  } else {
    newNode.next = d.head
    d.head.prev = newNode
    d.head = newNode
  }

  d.len++

  return
}

func (d *doublyLinkedList) removeNode(n *node) {
  if n.next == nil && n.prev == nil {
    return
  } else if n.next == nil {
    n.prev.next = nil
    d.tail = n.prev
    n.prev = nil
  } else if n.prev == nil {
    n.next.prev = nil
    d.head = n.next
    n.next = nil
  } else {
    n.prev.next = n.next
    n.next.prev = n.prev
    n.next = nil
    n.prev = nil
  }

  return

}

func (d *doublyLinkedList) appendToTail(data string) {
  newNode := &node{
    data: data,
  }

  if d.tail == nil {
    d.tail = newNode
    d.head = newNode
  } else {
    newNode.prev = d.tail
    d.tail.next = newNode
    d.tail = newNode
  }

  d.len ++

  return
}

func (d *doublyLinkedList) SizeOfDLL() int {
  return d.len
}

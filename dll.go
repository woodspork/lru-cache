package main

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

package main

import (
  "fmt"
)

func main(){
  dll := initDLL()
  dll.appendToHead("2")
  fmt.Println(dll)
  fmt.Println("length of the doubly linked list", dll.len)
  dll.appendToTail("4")
  fmt.Println("length of the doubly linked list", dll.len)
}

package main

import (
  "fmt"
)

func main(){
  cache := initCache()
  fmt.Println("cache", cache)
  cache.cacheAlter("string")
  cache.cacheAlter("string1")
  cache.cacheAlter("string2")
  cache.cacheAlter("string3")
  cache.cacheAlter("string4")
  cache.cacheAlter("string5")
  cache.cacheAlter("string6")
  cache.cacheAlter("string7")
  cache.cacheAlter("string8")
  cache.cacheAlter("string9")
}

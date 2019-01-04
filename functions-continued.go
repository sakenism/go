package main

import (
       "fmt"
)

func add (x, y int) int {
     return x + y
}

func main () {
     fmt.Printf("%d\n", add(12, 43))
}
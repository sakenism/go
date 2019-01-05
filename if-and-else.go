package main

import (
       "fmt"
       "math"
)

func pow(x, n, lim float64) float64 {
     if(math.Pow(x, n) < lim) {
     	return math.Pow(x, n)
     }
     fmt.Printf("%g >= %g\n", math.Pow(x, n), lim)
     return lim
}

func main() {
     fmt.Println(pow(3, 2, 10))
     fmt.Println(pow(3, 3, 20))
}
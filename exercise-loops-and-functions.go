package main

import "fmt"

const Eps = 1e-8

func Abs(x float64) float64 {
     if(x < 0) {
     	  return -x
     }
     return x
}

func Sqrt(x float64) float64 {
     var l, r, mid float64
     r = x
     mid = (l + r) / 2
     if (x == 1.0 || x == 0.0) {
     	return (x)
     }
     for (Abs(mid * mid - x) > Eps) {
     	 fmt.Printf("%g\n", mid)
     	 mid = (l + r) / 2
	 if (mid * mid > x) {
	    r = mid 
	 }	 
	 if (mid * mid < x) {
	    l = mid
	 }
     }
     return mid
}    

func main() {
     fmt.Println(Sqrt(0.0002))
}
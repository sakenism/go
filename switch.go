package main

import (
       "fmt"
       "runtime"
)

func main() {
     fmt.Printf("Go runs on ")
     switch os := runtime.GOOS; os {
     	    case "darwin" :
	    	 fmt.Printf("OS X.\n")
	    case "linux" :
	    	 fmt.Printf("Linux.\n")
	    default :
	    	 fmt.Printf("%s.", os)
     }
     os := runtime.GOOS
     fmt.Println(os)
}
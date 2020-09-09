package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
        t := time.Now().UnixNano()
        rand.Seed(t)
	s := rand.Intn(6)
	switch s + 1 {
	  case 6:
	    fmt.Printf("大吉")
	  case 4,5:
	    fmt.Printf("中吉")
	  case 2,3:
	    fmt.Printf("吉")
	  case 1:
	    fmt.Printf("凶")
	}
}

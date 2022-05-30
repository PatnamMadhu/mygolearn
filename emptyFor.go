package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := time.After(time.Second * 5)
	for {
		select {
		case <-timeout:
			fmt.Println("timeout")
			return
		}
	}
}
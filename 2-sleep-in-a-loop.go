package main

import (
	"fmt"
	"time"
)

func sleepInALoop() {
	for i := 1; i <= MAX; i++ {
		fmt.Printf("%d - Starting... \n", i)
		time.Sleep(TIME * time.Second)
		fmt.Printf("%d - It's over. \n", i)
	}
}

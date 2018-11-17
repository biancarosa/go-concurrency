package main

import (
	"fmt"
	"time"
)

func sleepInALoop() {
	for i := 1; i <= Max; i++ {
		fmt.Printf("%d - Starting... \n", i)
		time.Sleep(Time * time.Second)
		fmt.Printf("%d - It's over. \n", i)
	}
}

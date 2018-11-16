package main

import (
	"fmt"
	"time"
)

func justSleep() {
	fmt.Println("Starting...")
	time.Sleep(TIME * time.Second)
	fmt.Println("It's over.")
}

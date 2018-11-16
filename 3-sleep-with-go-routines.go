package main

import (
	"fmt"
	"time"
)

func sleepWithGoRoutines() {
	channel := make(chan string, MAX)
	for i := 1; i <= MAX; i++ {
		go func(counter int) {
			fmt.Printf("%d - Starting... \n", counter)
			time.Sleep(TIME * time.Second)
			channel <- fmt.Sprintf("%d - It's over. \n", counter)
		}(i)
	}
	for i := 1; i <= MAX; i++ {
		fmt.Printf(<-channel)
	}
}

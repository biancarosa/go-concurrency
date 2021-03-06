package main

import (
	"fmt"
	"time"
)

func sleepWithGoRoutines() {
	channel := make(chan string, Max)
	for i := 1; i <= Max; i++ {
		go func(counter int) {
			fmt.Printf("%d - Starting... \n", counter)
			time.Sleep(Time * time.Second)
			channel <- fmt.Sprintf("%d - It's over. \n", counter)
		}(i)
	}
	for i := 1; i <= Max; i++ {
		fmt.Printf(<-channel)
	}
}

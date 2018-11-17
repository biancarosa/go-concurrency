package main

import (
	"fmt"
)

func loop() {
	for i := 1; i <= MAX_BIGGER; i++ {
		fmt.Printf("%d - Starting... \n", i)
		//code without I/O
		fmt.Printf("%d - It's over. \n", i)
	}
}

func loopWithGoRoutines() {
	channel := make(chan string, MAX_BIGGER)
	for i := 1; i <= MAX_BIGGER; i++ {
		go func(counter int) {
			fmt.Printf("%d - Starting... \n", counter)
			//code without I/O
			channel <- fmt.Sprintf("%d - It's over. \n", counter)
		}(i)
	}
	for i := 1; i <= MAX_BIGGER; i++ {
		fmt.Printf(<-channel)
	}
}

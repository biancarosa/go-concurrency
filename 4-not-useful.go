package main

import (
	"fmt"
)

func loop() {
	for i := 1; i <= MaxBigger; i++ {
		fmt.Printf("%d - Starting... \n", i)
		fmt.Printf("%d - It's over. \n", i)
	}
}

func loopWithGoRoutines() {
	channel := make(chan string, MaxBigger)
	for i := 1; i <= MaxBigger; i++ {
		go func(counter int) {
			fmt.Printf("%d - Starting... \n", counter)
			channel <- fmt.Sprintf("%d - It's over. \n", counter)
		}(i)
	}
	for i := 1; i <= MaxBigger; i++ {
		fmt.Printf(<-channel)
	}
}

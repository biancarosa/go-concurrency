package main

import (
	"fmt"
	"time"
)

func raceOnALoop() {
	channel := make(chan string, Max)
	for i := 1; i <= Max; i++ {
		go func() {
			fmt.Printf("%d - Starting... \n", i) // this is not the the i you want
			time.Sleep(Time * time.Second)
			channel <- fmt.Sprintf("%d - It's over. \n", i) //this is also not the i you want
		}()
	}
	for i := 1; i <= Max; i++ {
		fmt.Printf(<-channel)
	}
}

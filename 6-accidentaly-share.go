package main

import (
	"fmt"
	"os"
	"time"
)

func accidentalyShare() error {
	channel := make(chan string, Max)
	f1, err := os.Create("file1")
	if err != nil {
		return err
	}
	for i := 1; i <= Max; i++ {
		go func(counter int) {
			fmt.Printf("%d - Starting... \n", counter)
			time.Sleep(Time * time.Second)
			_, err = f1.Write([]byte(fmt.Sprintf("Line %d\n", counter)))
			if err != nil {
				fmt.Println("error")
			}
			channel <- fmt.Sprintf("%d - It's over. \n", counter)
		}(i)
	}
	for i := 1; i <= Max; i++ {
		fmt.Printf(<-channel)
	}
	return nil
}

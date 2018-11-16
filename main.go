package main

import (
	"fmt"
	"time"
)

const (
	//MAX defines how many loops are going to run
	MAX int = 5
	//TIME defines how many seconds we will wait
	TIME time.Duration = time.Duration(1)
)

func main() {
	fmt.Println("Hello")
}

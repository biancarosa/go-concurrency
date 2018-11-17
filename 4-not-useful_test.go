package main

import "testing"

func TestNotUsefulLoop(t *testing.T) {
	loop()
}

func TestNotUsefulLoopWithGoRoutines(t *testing.T) {
	loopWithGoRoutines()
}

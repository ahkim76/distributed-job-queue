package main

import "fmt"

// EXECUTABLE THAT  STARTS THE WORKER LOOP
// Connects to database, starts a loop of claiming job, running handler, updating job row in DB
// Also starts the sweeper (goroutine) to reclaim stuck jobs

func main() {
	fmt.Println("WHATS UP GUYS!")
}
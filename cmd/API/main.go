package main

import "fmt"

// EXECUTABLE THAT STARTS THE HTTP SERVER
// Connects to database, creates HTTP router, calls http.ListenAndServe(":8080", router)
// No code logic; wires things together

// POST /jobs -> enqueue a job

// GET /jobs -> get all job statuses

// GET /jobs/:id -> get job status

// GET /stats -> queue depths

func main() {
	fmt.Println("Hello world!!!")
}
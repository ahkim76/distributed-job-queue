package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/ahkim76/distributed-job-queue/internal/jobs"
	"github.com/gin-gonic/gin"
)

// EXECUTABLE THAT STARTS THE HTTP SERVER
// Connects to database, creates HTTP router, calls http.ListenAndServe(":8080", router)
// No code logic; wires things together

var errMsg = "timeout contacting external API"
var lease = time.Now().Add(30 * time.Second)

var jobStore = []jobs.Job{
	{
		ID:          1, QueueName:   "default", JobType:     "log_message", Payload:     json.RawMessage(`{"msg":"hello world"}`),
		Status:      "pending", Attempts:    0, MaxAttempts: 5,
		Priority:    0, VisibleAt:   time.Now(), CreatedAt:   time.Now(), UpdatedAt:   time.Now(),
	},
	{
		ID:        2, QueueName: "emails", JobType:   "send_email", Payload:   json.RawMessage(`{"to":"test@example.com","subject":"Welcome!"}`),
		Status:      "pending", Attempts:    1, MaxAttempts: 3, Priority:    10,
		VisibleAt: time.Now(), CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},
	{
		ID:        3, QueueName: "payments", JobType:   "charge_card", Payload:   json.RawMessage(`{"user_id":42,"amount":1999}`),
		Status:         "processing", Attempts:       2, MaxAttempts:    5, Priority:       5,
		VisibleAt:      time.Now(), LeaseExpiresAt: &lease, LastError:      &errMsg,
		CreatedAt: time.Now(), UpdatedAt: time.Now(),
	},

}

// POST /jobs -> enqueue a job
func postJobs(c *gin.Context) {
	var newJob jobs.Job
	
	// BindJSON reads HTTP request body, parses JSON, writes parsed fields into the memory at &newJob
	err := c.BindJSON(&newJob)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	jobStore = append(jobStore, newJob)
	c.IndentedJSON(http.StatusOK, newJob)
}

// GET /jobs -> get all jobs
func getJobs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobStore)
}

// GET /jobs/:id -> get a job
func getJobByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid job ID"})
		return
	}
	for _, a := range jobStore {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "job not found"})
}

// DELETE /jobs/:id -> deletes job with id
func deleteJob(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid ID value"})
	}

	for i, a := range jobStore {
		if a.ID == id {
			jobStore = append(jobStore[:i], jobStore[i+1:]...)
			c.Status(http.StatusNoContent)
			return 
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"error": "job not found"})
}

// GET /stats -> queue depths
func getStats(c *gin.Context) {
	// initialize an accumulator
	statMap := make(map[string]map[string]int)

	// loop over jobStore and increment respective (queue_name, status) pair
	for _, a := range jobStore {
		queue_name := a.QueueName
		status := a.Status

		// check if queue_name exists
		if _, ok := statMap[queue_name]; !ok {
			statMap[queue_name] = make(map[string]int)
		} 

		// check status
		current_queue := statMap[queue_name]
		if _, ok := current_queue[status]; !ok {
			current_queue[status] = 0
		} 
		current_queue[status]++
		statMap[queue_name] = current_queue
	}

	// convert map -> slice for JSON response
	type StatRow struct {
		QueueName string `json:"queue_name"`
		Status string `json:"status"`
		Count int `json:"count"`
	}

	stats := []StatRow{}

	for queue, inner := range statMap {
		for status, count := range inner {
			stats = append(stats, StatRow{
				QueueName: queue,
				Status: status,
				Count: count,
			})
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"stats": stats})
}

func main() {
	fmt.Println("Hello world!!!")

	// Initialize HTTP router with default middleware
	router := gin.Default()

	// Endpoints
	router.GET("/jobs", getJobs)
	router.GET("/jobs/:id", getJobByID)
	router.POST("/jobs", postJobs)
	router.GET("/jobs/status", getStats)
	router.DELETE("jobs/:id", deleteJob)

	// Run HTTP server
	router.Run("localhost:8080")
}


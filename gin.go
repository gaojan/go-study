package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Agent ...
type Agent struct {
	AgentID  string `json:"agent_id"`
	QueuedAt string `json:"queued_at"`
	QueuedBy string `json:"queued_by"`
}

// Details ...
type Details struct {
	EventID  string `json:"event_id"`
	Endpoint string
	Metric   string
	Content  string
	Priority int
	Status   string
}

// Test1 ...
type Test1 struct {
	Agent       Agent
	Details     Details
	Description string
	EventType   string `json:"event_type"`
	ServiceKey  string `json:"service_key"`
}

// Test2 test2
type Test2 struct {
	Data []*Test1
}

func main() {
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/v1/test", func(c *gin.Context) {
		var test Test1

		if err := c.BindJSON(&test); err == nil {
			log.Println("========================start io=====================")
			time.Sleep(time.Duration(1) * time.Second)
			log.Println("=========================end io=======================")
			c.JSON(http.StatusOK, test)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	})
	r.POST("/v2/test", func(c *gin.Context) {
		var test Test2

		if err := c.BindJSON(&test); err == nil {
			log.Println("========================start io=====================")
			time.Sleep(time.Duration(1) * time.Second)
			log.Println("=========================end io=======================")
			c.JSON(http.StatusOK, test)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	r.POST("/v3/test", func(c *gin.Context) {
		var test Test2

		if err := c.BindJSON(&test); err == nil {
			log.Println("========================start io=====================")
			time.Sleep(time.Duration(1) * time.Second)
			log.Println("=========================end io=======================")
			c.JSON(http.StatusOK, test)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	})
	r.POST("/v4/test", func(c *gin.Context) {
		var test Test2

		if err := c.BindJSON(&test); err == nil {
			log.Println("========================start io=====================")
			time.Sleep(time.Duration(1) * time.Second)
			log.Println("=========================end io=======================")
			c.JSON(http.StatusOK, test)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

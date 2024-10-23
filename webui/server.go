package webui

import (
	"fmt"
	"net/http"

	// "os"
	"test-in-go/utils/report_helpers"

	"github.com/gin-gonic/gin"
)

// StartWebServer starts the web server for the UI
func StartWebServer() {
	router := gin.Default()

	// Serve the static HTML files
	router.Static("/static", "./webui/static")
	router.LoadHTMLGlob("webui/templates/*")

	// API for running tests
	router.POST("/api/run-tests", func(c *gin.Context) {
		err := report_helpers.InitPrettyReport()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Error initializing report"})
			return
		}

		status := runTests()
		report_helpers.FinalizePrettyReport()

		if status != 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Tests failed"})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "Tests passed successfully"})
		}
	})

	// Serve the index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// Start the server
	router.Run(":8080")
}

// runTests is a helper function to run tests programmatically
func runTests() int {
	// Here, you can call the main test function or godog to run tests
	fmt.Println("Running tests...")
	return 0 // Return test status
}

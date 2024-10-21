package main

import (
	"fmt"
	"os"

	"test-in-go/config"
	"test-in-go/utils/logging_helpers"

	"github.com/cucumber/godog"
)

func main() {
	// Set up logging
	log := logging_helpers.SetupLogger()

	// Load environment variables
	env := config.GetEnv("ENVIRONMENT", "dev")
	log.Infof("Running in %s environment", env)

	// Define options for the test run
	opts := godog.Options{
		Format:    "pretty", // Use a human-readable format
		Paths:     []string{"./features"},
		Output:    os.Stdout,
		Randomize: 0, // Don't randomize the test order
	}

	// Run Godog tests
	status := godog.TestSuite{
		Name:                 "Test-in-Go",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if status != 0 {
		log.Error("Test execution failed.")
	} else {
		log.Info("Test execution completed successfully.")
	}

	os.Exit(status)
}

// InitializeTestSuite - this can be used to prepare data, etc.
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	fmt.Println("Initializing the test suite...")
	// Any necessary suite-level setup goes here
}

// InitializeScenario - this runs before each test scenario
func InitializeScenario(ctx *godog.ScenarioContext) {
	fmt.Println("Initializing scenario...")
	// Initialize step definitions here
}

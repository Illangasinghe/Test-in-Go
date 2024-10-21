package main

import (
	// "log"
	"os"
	"test-in-go/config"
	"test-in-go/steps/inbound"
	"test-in-go/utils/db_helpers"
	"test-in-go/utils/logging_helpers"

	// "testing"

	"github.com/cucumber/godog"
	// "github.com/ozontech/allure-go/pkg/allure"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func main() {

	// Set up logger
	logger = logging_helpers.SetupLogger()

	// Set up the database connection
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Error loading configuration: ", err)
		os.Exit(1)
	}

	// Connect to the database using the configuration
	err = db_helpers.ConnectPostgres(cfg.DB.URL)
	if err != nil {
		logger.Fatal("Failed to connect to the database: ", err)
		os.Exit(1)
	}
	defer db_helpers.ClosePostgres()

	// Run the test suite
	status := runGodogTests()

	// Exit with appropriate status code
	if status != 0 {
		logger.Error("Test suite failed.")
	} else {
		logger.Info("Test suite completed successfully.")
	}
	os.Exit(status)
}

func runGodogTests() int {
	// Godog test options
	opts := godog.Options{
		Format: "pretty",
		Paths:  []string{"./features/inbound/product_creation.feature"},
		Output: os.Stdout,
		Strict: true,
	}

	// Run the test suite and return the status
	return godog.TestSuite{
		Name:                 "Product Creation Test",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()
}

// InitializeTestSuite - this can be used to prepare data, etc.
func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	// Any necessary suite-level setup goes here
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	inbound.InitializeProductSteps(ctx)
}

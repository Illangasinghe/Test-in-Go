package main

import (
	"os"
	"test-in-go/config"
	"test-in-go/steps/inbound"
	"test-in-go/utils/db_helpers"
	"test-in-go/utils/logging_helpers"
	"test-in-go/utils/report_helpers"

	"github.com/cucumber/godog"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func main() {

	// Set up logger
	logger = logging_helpers.SetupLogger()

	// Initialize the pretty report
	err := report_helpers.InitPrettyReport()
	if err != nil {
		logger.Fatal("Error initializing pretty report: ", err)
		os.Exit(1)
	}

	// Set up the database connection
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("Error loading configuration: ", err)
		os.Exit(1)
	}

	// Connect to the database using the configuration
	err = db_helpers.ConnectPostgres(cfg.DB.URL)
	if err != nil {
		logger.Fatal("Failed to connect to the database "+config.GetEnv("DB_URL")+" | "+cfg.DB.URL+" due to:", err)
		os.Exit(1)
	}
	defer db_helpers.ClosePostgres()

	// Run the test suite
	status := runGodogTests()

	// Finalize the report without passing parameters
	err = report_helpers.FinalizePrettyReport()
	if err != nil {
		logger.Error("Error finalizing the report: ", err)
	}

	// Exit with appropriate status code
	if status != 0 {
		logger.Error("Test suite failed.")
	} else {
		logger.Info("Test suite completed successfully.")
	}
	os.Exit(status)
}

func runGodogTests() int {
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

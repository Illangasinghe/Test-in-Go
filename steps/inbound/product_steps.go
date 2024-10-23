package inbound

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-in-go/utils/data_helpers"
	"test-in-go/utils/db_helpers"
	"test-in-go/utils/protocol_helpers"
	"test-in-go/utils/report_helpers"

	"github.com/cucumber/godog"
)

// Declare a global variable for the product ID.
var createdProductID string

// Helper function to handle API requests, now taking the Root structure that includes a products array.
func postProductToAPI(root data_helpers.Root) (*http.Response, error) {
	return protocol_helpers.PostRequest("/product", root)
}

// Step 1: Define a new test case with a dynamic test code based on the test case ID.
func aNewTestcaseWithID(testcaseID string) error {
	stepName := "Define a new test case with dynamic TestCode"
	report_helpers.PrettyLogStep(stepName, "Started", fmt.Sprintf("Testcase ID: %s", testcaseID))

	err := data_helpers.GenerateTestCode(testcaseID)
	if err != nil {
		report_helpers.FailedStep()
		report_helpers.PrettyLogStep(stepName, "Failed", fmt.Sprintf("Error: %v", err))
		return err
	}

	report_helpers.PassedStep()
	report_helpers.PrettyLogStep(stepName, "Passed", "TestCode generated successfully")
	return nil
}

// Step 2: Create the product with the given description using the previously generated TestCode, wrapped inside a Root structure.
func aProductWithTheDescriptionIsCreated(description string) error {
	stepName := "Create product with dynamic TestCode"
	report_helpers.PrettyLogStep(stepName, "Started", fmt.Sprintf("Description: %s", description))

	// Build a product using the builder pattern.
	product := data_helpers.NewProductBuilder().
		WithTestCode().                      // Use the generated TestCode.
		WithShortDescription(description).   // Set the description dynamically.
		WithSKU(data_helpers.GenerateSKU()). // Generate dynamic SKU.
		Build()

	// Wrap the product inside a Root structure (with products array).
	root := data_helpers.Root{
		Products: []data_helpers.Product{product},
	}

	// Marshal the Root structure to JSON for sending in API requests.
	jsonData, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		report_helpers.FailedStep()
		report_helpers.PrettyLogStep(stepName, "Failed", fmt.Sprintf("Failed to marshal product: %v", err))
		return err
	}

	// Optional: Print the generated JSON data for debugging.
	fmt.Println(string(jsonData))

	// Send the Root structure (containing the product) to the API.
	resp, err := postProductToAPI(root)
	if err != nil {
		report_helpers.FailedStep()
		report_helpers.PrettyLogStep(stepName, "Failed", fmt.Sprintf("Failed to send API request: %v", err))
		return err
	}
	defer resp.Body.Close()

	// Check if the API returned the correct status code.
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
		report_helpers.FailedStep()
		report_helpers.PrettyLogStep(stepName, "Failed", fmt.Sprintf("Expected status 201 or 200, got %d", resp.StatusCode))
		return fmt.Errorf("expected status 201 or 200, got %d", resp.StatusCode)
	}

	// Store the created product's ID.
	createdProductID = product.ProductCode

	// Log success.
	report_helpers.PassedStep()
	report_helpers.PrettyLogStep(stepName, "Passed", fmt.Sprintf("Product created successfully with ID %s", createdProductID))
	return nil
}

// Step 3: Validate that the product was created successfully with the expected description.
func theProductShouldBeCreatedSuccessfullyWithDescription(expectedDescription string) error {
	stepName := "Validate product creation in the database"
	report_helpers.PrettyLogStep(stepName, "Started", fmt.Sprintf("Expected Description: %s", expectedDescription))

	// Define the query to check the product description in the database.
	query := "SELECT shortDescription FROM product WHERE productid = $1"
	dbValue, err := db_helpers.QueryDatabase(query, createdProductID)
	if err != nil {
		report_helpers.FailedStep()
		report_helpers.PrettyLogStep(stepName, "Failed", fmt.Sprintf("Error: %v", err))
		return err
	}

	// Check if the product description matches the expected value.
	if dbValue != expectedDescription {
		report_helpers.FailedStep()
		report_helpers.PrettyLogStep(stepName, "Failed", fmt.Sprintf("Expected %s, got %s", expectedDescription, dbValue))
		return fmt.Errorf("expected product description %s, got %s", expectedDescription, dbValue)
	}

	// Log success.
	report_helpers.PassedStep()
	report_helpers.PrettyLogStep(stepName, "Passed", fmt.Sprintf("Product validation successful. Description: %s", dbValue))
	return nil
}

// InitializeProductSteps registers the step definitions for the scenario.
func InitializeProductSteps(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(sc *godog.Scenario) {
		report_helpers.PrettyLogScenario(sc.Name, "Started")
		data_helpers.TotalScenarios++
	})

	ctx.AfterScenario(func(sc *godog.Scenario, err error) {
		if err != nil {
			report_helpers.FailedScenario()
			report_helpers.PrettyLogScenario(sc.Name, "Failed")
		} else {
			report_helpers.PassedScenario()
			report_helpers.PrettyLogScenario(sc.Name, "Passed")
		}
	})

	ctx.Step(`^a new testcase with ID "([^"]*)"$`, aNewTestcaseWithID)
	ctx.Step(`^a product with the description "([^"]*)" is created$`, aProductWithTheDescriptionIsCreated)
	ctx.Step(`^the product should be created successfully with description "([^"]*)"$`, theProductShouldBeCreatedSuccessfullyWithDescription)
}

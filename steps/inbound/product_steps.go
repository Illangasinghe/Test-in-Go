package inbound

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-in-go/utils/data_helpers"
	"test-in-go/utils/protocol_helpers"
	"test-in-go/utils/report_helpers"

	"github.com/cucumber/godog"
)

var createdProductID string

// Helper function to handle API requests, now taking the Root structure that includes a products array.
func postProductToAPI(root data_helpers.Root) (*http.Response, error) {
	return protocol_helpers.PostRequest("/product", root)
}

// Step 1: Define a new test case with a dynamic test code based on the test case ID.
func aNewTestcaseWithID(testcaseID string) error {
	return report_helpers.RunAllureStep("Define a new test case with dynamic TestCode", func() error {
		return data_helpers.GenerateTestCode(testcaseID)
	})
}

// Step 2: Create the product with the given description using the previously generated TestCode, wrapped inside a Root structure.
func aProductWithTheDescriptionIsCreated(description string) error {
	return report_helpers.RunAllureStep("Create product with dynamic TestCode", func() error {
		// Create a single product using the ProductBuilder
		product := data_helpers.NewProductBuilder().
			WithTestCode().                      // Use the global TestCode
			WithShortDescription(description).   // Set the description dynamically
			WithSKU(data_helpers.GenerateSKU()). // Assuming SKU generation is dynamic
			Build()

		// Wrap the product in the Root structure (products array)
		root := data_helpers.Root{
			Products: []data_helpers.Product{product},
		}

		// Marshal the Root structure to JSON
		jsonData, err := json.MarshalIndent(root, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to marshal product: %v", err)
		}

		// Optional: Print the generated JSON data for debugging
		fmt.Println(string(jsonData))

		// Send the Root structure containing the product to the API
		resp, err := postProductToAPI(root)
		if err != nil {
			return fmt.Errorf("failed to send API request: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			return fmt.Errorf("expected status 201 or 200, got %d for %s", resp.StatusCode, resp.Request.RequestURI)
		}

		// Save the product code (assuming it's unique) to check later
		createdProductID = product.ProductCode
		return nil
	})
}

// Step 3: Validate that the product was created successfully with the expected description.
func theProductShouldBeCreatedSuccessfullyWithDescription(expectedDescription string) error {
	return report_helpers.RunAllureStep("Validate product creation in the database", func() error {
		query := "SELECT shortDescription FROM product WHERE productid = $1"
		dbValue, err := report_helpers.QueryDatabase(query, createdProductID)
		if err != nil {
			return fmt.Errorf("could not find product with ID %s: %v", createdProductID, err)
		}
		if dbValue != expectedDescription {
			return fmt.Errorf("expected product description %s, got %s", expectedDescription, dbValue)
		}
		return nil
	})
}

// InitializeProductSteps registers the step definitions for the scenario
func InitializeProductSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^a new testcase with ID "([^"]*)"$`, aNewTestcaseWithID)
	ctx.Step(`^a product with the description "([^"]*)" is created$`, aProductWithTheDescriptionIsCreated)
	ctx.Step(`^the product should be created successfully with description "([^"]*)"$`, theProductShouldBeCreatedSuccessfullyWithDescription)
}

package inbound

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test-in-go/db"
	"test-in-go/utils/data_helpers"
	"test-in-go/utils/protocol_helpers"

	"github.com/cucumber/godog"
	"github.com/ozontech/allure-go/pkg/allure"
)

var createdProductID string

func iCreateANewTestcaseWithTestNumbers(categoryNumber, methodNumber, testcaseNumber int) error {
	// Begin the Allure step
	step := allure.NewSimpleStep("Create new test case with dynamic test code").Begin()

	// Generate a product with a dynamically calculated TestCode using the provided numbers
	product := data_helpers.NewProductBuilder().
		WithCalculatedTestCode(categoryNumber, methodNumber, testcaseNumber).
		WithSKU(data_helpers.GenerateSKU()).
		Build()

	// Convert the product to JSON format
	jsonData, err := json.MarshalIndent(product, "", "  ")
	if err != nil {
		step.Failed().Finish()
		return fmt.Errorf("failed to marshal product: %v", err)
	}

	// Print the generated product data (can be removed or logged as needed)
	fmt.Println(string(jsonData))

	// Call the helper function to send the product data to the API
	resp, err := protocol_helpers.PostRequest("/product", product)
	if err != nil {
		step.Failed().Finish() // Mark the step as failed if API call fails
		return fmt.Errorf("failed to send API request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response from the API
	if resp.StatusCode != http.StatusCreated {
		step.Failed().Finish()
		return fmt.Errorf("expected status 201, got %d for "+resp.Request.RequestURI, resp.StatusCode)
	}

	// Mark the step as passed and finalize it
	step.Passed().Finish()

	createdProductID = product.ProductCode
	return nil
}

func theProductShouldBeCreatedSuccessfully(shortDescription string) error {
	// You can add any validation logic here, such as checking the database to see if the product was created.
	// For now, we will assume that the previous step (the POST request) checks for success.
	// Query the database to verify product exists
	var dbValue string
	query := "SELECT shortDescription FROM product WHERE productid = $1"
	err := db.DB.QueryRow(query, createdProductID).Scan(&dbValue)
	if err != nil {
		return fmt.Errorf("could not find product with ID %s: %v", createdProductID, err)
	}
	if dbValue != shortDescription {
		return fmt.Errorf("expected product name %s, got %s", shortDescription, dbValue)
	}
	return nil
}

func InitializeProductSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a new testcase with the following test numbers (\d+), (\d+), (\d+)$`, iCreateANewTestcaseWithTestNumbers)
	ctx.Step(`^the product should be created successfully$`, theProductShouldBeCreatedSuccessfully)
}

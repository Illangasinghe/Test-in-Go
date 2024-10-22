package inbound

import (
	"fmt"
	"net/http"

	"test-in-go/utils/protocol_helpers"

	// "test-in-go/data_helpers"  // Import the data_helpers

	// "test-in-go/utils/data_helpers"

	"github.com/cucumber/godog"
	"github.com/ozontech/allure-go/pkg/allure"
)

var createdProductID string

// Product structure for the API payload
type Product struct {
	ProductID string `json:"product_id"`
	Name      string `json:"name"`
}

func iCreateAProductWithTheFollowingDetails(productID, name string) error {
	// Begin the step for product creation
	step := allure.NewSimpleStep(fmt.Sprintf("Create Product with ID: %s", productID)).Begin()

	// Create the product payload
	product := Product{
		ProductID: productID,
		Name:      name,
	}
	// Call the helper function with authentication
	resp, err := protocol_helpers.PostRequest("/products", product)
	if err != nil {
		step.Failed().Finish() // Mark the step as failed if API call fails
		return fmt.Errorf("failed to send API request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		step.Failed().Finish()
		return fmt.Errorf("expected status 201, got %d", resp.StatusCode)
	}

	// If everything is successful, mark the step as passed
	step.Passed().Finish()

	createdProductID = productID
	return nil
}

func InitializeProductSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a product with the following details "([^"]*)" and "([^"]*)"$`, iCreateAProductWithTheFollowingDetails)
}

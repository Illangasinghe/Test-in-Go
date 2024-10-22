package inbound

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"test-in-go/config"

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

	// Convert the product to JSON
	productJSON, err := json.Marshal(product)
	if err != nil {
		step.Failed().Finish() // Mark the step as failed if marshaling fails
		return fmt.Errorf("failed to marshal product: %v", err)
	}

	// Send a POST request to the product API
	apiURL := config.GetEnv("API_URL") + "/products"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(productJSON))
	if err != nil {
		step.Failed().Finish() // Mark the step as failed if request creation fails
		return fmt.Errorf("failed to create API request: %v", err)
	}

	// Set content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Execute the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		step.Failed().Finish() // Mark the step as failed if API call fails
		return fmt.Errorf("failed to send API request: %v", err)
	}
	defer resp.Body.Close()

	// Validate the response status code
	if resp.StatusCode != http.StatusCreated {
		step.Failed().Finish() // Mark the step as failed if response is not 201
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

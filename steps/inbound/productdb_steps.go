package inbound

import (
	"fmt"
	"test-in-go/db"

	"github.com/cucumber/godog"
)

var createdProductIDDC string

func iCreateAProductOnDbWithTheFollowingDetails(productID, name string) error {
	// Insert product into the database
	query := "INSERT INTO product (productid, productlevel, parentid, parentlevel, longdescription) VALUES ($1, 'PRD', 'LOREWI', 'PH1', $2)"
	_, err := db.DB.Exec(query, productID, name)
	if err != nil {
		return fmt.Errorf("failed to create product: %v", err)
	}
	createdProductID = productID
	return nil
}

func theProductShouldBeInTheDatabaseWithName(name string) error {
	// Query the database to verify product exists
	var dbName string
	query := "SELECT longdescription FROM product WHERE productid = $1"
	err := db.DB.QueryRow(query, createdProductID).Scan(&dbName)
	if err != nil {
		return fmt.Errorf("could not find product with ID %s: %v", createdProductID, err)
	}
	if dbName != name {
		return fmt.Errorf("expected product name %s, got %s", name, dbName)
	}
	return nil
}

func InitializeProductDBSteps(ctx *godog.ScenarioContext) {
	ctx.Step(`^I create a product with the following details "([^"]*)" and "([^"]*)"$`, iCreateAProductWithTheFollowingDetails)
	ctx.Step(`^the product should be in the database with name "([^"]*)"$`, theProductShouldBeInTheDatabaseWithName)
}

Feature: Product Creation
  Scenario: Create a new product
    Given I create a product with the following details "sample-123" and "Test Description"
    Then the product should be in the database with name "Test Description"
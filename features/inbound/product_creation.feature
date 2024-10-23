Feature: Product Creation

  Scenario: Create a product using dynamic TestCode
    Given a new testcase with ID "110-010-001"
    When a product with the description "Test Product" is created
    Then the product should be created successfully with description "Test Product"

Feature: Product Creation

  Scenario: Create a product using dynamic TestCode
    Given I create a new testcase with the following test numbers 110, 10, 1
    Then the product should be created successfully

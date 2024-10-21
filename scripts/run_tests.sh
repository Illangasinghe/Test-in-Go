#!/bin/bash

echo "Preparing to run tests..."

# Set environment variables
export ENVIRONMENT=${ENVIRONMENT:-dev}
export DB_URL=${DB_URL:-postgres://user:pass@localhost/test_db?sslmode=disable}

# Run the tests
echo "Running tests in $ENVIRONMENT environment"
go run main.go

# Check the result of the tests and act accordingly
if [ $? -ne 0 ]; then
  echo "Tests failed. Please check the logs."
  exit 1
else
  echo "Tests passed successfully!"
fi

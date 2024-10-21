#!/bin/bash

echo "Serving Allure reports..."

# Serve the Allure report
allure serve ./reports/allure-results

if [ $? -ne 0 ]; then
  echo "Failed to serve Allure reports."
  exit 1
else
  echo "Allure report is now accessible."
fi

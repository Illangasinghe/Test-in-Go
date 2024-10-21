#!/bin/bash

echo "Starting the web UI for test execution..."

# You can replace this with the actual command to run your custom Web UI, if applicable
# Example: Running a node server or Python Flask app
npm start  # Or any command to serve your web UI

if [ $? -ne 0 ]; then
  echo "Failed to start the web UI."
  exit 1
else
  echo "Web UI is running. Access it via http://localhost:3000"
fi

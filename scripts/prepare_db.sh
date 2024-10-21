#!/bin/bash

echo "Preparing the database..."

# Set up the database using the init script
psql -h localhost -U user -d test_db -f ./fixtures/init_db.sql

# Check for errors
if [ $? -ne 0 ]; then
  echo "Database setup failed."
  exit 1
else
  echo "Database setup completed successfully."
fi

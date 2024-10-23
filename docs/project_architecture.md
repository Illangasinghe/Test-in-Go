# Project Architecture

## Overview

The **Test in Go** framework is a modular and scalable BDD test automation framework built in Go. It supports multiple protocols (SOAP, REST, Kafka), database interactions, and Docker integration for isolated environments. The architecture is designed for flexibility, maintainability, and ease of extension, with clear separation between test definitions, logic, and utilities.

---

## Key Architectural Principles

1. **Separation of Concerns**: The framework clearly separates test definitions (feature files), logic (step definitions), and utilities (helpers), ensuring a clean, maintainable codebase.
   
2. **Modularity and Extensibility**: New test scenarios, protocols, or database support can be easily added without affecting the core functionality.

3. **Scalability**: The framework scales well across various environments, supporting both local and Docker-based execution, and integrates seamlessly with CI/CD pipelines.

4. **Protocol-Agnostic**: The framework supports multiple protocols (SOAP, REST, Kafka) for integration tests, making it versatile across different systems.

5. **BDD-Driven**: Tests are written in Gherkin syntax using feature files, making it easy for both technical and non-technical stakeholders to collaborate through BDD (Behavior-Driven Development).

---

## High-Level Architecture

### Test Definitions Layer

- **Features**: Gherkin feature files (`.feature`) define test scenarios using Given-When-Then syntax. Files are categorized into inbound, outbound, and integration scenarios.
- **Steps**: Go step definition files implement the logic for each Gherkin step and are linked to the feature files, handling API calls, database validation, and other test operations.

### Utility Layer

Reusable helper functions for tasks such as API requests, Kafka messaging, and database interactions. The utility functions are modular and extendable.

- **API Helpers**: Manage SOAP, REST, and Kafka interactions.
- **Database Helpers**: Handle querying, data validation, and cleanups.
- **Validation Helpers**: Perform schema and field-level validation for JSON and XML.
- **Logging Helpers**: Capture detailed logs for requests, responses, and errors during test execution.

### Configuration Layer

Configuration files and environment variables make the framework adaptable to different environments.

- **config.json**: Stores database connections, API URLs, and other settings.
- **env.go**: Handles environment variable management for Docker and local execution.

### Execution Layer

- **Docker Integration**: Tests can be run in isolated Docker environments using the provided `Dockerfile` and `docker-compose.yml`.
- **Standalone Execution**: Tests can also be run locally or integrated into CI/CD pipelines via `main.go`.
- **Automation Scripts**: Scripts for running tests, preparing databases, and generating reports in `scripts/`.

### Reporting Layer

- **Allure Reports**: The framework integrates with Allure for detailed visual test reports. Reports are stored in the `reports/` directory and can be served locally or via CI pipelines.

---

## Component Interaction

### BDD Flow

1. **Feature File**: Defines the test scenario using Gherkin.
2. **Step Definitions**: Links Gherkin steps to Go functions.
3. **Utility Functions**: Handles API calls, validations, and database interactions.
4. **Execution**: Tests are run in either local or Docker environments, with results captured in reports.

### Docker and CI/CD Integration

- **Docker**: The framework supports isolated execution using Docker containers.
- **CI/CD**: Integration with CI tools (Jenkins, GitLab CI, GitHub Actions) allows for automated test execution in pipelines, with results logged and visualized in Allure reports.

### Database Interaction

Database-related tests interact with PostgreSQL using helper functions to query, verify, and clean up the database during and after tests.

---

## Project Structure

test-in-go/
│
├── Dockerfile                           # Dockerfile for building the test framework container
├── docker-compose.yml                   # Docker Compose config to manage DB, SUT, and test framework
├── go.mod                               # Go modules file for dependency management
├── go.sum                               # Go module dependency checksums
├── README.md                            # Project documentation, usage examples, and contribution guidelines
├── LICENSE.md                           # License for public usage
├── config/                              # Configuration files
│   ├── config.json                      # Main config file (DB, API settings, etc.)
│   └── env.go                           # Helper for loading environment variables
│
├── features/                            # Gherkin feature files, organized by message type/functionality
│   ├── inbound/                         # Inbound calls (messages from external systems)
│   │   ├── product_creation.feature
│   │   ├── stock_update.feature
│   │   └── order_placement.feature
│   ├── outbound/                        # Outbound calls (messages from core)
│   │   ├── stock_balance_update.feature
│   │   └── order_status_update.feature
│   └── integration_scenarios/           # Complex scenarios (inbound + outbound combined)
│       └── product_order_scenario.feature
│
├── steps/                               # Step definitions for Gherkin feature files
│   ├── inbound/                         # Step definitions for inbound features
│   │   ├── product_steps.go
│   │   ├── stock_steps.go
│   │   └── order_steps.go
│   ├── outbound/                        # Step definitions for outbound features
│   │   ├── stock_balance_steps.go
│   │   └── order_status_steps.go
│   └── integration_scenarios/           # Step definitions for complex integration scenarios
│       └── product_order_scenario_steps.go
│
├── utils/                               # Utility functions
│   ├── protocol_helpers/                # Protocol dealing helpers
│   │   ├── soap_helpers.go              # SOAP-specific message generation, sending, and parsing
│   │   ├── rest_helpers.go              # REST-specific HTTP requests and responses
│   │   ├── kafka_helpers.go             # Kafka producer and consumer helpers
│   ├── api_helpers/                     # API calling functions (HTTP requests)
│   │   ├── inbound_api_helpers.go       # Helper functions for inbound API calls
│   │   └── outbound_api_helpers.go      # Helper functions for outbound API calls
│   ├── data_helpers/                    # Dynamic generic test data for tests
│   │   ├── dynamic_values.go            # Dynamic test data
│   │   ├── default_values.go            # Default values
│   │   ├── inbound/                     # Sample test data for inbound messages
│   │   |   ├── sample_product.go        # Static product test data
│   │   |   └── sample_order.go          # Static order test data
│   │   └── outbound/                    # Sample test data for outbound messages
│   │       ├── sample_stock_balance.go  # Static stock balance test data
│   │       └── sample_order_status.go   # Static order test data
│   ├── response_helpers/                # Helpers for storing and retrieving responses
│   │   ├── inbound_response_helpers.go
│   │   └── outbound_response_helpers.go
│   ├── validation_helpers/              # Helpers for validation (response checks, schemas)
│   │   ├── json_schema_validation.go
│   │   ├── json_field_validation.go
│   │   ├── xml_schema_validation.go
│   │   └── xml_field_validation.go
│   ├── message_helpers/                 # Helpers for building dynamic messages
│   │   ├── json_message_builder.go
│   │   └── xml_message_builder.go
│   ├── logging_helpers/                 # Logging system for API requests, responses, and errors
│   │   └── logrus_setup.go
│   └── db_helpers/                      # Database interaction helpers
│       ├── db_helper.go                 # Generic DB helper functions (supports multiple databases)
│       └── postgres_db_helper.go        # PostgreSQL-specific helper functions
│
├── templates/                           # Directory for JSON/XML templates
│   ├── inbound/                         # Templates for inbound messages
│   │   ├── call/                        # Original inbound message inputs to SUT (from host)
│   │   │   ├── product_update_template.json
│   │   │   └── order_placement_template.json
│   │   └── callout/                     # Translated outbound message outputs of SUT (to core)
│   │       ├── product_update_template.json
│   │       └── order_placement_template.json
│   └── outbound/                        # Templates for outbound messages
│       ├── call/                        # Original outbound message inputs to SUT (from core)
│       │   ├── stock_balance_update_template.json
│       │   └── order_status_update_template.json
│       └── callout/                     # Translated outbound message outputs of SUT (to host)
│           ├── stock_balance_update_template.json
│           └── order_status_update_template.json
│
├── db/                                  # Database management (database-agnostic)
│   ├── postgres.go                      # Database interaction module (Primarily PostgreSQL) as an example DB
│   └── db_interface.go                  # Interface for abstracting database methods
│
├── docs/                                # Documentation directory (public resource focus - Phase 6)
│   ├── usage.md                         # How to use the framework
│   ├── contribution.md                  # Contribution guide for adding new features, databases, or tests
│   └── database_setup.md                # Instructions for setting up supported databases (e.g., PostgreSQL, MySQL)
│
├── fixtures/                            # Predefined test fixtures for databases
│   ├── init_db.sql                      # SQL script for initializing db with sample data scripts (if required)
│   └── clean_db.sql                     # SQL script for cleaning database after test execution
│
├── reports/                             # Directory for storing test execution results and reports
│   └── allure-results                   # Allure-specific test result reports
│
├── docker/                              # Docker setup for running isolated tests (Phase 2)
│   ├── Dockerfile                       # The project dockerfile
│   └── docker-compose.yml               # Docker Compose configuration for linked services (e.g., API and DB)
│
├── scripts/                             # Automation scripts
│   ├── run_tests.sh                     # Script to run the tests (ideal for CI pipelines)
│   ├── prepare_db.sh                    # Script to set up databases (e.g., run init scripts for data pre-requisites)
│   ├── run_allure_report.sh             # Script to serve allure reports locally for CLI-based execution
│   └── run_web_ui.sh                    # Script to serve test execution dashboard for web-based execution (Phase 4)
│
├── logs/                                # Logging for better debug (Phase 3)
│   ├── execution.log                    # Stores logs of each test run (stdout, errors)
│   └── results.log                      # Logs test against timestamp and results on each step on each test execution
│
├── ci_config/                           # Add configurations for popular CI/CD tools (Phase 5)
│   ├── jenkinsfile                      # Jenkins pipeline config for running the test framework
│   ├── gitlab-ci.yml                    # GitLab CI configuration file
│   └── github-actions.yml               # GitHub Actions config
│
└── main.go                              # Main entry point for executing the tests
```

---

## Extensibility

- **Add New Protocols**: Easily add support for new protocols (e.g., gRPC) in the `utils/` directory.
- **Support New Databases**: Extend database support by adding new database helper modules in `db_helpers/`.
- **Expand Test Scenarios**: Add new feature files and step definitions without affecting existing ones.

---

For more details on setup and usage, see the [Getting Started Guide](getting_started.md) or explore other documents in the `docs/` folder.
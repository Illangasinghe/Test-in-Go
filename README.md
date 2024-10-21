# Test in Go

Welcome to **Test in Go** - a Go-based test automation framework. This project is designed to help developers perform API testing, database validations, and Docker integration in a clean, modular, and scalable manner. The framework is built with Go and leverages Gherkin-based Godog for behavior-driven testing, with step definitions linked to BDD steps for improved readability and collaboration.

## Table of Contents

- [Project Structure](#project-structure)
- [Getting Started](#getting-started)
- [Installation](#installation)
- [Running Tests](#running-tests)
- [Configuration](#configuration)
- [Directory Overview](#directory-overview)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

The framework has been structured to make it easy to manage, extend, and scale. Here is a brief overview of the key directories and files:

```
test-in-go/
│
├── Dockerfile                           # Dockerfile for building the test framework container
├── docker-compose.yml                   # Docker Compose configuration to manage services (DB, SUT, test framework)
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
├── data/                                # Sample test data directory
│   ├── dynamic/                         # Dynamic generic test data for tests
│   |   ├── dynamic_variables.go         # Dynamic test data
│   |   └── timebound_variables.go       # Tim-bound variouse test data
│   ├── static/                          # Static generic test data for tests
│   |   ├── default_values.go            # Default values
│   |   └── static_variables.go          # Static test data
│   ├── inbound/                         # Sample test data for inbound messages
│   |   ├── sample_product.go            # Static product test data
│   |   └── sample_order.go              # Static order test data
│   └── outbound/                        # Sample test data for outbound messages
│       ├── sample_stock_balance.go      # Static stock balance test data
│       └── sample_order_status.go       # Static order test data
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
│   ├── prepare_db.sh                    # Script to set up databases (e.g., run init scripts)
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

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or above)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Allure](https://docs.qameta.io/allure/#_get_started) for generating and viewing reports

### Installation

1. **Clone the Repository**

   ```bash
   git clone https://github.com/Illangasinghe/Test-in-Go.git
   cd Test-in-Go
   ```

2. **Install Dependencies**

   Use `go mod` to download the necessary Go packages:

   ```bash
   go mod tidy
   ```

3. **Build Docker Containers**

   Build the Docker containers for your services using Docker Compose:

   ```bash
   docker-compose up --build
   ```

## Running Tests

To run all tests:

```bash
# Run Godog tests
godog ./tests/features
```

Or, if you prefer using Docker to run the tests in an isolated environment:

```bash
docker-compose run test-framework
```

You can also use the included script for running the tests conveniently:

```bash
./scripts/run_tests.sh
```

## Configuration

All configuration parameters are located in the `config/config.json` file. This includes database connection details, API URLs, and other key settings for the framework. The `env.go` file manages environment variables to make the framework flexible in Docker environments.

## Directory Overview

- **config/**: Contains configuration files used for the framework, including general settings (`config.json`).
- **templates/**: Predefined JSON/XML templates for generating dynamic test data, categorized by inbound and outbound messages.
- **data/**: Static test data files that are used for testing without modification.
- **utils/**: Helper functions, such as JSON/XML parsing, HTTP client utilities, response handling, and database interactions.
  - **api_helpers/**: Functions for making API calls.
  - **response_helpers/**: Functions to manage responses.
  - **validation_helpers/**: Response and schema validation helpers.
  - **message_helpers/**: Building dynamic inbound and outbound messages.
  - **db_helpers/**: Database interactions and validation.
- **db/**: Handles database connections and validation functions, and supports multiple databases (e.g., PostgreSQL, MySQL, MongoDB).
- **tests/**: Organized by entity and type of call (e.g., `inbound/`, `outbound/`, `integration_scenarios/`), and contains feature files for behavior-driven tests and corresponding step definitions.
  - **features/**: Gherkin feature files that define the test scenarios.
  - **steps/**: Step definition files that link the Gherkin steps to Go code.
- **docs/**: Documentation for using, contributing, and setting up the framework, making it easier for public users.
- **fixtures/**: SQL scripts for initializing and cleaning up the database before and after test execution.
- **reports/**: Stores test reports for tracking results and analysis, useful in CI/CD environments (Allure results are stored here).
- **docker/**: Docker setup for isolated test environments.
- **scripts/**: Contains scripts to help automate tasks, such as `run_tests.sh` for executing tests or `prepare_db.sh` for preparing databases.
- **main.go**: Entry point to run the test framework as a standalone executable.

## Contributing

Contributions are welcome! If you have any ideas, improvements, or bugs to report, feel free to open an issue or create a pull request. Please follow the standard guidelines for contributing:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/YourFeature`).
3. Commit your changes (`git commit -m 'Add some feature'`).
4. Push to the branch (`git push origin feature/YourFeature`).
5. Open a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

---

Thank you for using **Test in Go**! We hope this helps make your testing efficient, organized, and enjoyable.


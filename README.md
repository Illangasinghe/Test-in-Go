# Test in Go

Welcome to **Test in Go** - a Go-based test automation framework. This project is designed to help developers perform API testing, database validations, and Docker integration in a clean, modular, and scalable manner. The framework is built with Go and leverages Ginkgo and Gomega for structured and expressive test cases.

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
├── Dockerfile                     # Dockerfile for the test framework
├── docker-compose.yml              # Docker Compose for managing services (database, SUT, test framework)
├── go.mod                          # Go modules for dependency management
├── go.sum                          # Go checksum
├── README.md                       # Documentation, usage examples, contributing guidelines
├── LICENSE.md                      # License for public usage
├── config/                         # Configuration files
│   ├── config.yaml                 # Main configuration file for database and API settings
│   └── env.go                      # Helper to load environment variables
│
├── templates/                      # Directory for JSON/XML templates
│   ├── product_template.json       # Sample product JSON template
│   └── order_template.xml          # Sample order XML template
│
├── data/                           # Sample test data directory
│   ├── sample_product.json         # Static product test data
│   └── sample_order.xml            # Static order test data
│
├── utils/                          # Utility functions
│   ├── json_helper.go              # JSON parsing, validation, and manipulation
│   ├── xml_helper.go               # XML parsing, validation, and manipulation
│   ├── template_loader.go          # Functions to load and apply dynamic templates
│   ├── db_helper.go                # Generic DB helper functions (supports multiple databases)
│   ├── api_client.go               # API request handler
│   └── config_helper.go            # Configuration loader from YAML/environment
│
├── db/                             # Database management (database-agnostic)
│   ├── postgres.go                 # PostgreSQL database interaction module
│   ├── mysql.go                    # MySQL database interaction module (as an example)
│   ├── mongo.go                    # MongoDB database interaction module (as an example)
│   └── db_interface.go             # Interface for abstracting database methods
│
├── tests/                          # Ginkgo test cases organized by entity or function
│   ├── product/
│   │   ├── product_create_test.go  # API test for creating product
│   │   ├── product_update_test.go  # API test for updating product
│   │   └── product_db_test.go      # DB test for validating product details
│   ├── order/
│   │   └── order_create_test.go    # API and DB validation tests for order creation
│   └── common/
│       └── api_test_helpers.go     # Common helper functions for API tests
│
├── docs/                           # Documentation directory (public resource focus)
│   ├── usage.md                    # How to use the framework
│   ├── contribution.md             # Contribution guide for adding new features, databases, or tests
│   └── database_setup.md           # Instructions for setting up supported databases (e.g., PostgreSQL, MySQL)
│
├── fixtures/                       # Predefined test fixtures for databases
│   ├── init_db.sql                 # SQL script for initializing PostgreSQL with sample data
│   ├── init_mysql.sql              # MySQL initialization script (if using MySQL)
│   └── clean_db.sql                # SQL script for cleaning database after test execution
│
├── reports/                        # Directory for storing test execution results and reports
│   └── junit_report.xml            # Example test report output for CI tools (e.g., Jenkins, GitLab)
│
├── scripts/                        # Automation scripts
│   ├── run_tests.sh                # Script to run the tests (ideal for CI pipelines)
│   └── prepare_db.sh               # Script to set up databases (e.g., run init scripts)
│
└── main.go                         # Main entry point for executing the tests
```

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or above)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

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
go test ./...
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

All configuration parameters are located in the `config/config.yaml` file. This includes database connection details, API URLs, and other key settings for the framework. The `env.go` file manages environment variables to make the framework flexible in Docker environments.

## Directory Overview

- **config/**: Contains configuration files used for the framework, including general settings (`config.yaml`).
- **templates/**: Predefined JSON/XML templates for generating dynamic test data.
- **data/**: Static test data files that are used for testing without modification.
- **utils/**: Helper functions, such as JSON/XML parsing, HTTP client utilities, and template loaders.
- **db/**: Handles database connections and validation functions, and supports multiple databases (e.g., PostgreSQL, MySQL, MongoDB).
- **tests/**: Organized by entity (e.g., `product/`, `order/`), and contains API tests, database validation tests, etc.
- **docs/**: Documentation for using, contributing, and setting up the framework, making it easier for public users.
- **fixtures/**: SQL scripts for initializing and cleaning up the database before and after test execution.
- **reports/**: Stores test reports for tracking results and analysis, useful in CI/CD environments.
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


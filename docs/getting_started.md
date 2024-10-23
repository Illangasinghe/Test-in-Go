# Getting Started with Test in Go

Welcome to **Test in Go**, a Go-based test automation framework that simplifies API testing, database validations, and Docker integration. This guide will help you set up the framework, understand its key components, and run your first tests.

## Prerequisites

Before you begin, ensure that you have the following installed on your machine:

- **[Go](https://golang.org/doc/install)** (version 1.18 or above)
- **[Docker](https://docs.docker.com/get-docker/)** for containerized test environments
- **[Docker Compose](https://docs.docker.com/compose/install/)** for managing multi-container Docker applications
- **[Allure](https://docs.qameta.io/allure/#_get_started)** for generating and viewing test reports

## Installation

### Step 1: Clone the Repository

Start by cloning the **Test in Go** repository to your local machine:

```bash
git clone https://github.com/Illangasinghe/Test-in-Go.git
cd Test-in-Go
```

### Step 2: Install Dependencies

Use Go's dependency management tool, `go mod`, to install all necessary packages for the framework:

```bash
go mod tidy
```

### Step 3: Build Docker Containers

If you're planning to use Docker for isolated testing environments, build the Docker containers for the services with Docker Compose:

```bash
docker-compose up --build
```

This step sets up the required services, such as databases or application servers, required for the tests.

## Configuration

All framework configurations are managed in the `config/config.json` file. This includes:

- **Database connection settings** (e.g., PostgreSQL)
- **API URLs** for SOAP, REST, or Kafka endpoints
- **Environment variables** handled by `env.go` to adapt the framework to different environments (especially useful in Docker-based setups)

You may need to adjust these configurations to suit your environment or test requirements.

## Running Tests

Once the setup is complete, you're ready to run your tests. Tests are written in **Gherkin** format and use **Godog** for executing BDD-style scenarios.

### Option 1: Run Tests Locally

You can run all the tests directly using `godog`:

```bash
godog ./features/inbound/product
```

This command runs the tests located in the `features/inbound/product` directory.

### Option 2: Run Tests in Docker

Alternatively, you can run the tests inside a Docker container:

```bash
docker-compose run test-framework
```

This approach provides an isolated environment to ensure consistency across different systems.

### Option 3: Use Predefined Script

A convenient script, `run_tests.sh`, is available for running tests with predefined configurations:

```bash
./scripts/run_tests.sh
```

This script streamlines the execution process, automatically preparing the environment and running the test suite.

## Viewing Test Reports

The **Test in Go** framework uses **Allure** for generating test reports. Once tests are completed, you can generate and view the reports using the following script:

```bash
./scripts/run_allure_report.sh
```

This command will serve the Allure report on a local server, allowing you to interactively browse the test results.

## Directory Structure Overview

Here’s a brief explanation of the key directories relevant to getting started with **Test in Go**:

- **config/**: Holds the configuration files (`config.json`) to manage database connections, API settings, and environment variables.
- **features/**: Contains Gherkin feature files for behavior-driven tests.
- **steps/**: Step definition files in Go, linking Gherkin steps to Go code for executing the test logic.
- **utils/**: Helper functions, such as protocol-specific utilities, validation helpers, and data generators.
- **docker/**: Contains Docker configuration files, including `Dockerfile` and `docker-compose.yml`, for setting up the test environment.
- **scripts/**: Automation scripts to run tests, generate reports, and manage databases.
- **reports/**: Stores test execution results and Allure reports for easy tracking and visualization.

## Next Steps

Once you have your environment set up and you've run your first tests, you can explore additional functionality:

- **Writing Tests**: Learn how to write effective BDD tests by following the Gherkin test conventions in `docs/naming_conventions.md`.
- **Database Setup**: Configure supported databases (e.g., PostgreSQL) by referring to `[docs/database_setup.md](./database_setup.md)`.
- **Contributing**: If you wish to contribute to this framework, check the guidelines in `docs/contribution.md`.

## Troubleshooting

- **Docker Issues**: If you encounter problems with Docker, ensure that Docker and Docker Compose are installed correctly and that your system meets the requirements.
- **Dependency Errors**: If you face issues during Go module installation, run `go mod tidy` to clean and update the dependencies.
---

## Additional Documentation and Resources

To explore further details and specific guides, refer to the following documentation:

- **[README.md](../README.md)**: The main documentation file in the root directory, providing an overview of the framework, installation steps, and key features.
- **[Contribution Guide](./contribution.md)**: Guidelines on how to contribute to the project, including adding new features, bug fixes, and best practices.
- **[Database Setup](./database_setup.md)**: Instructions for configuring and setting up supported databases, such as PostgreSQL and MySQL, for testing purposes.
- **[Naming Conventions](./naming_conventions.md)**: Detailed guidelines for writing and organizing test scenarios and feature files following the Gherkin format.
- **[Priorities](./priorities.md)**: Insights into the development priorities and roadmap for future improvements to the framework.
- **[Change Log](./change_log.md)**: (Coming soon) A record of changes, updates, and new features introduced in each version of the framework.

For further details and advanced usage, refer to the full documentation in the `/docs` directory.

---

Thank you for using **Test in Go**! We're excited to see how you apply this framework in your testing workflows.
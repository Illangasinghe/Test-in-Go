# Test-in-Go Framework Usage

## Getting Started

This framework is designed for behavior-driven development (BDD) using **Godog** in Go. It supports testing across different protocols (SOAP, REST, Kafka) and includes logging, reporting, and database interactions.

### Prerequisites

- **Docker**: Ensure Docker is installed to run the tests in an isolated environment.
- **Go**: Version 1.17 or higher is required to build and run the test framework.

### 1. Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/your-username/Test-in-Go.git
cd Test-in-Go
```

### 2. Build the Docker Environment

Use Docker Compose to build and run the services (test framework, SUT, and database):

```bash
docker-compose up --build
```

### 3. Prepare the Database

Run the `prepare_db.sh` script to set up the PostgreSQL database with the necessary schema and test data:

```bash
./scripts/prepare_db.sh
```

### 4. Running Tests

Run the tests using the following script:

```bash
./scripts/run_tests.sh
```

This will run the tests using **Godog** and generate a **Pretty Report**.

### 5. Pretty Report

Test results are logged in a human-readable format in the `./reports/pretty-report.txt` file. This report includes step-by-step status logs (Passed, Failed, Skipped) and a summary of the scenarios and steps executed.

To manually inspect the report, you can check the file generated after the test run:

```bash
cat ./reports/pretty-report.txt
```

### 6. Optional: Web UI Mode

You can also run a web UI to view, execute, and manage your feature files and reports. The web UI can be launched with:

```bash
go run main.go --web-ui
```

The web interface provides features such as viewing the feature files, executing tests, and displaying results interactively.

If you prefer to run tests without the web UI (e.g., for CI/CD purposes), simply run:

```bash
go run main.go --run-tests
```

---

## Environment Variables

You can configure the environment variables as needed:

- **ENVIRONMENT**: The environment in which the tests will run (e.g., `dev`, `staging`).
- **DB_URL**: The PostgreSQL connection URL.

These can be set in your terminal or as part of the `docker-compose.yml` file.

---

## Troubleshooting

### 1. Database Connection Issues

If you encounter problems connecting to the database, ensure the following:

- The `DB_URL` is correctly configured in `config/config.json` or as an environment variable.
- PostgreSQL is running and accessible (check Docker container logs).

### 2. Test Failures

Check the logs in `./logs/execution.log` to see detailed error messages for test failures. Review the pretty report at `./reports/pretty-report.txt` for more details.

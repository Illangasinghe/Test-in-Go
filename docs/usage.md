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

This will run the tests using **Godog** and generate an Allure report.

### 5. Viewing Allure Reports

To view the test results in a browser, use the following script:

```bash
./scripts/run_allure_report.sh
```

The Allure report will be accessible at `http://localhost:4040`.

---

## Environment Variables

You can configure the environment variables as needed:

- **ENVIRONMENT**: The environment in which the tests will run (e.g., `dev`, `staging`).
- **DB_URL**: The PostgreSQL connection URL.
- **ALLURE_RESULTS_DIRECTORY**: The directory where Allure stores its test results.

These can be set in your terminal or as part of the `docker-compose.yml` file.

---

## Troubleshooting

### 1. Database Connection Issues

If you encounter problems connecting to the database, ensure the following:

- The `DB_URL` is correctly configured in `config/config.json` or as an environment variable.
- PostgreSQL is running and accessible (check Docker container logs).

### 2. Test Failures

Check the logs in `./logs/execution.log` to see detailed error messages for test failures. Additionally, review the Allure report for more detailed insights.

---
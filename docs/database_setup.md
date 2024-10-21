# Database Setup

The **Test-in-Go** framework uses **PostgreSQL** as the primary database. This guide will walk you through setting up the database, both locally and in Docker.

---

## Local PostgreSQL Setup

If you want to set up PostgreSQL locally (instead of using Docker), follow these steps:

### 1. Install PostgreSQL

Install PostgreSQL on your machine. For Mac users, you can use Homebrew:

```bash
brew install postgresql
```

For Ubuntu users, use `apt`:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

### 2. Start PostgreSQL

Start the PostgreSQL service:

```bash
brew services start postgresql  # For Mac
sudo service postgresql start   # For Ubuntu
```

### 3. Create a Database and User

Create the `test_db` database and a user for the tests:

```bash
psql postgres
CREATE DATABASE test_db;
CREATE USER user WITH ENCRYPTED PASSWORD 'pass';
GRANT ALL PRIVILEGES ON DATABASE test_db TO user;
```

### 4. Update Configuration

Ensure that the `config/config.json` file points to your local PostgreSQL instance:

```json
{
  "db": {
    "url": "postgres://user:pass@localhost:5432/test_db?sslmode=disable"
  }
}
```

---

## Docker-Based PostgreSQL Setup

Alternatively, you can use **Docker Compose** to spin up PostgreSQL alongside the test framework.

### 1. Build and Run Docker Compose

Run the following command to bring up the PostgreSQL service defined in `docker

-compose.yml`:

```bash
docker-compose up db
```

This will create a PostgreSQL instance with the default credentials set in the compose file.

### 2. Access the PostgreSQL Database

To access the PostgreSQL instance running in Docker, use the following command:

```bash
docker exec -it test-in-go-db psql -U user -d test_db
```

### 3. Database Initialization

Run the SQL script to initialize the schema and test data:

```bash
./scripts/prepare_db.sh
```

---

## Troubleshooting

### Connection Issues

If you're unable to connect to the database:

- Verify the **DB_URL** in your `config/config.json` file or environment variables.
- Ensure the PostgreSQL service is running (`docker-compose logs db`).
- Check that the correct database credentials are being used.
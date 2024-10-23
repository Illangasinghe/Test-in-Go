# Test in Go 🚀

Welcome to **Test in Go** – a Go-based test automation framework built for scalability, modularity, and behavior-driven testing (BDD). This framework leverages Gherkin feature files with step definitions written in Go, enabling seamless collaboration between developers and stakeholders while ensuring robust API testing, database validation, and Docker integration.

## Key Features ✨

- 📝 **BDD-Driven**: Use Gherkin syntax for behavior-driven development, making test scenarios readable and maintainable.
- 🔗 **Protocol Support**: Built-in support for SOAP, REST, and Kafka for API and integration testing.
- 🗄️ **Database Validation**: Seamlessly interact with databases (PostgreSQL support) for comprehensive validation.
- 🐳 **Docker Integration**: Run tests in isolated, containerized environments.
- 🔧 **CI/CD Friendly**: Integrates with Jenkins, GitLab CI, GitHub Actions, and more.
- 📊 **Simple Reporting**: Generate detailed visual reports for easy test analysis.

---

## Table of Contents 📚

- [Installation](#installation)
- [Running Tests](#running-tests)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [Documentation](#documentation)
- [License](#license)

---

## Installation ⚙️

To get started with **Test in Go**, follow the [Getting Started Guide](docs/getting_started.md) for a comprehensive walkthrough of setup, installation, and configuration.

## Running Tests 🏃‍♂️

You can run tests locally or in Docker. Detailed instructions for running tests can be found in the [Getting Started Guide](docs/getting_started.md).

- **Local Execution**: Run BDD tests using the `godog` command.
- **Docker Execution**: Use Docker Compose to run tests in an isolated environment.

## Configuration 🛠️

The framework’s configuration, including database settings and API endpoints, is managed through the `config/config.json` file. For more details on configuring the framework, refer to the [Getting Started Guide](docs/getting_started.md#configuration).

## Contributing 🤝

We welcome contributions! If you’d like to add new features, improve documentation, or report issues, please follow our [Contribution Guidelines](docs/contribution.md).

### Steps to Contribute:

1. 🍴 **Fork** the repository.
2. 🌿 **Create a new branch** (`git checkout -b feature/YourFeature`).
3. 🛠️ **Make your changes**.
4. 📤 **Push the branch** and open a pull request.

## Documentation 📖

For comprehensive guides and in-depth usage instructions, explore the documentation in the `docs/` folder:

- 📗 **[Getting Started Guide](docs/getting_started.md)**: Step-by-step setup and usage instructions.
- 🏗️ **[Project Architecture](docs/project_architecture.md)**: In-depth architectural overview of the framework.
- 🗄️ **[Database Setup](docs/database_setup.md)**: Instructions for configuring databases.
- 🔧 **[Contribution Guide](docs/contribution.md)**: Guidelines for contributing to the project.
- 📋 **[Naming Conventions](docs/naming_conventions.md)**: Best practices for writing Gherkin feature files and step definitions.

## License 📜

This project is licensed under the MIT License. See the [LICENSE](LICENSE.md) file for more details.

---

Thank you for using **Test in Go**! We hope this framework helps streamline your test automation and fosters collaboration in your development process. 🙌
# Contribution Guide

We welcome contributions to the **Test-in-Go** framework! Whether it's fixing bugs, adding new features, or improving the documentation, your help is greatly appreciated.

## How to Contribute

### 1. Fork the Repository

Start by forking the repository to your GitHub account:

1. Navigate to the repository on GitHub.
2. Click the "Fork" button at the top-right of the page.
3. Clone the forked repository to your local machine.

```bash
git clone https://github.com/your-username/Test-in-Go.git
```

### 2. Create a New Branch

Create a new branch for your feature or bug fix:

```bash
git checkout -b feature/new-feature
```

### 3. Make Your Changes

Make your changes in the appropriate files. If you're adding new functionality, be sure to:

- Write tests for the new feature.
- Ensure that all tests pass using `./scripts/run_tests.sh`.
- Update the documentation (if necessary).

### 4. Commit and Push

Once your changes are ready, commit them to your branch:

```bash
git add .
git commit -m "Add new feature"
git push origin feature/new-feature
```

### 5. Open a Pull Request

Go to the original repository on GitHub and open a **Pull Request** (PR) to the `main` branch. In the PR description, explain what the feature or fix does, and reference any relevant issue numbers.

We will review your PR and provide feedback if necessary. Once approved, your changes will be merged into the main branch!

---

## Code Style

- Use **Go modules** for dependency management.
- Follow Go's standard formatting (`go fmt`).
- Write clear and concise comments, especially for exported functions.
- Ensure your code is linted using a tool like **golangci-lint**.

---

## Running Tests

Be sure to run the full test suite before submitting your PR to ensure nothing is broken:

```bash
./scripts/run_tests.sh
```

---

## Thank You

Thank you for contributing to the project! Every bit of help is appreciated, and we value the time you take to improve the framework.
# Naming Conventions

This document outlines the naming conventions for files, variables, and other components in this project.

### File Naming
- Use **lowercase with underscores** (`_`).
  - E.g., `product_creation.feature`, `api_helpers.go`

### Directory Naming
- Use **lowercase**, no special characters.
  - E.g., `config`, `features`

### Variable Naming
- Use **camelCase** for variables.
  - E.g., `userDetails`, `apiResponse`

### Function Naming
- **PascalCase** for exported, **camelCase** for unexported functions.
  - E.g., `HandleRequest` (exported), `parseResponse` (unexported)

### Constants Naming
- Use **UPPERCASE with underscores**.
  - E.g., `MAX_CONNECTIONS`, `API_TIMEOUT`

### Test Function Naming
- Prefix with `Test` and use **PascalCase**.
  - E.g., `TestHandleRequest`

### Feature Step Definitions
- Use **camelCase** for methods matching feature steps.
  - E.g., `givenAUserCreatesAProduct`

### Struct and Interface Naming
- Use **PascalCase**.
  - E.g., `ProductDetails`, `StockBalanceData`

### JSON Key Naming
- Use **snake_case** for JSON keys in struct tags.
  - E.g., `json:"product_id"`

### SOAP Message Naming
- **PascalCase** for SOAP struct fields.
  - E.g., `ProductRequest`, `CustomerResponse`
- Use **camelCase** or **PascalCase** for XML tags.
  - E.g., `xml:"ProductID"`

### Kafka Message Naming
- **snake_case** for message keys/values.
  - E.g., `"product_id": "12345"`
- **lowercase with hyphens** for topic names.
  - E.g., `order-events`, `product-updates`

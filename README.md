# Go API Delivery Platform

A DevOps-focused project exploring Infrastructure as Code, containerization, and automated software delivery using **Go, Docker, Terraform, AWS, and GitHub Actions**.

The project builds on an existing Go REST API and focuses on developing a reproducible deployment workflow, from automated testing and container builds to provisioning cloud infrastructure.

> **Status:** In progress. The Go API, Docker configuration, and CI validation workflows are implemented. AWS infrastructure provisioning and automated deployment are upcoming stages.

## Tech Stack

| Area | Technologies |
|---|---|
| Application | Go, REST API |
| Containerization | Docker, multi-stage builds |
| Infrastructure as Code | Terraform, AWS Provider |
| Cloud Infrastructure | AWS EC2 (planned) |
| CI/CD | GitHub Actions |
| Testing | Go testing package, race detector, coverage, `go vet` |

## Architecture

The intended deployment architecture is:

```text
GitHub Repository
       |
       v
GitHub Actions
  - Formatting checks
  - Unit tests and race detection
  - Static analysis
  - Build verification
  - Terraform validation
       |
       v
AWS EC2 (planned)
       |
       v
Docker Container
       |
       v
Go REST API
  - Calculator endpoints
  - Health check
```

Terraform will manage the AWS infrastructure required to host the application. A future deployment workflow will automate delivery of new application versions.

## Project Structure

```text
.
├── api/
│   ├── cmd/
│   ├── internal/
│   ├── go.mod
│   └── Dockerfile
├── terraform/
│   ├── main.tf
│   └── .terraform.lock.hcl
├── .github/
│   └── workflows/
│       └── ci.yml
└── README.md
```

## Running the API Locally

**Prerequisites:** Go and Docker.

To run the API directly:

```bash
cd api
go run ./cmd/server
```

To run the test suite:

```bash
cd api
go test ./...
```

To run tests with race detection and generate a coverage report:

```bash
go test -race -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```

### Docker

Build the container image from the API directory:

```bash
cd api
docker build -t go-calculator-api .
```

Run the container:

```bash
docker run --rm -p 8080:8080 go-calculator-api
```

The API includes a health check endpoint that can be used to verify application availability.

> The port and startup command above assume the application listens on port 8080 and uses `cmd/server`. Adjust them to match the actual implementation.

## Continuous Integration

GitHub Actions automatically validates changes on pushes and pull requests.

### Go API Checks

- Verify source formatting using `gofmt`
- Run unit tests with race detection
- Generate test coverage data
- Run static analysis using `go vet`
- Verify that the application builds successfully

### Terraform Checks

- Validate Terraform formatting
- Initialize required providers without configuring a remote backend
- Validate Terraform configuration syntax and structure

These checks run independently and do not require AWS credentials or provision cloud resources.

## Infrastructure as Code

The `terraform/` directory contains the initial Terraform configuration using the AWS provider.

The planned infrastructure includes:

- An EC2 instance hosting the containerized Go API
- A security group restricting inbound network traffic
- EC2 bootstrap configuration for installing and running Docker
- Infrastructure outputs for accessing the deployed service

Terraform will be used to provision, update, and destroy the infrastructure through a repeatable workflow.

## Roadmap

- [x] Implement Go REST API and unit tests
- [x] Add application health check
- [x] Containerize API using a multi-stage Docker build
- [x] Configure Terraform AWS provider
- [x] Implement GitHub Actions CI for Go
- [x] Implement Terraform formatting and validation in CI
- [ ] Provision AWS EC2 infrastructure using Terraform
- [ ] Deploy containerized API to EC2
- [ ] Automate application deployment using GitHub Actions
- [ ] Configure basic application monitoring and logging
- [ ] Document deployment, troubleshooting, and infrastructure cleanup

## Project Goals

This project is intended to develop practical experience with:

- Infrastructure provisioning and lifecycle management
- Containerized application delivery
- CI/CD automation
- Cloud deployment troubleshooting
- Infrastructure validation and reproducibility
- AWS resource management and operational considerations

The implementation is intentionally incremental, prioritizing understanding of each infrastructure component before introducing additional services and complexity.

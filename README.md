# Deterministic State Authority (BSA) Core: Solving LLM State Drift

## Problem
In multi-agent LLM systems, "state drift" and hallucinations are critical issues. Agents often operate on outdated or conflicting views of the world, leading to inconsistent actions and corrupted data. Without a single source of truth that enforces strict versioning and integrity, autonomous systems degrade into chaos.

## Solution
The **Deterministic State Authority (BSA)** solves this by providing a centralized, high-performance state management service.
- **Version-Locked Live Index**: Ensures every agent sees a consistent, immutable snapshot of the state for any given version.
- **Go Reconciliation Loop**: A robust, concurrent-safe background process that validates proposals and enforces atomic commits to the canonical state.
- **Hybrid Architecture**: Leverages Go for the high-throughput core and Python for easy agent integration.

## Quickstart

### Prerequisites
- Docker (recommended) OR Go 1.21+
- Python 3.8+

### 1. Build and Run the Core Service (Go)

**Using Docker:**
```bash
docker build -t bsa-core .
docker run -p 8080:8080 bsa-core
```

**Using Go directly:**
```bash
cd bsa-core-go
go mod tidy
go run main.go
```

### 2. Run the Python Client Example

Install dependencies:
```bash
pip install requests pydantic notebook
```

Run the example notebook:
```bash
jupyter notebook example_usage.ipynb
```

## API Overview

- `GET /api/v1/state?version={v}`: Get the canonical state.
- `POST /api/v1/propose`: Submit a state change proposal.

## License
MIT

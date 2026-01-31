# Go Worker Data Transform Engine

A Go-based worker service that processes stock market data by aggregating multiple records into averaged summaries.

## Overview

This service receives stock data records via HTTP POST and returns aggregated results with averaged numeric fields. It's designed to work as a backend worker for the data-transform-engine system.

## API Endpoints

| Endpoint | Method | Description |
|----------|--------|-------------|
| `/health` | GET | Health check |
| `/process` | POST | Process stock records |

### POST /process

**Request:**
```json
{
  "records": [
    {
      "symbol": "^AXJO",
      "name": "S&P/ASX 200",
      "price": 8929.41,
      "change": 7.53,
      "change_percent": 0.084399,
      "day_high": 8948.55,
      "day_low": 8903.81,
      "previous_close": 8921.88
    }
  ]
}
```

**Response:**
```json
{
  "success": true,
  "data": {
    "symbol": "^AXJO",
    "name": "S&P/ASX 200",
    "price": 8929.41,
    "change": 7.53,
    "change_percent": 0.084399,
    "day_high": 8948.55,
    "day_low": 8903.81,
    "previous_close": 8921.88,
    "timestamp": "2026-01-31T01:30:00Z"
  }
}
```

## Quick Start

```bash
# Run the server
go run main.go

# Build and run
go build -o worker . && ./worker
```

## Configuration

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |

## Project Structure

```
go-worker-data-transform-engine/
├── main.go              # Server entry point
├── handlers/
│   └── process.go       # HTTP handlers
├── models/
│   └── stock.go         # Data structures
├── services/
│   └── aggregator.go    # Processing logic
└── go.mod
```

## Testing

```bash
go test ./...
```

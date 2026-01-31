# Go Worker Commands

## Development

```bash
# Run the server
go run main.go

# Run with custom port
PORT=9090 go run main.go
```

## Build

```bash
# Build binary
go build -o worker .

# Run the binary
./worker
```

## Testing

```bash
# Run all tests
go test ./...

# Run with verbose output
go test -v ./...

# Run with coverage
go test -cover ./...
```

## API Endpoints

```bash
# Health check
curl http://localhost:8080/health

# Process stock data
curl -X POST http://localhost:8080/process \
  -H "Content-Type: application/json" \
  -d '{"records":[{"symbol":"^AXJO","name":"S&P/ASX 200","price":100.0,"change":1.5,"change_percent":0.015,"day_high":101.0,"day_low":99.0,"previous_close":98.5}]}'
```

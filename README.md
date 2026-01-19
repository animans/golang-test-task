## Run
docker compose up --build

## Test
go test ./... -v

## API
POST /numbers
Body: {"value": 3}

Examples:
curl -X POST http://localhost:8080/numbers -H "Content-Type: application/json" -d '{"value":3}'

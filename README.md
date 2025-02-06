### event-booking-api-go

Simple API in Go using Gin and SQLite for study purposes.

Example `POST` request:

```bash
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-d '{
  "Name": "Devops na Praia fev/2025",
  "Description": "Encontro de devops no Porto.",
  "Location": "Porto, PT",
  "DateTime": "2025-02-15T09:00:00Z"
}'
```

Example `GET` request:

```bash
curl -X GET http://localhost:8080/events
```

Example `DELETE` request

```bash
curl -X DELETE http://localhost:8080/events/2
```

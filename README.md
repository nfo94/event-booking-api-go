### event-booking-api-go

Simple API in Go using Gin and SQLite. To run this project clone it in your machine
and run:

```bash
go mod tidy
```

And then run:

```bash
go run main.go
```

This will make the API available locally in http://localhost:8080/. If you wish to debug
this app in VSCode you can use the Debugger with the configuration in `.vscode`. Hit
`cmd/ctrl+shift+D` and click on play "Launch package".

Example for creating a user:

```bash
curl -X POST http://localhost:8080/signup \
-H "Content-Type: application/json" \
-d '{
  "email": "test1@gmail.com",
  "password": "totally-not-safe"
}'
```

Example login:

```bash
curl -X POST http://localhost:8080/login \
-H "Content-Type: application/json" \
-d '{
  "email": "test1@gmail.com",
  "password": "totally-not-safe"
}'
```

Grab the JWT token and use for protected routes. Example `POST` request for creating
an event:

```bash
curl -X POST http://localhost:8080/events \
-H "Content-Type: application/json" \
-H "Authorization: token" \
-d '{
  "Name": "Devops na Praia fev/2025",
  "Description": "Encontro de devops no Porto.",
  "Location": "Porto, PT",
  "DateTime": "2025-02-15T09:00:00Z"
}'
```

Example `GET` request to check all created events:

```bash
curl -X GET http://localhost:8080/events
```

Example `DELETE` request to delete an event:

```bash
curl -X DELETE http://localhost:8080/events/2
```

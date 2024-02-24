GopherDay 2024
===

## RESTful API

```bash
go run ./cmd/http
```

### v1

* `GET /v1/events` - List all events
* `POST /v1/events` - Create a new event

###  v2

* `GET /v2/events` - List all events

## gRPC

```bash
go run ./cmd/grpc
```

### Events

* `events.Events/List` - List all events

> You can use [grpcurl](https://github.com/fullstorydev/grpcurl) to interact with the gRPC server.
> Example: `grpcurl -plaintext localhost:8080 events.Events/List`

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

### v3

* `GET /v3/shipments/{id}` - Get a shipment by ID
* `POST /v3/shipments` - Create a new shipment
* `PUT /v3/shipments/{id}` - Update a shipment by ID
* `POST /v3/shipments/{id}/items` - Add an item to a shipment

## gRPC

```bash
go run ./cmd/grpc
```

### Events

* `events.Events/List` - List all events

> You can use [grpcurl](https://github.com/fullstorydev/grpcurl) to interact with the gRPC server.
> Example: `grpcurl -plaintext localhost:8080 events.Events/List`

## Future Discussions

* The `internal/event` is entity?
* Where to define the application errors is better?
* How to split into domains?

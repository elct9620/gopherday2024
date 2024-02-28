protoc:
	protoc --go_out=. --go_opt=paths=source_relative \
				 --go-grpc_out=. --go-grpc_opt=paths=source_relative \
				 pkg/events/events.proto

new_event:
	curl -XPOST -H 'Content-Type: application/json' -d '{}' localhost:8080/v1/events | jq

get_events:
	curl -XGET localhost:8080/v1/events | jq

get_events_v2:
	curl -XGET localhost:8080/v2/events | jq

new_shipment:
	curl -XPOST -H 'Content-Type: application/json' -d '{"id": "d489ca17-509f-4fe9-a000-23514cc0b9fb"}' localhost:8080/v3/shipments | jq

get_shipment:
	curl -XGET localhost:8080/v3/shipments/d489ca17-509f-4fe9-a000-23514cc0b9fb | jq

add_shipment_item:
	curl -XPOST -H 'Content-Type: application/json' -d '{"name": "Macbook Max M3 14inch"}' localhost:8080/v3/shipments/d489ca17-509f-4fe9-a000-23514cc0b9fb/items | jq

shipment_ship:
	curl -XPUT -H 'Content-Type: application/json' -d '{"state": "shipping"}' localhost:8080/v3/shipments/d489ca17-509f-4fe9-a000-23514cc0b9fb | jq

shipment_delivered:
	curl -XPUT -H 'Content-Type: application/json' -d '{"state": "delivered"}' localhost:8080/v3/shipments/d489ca17-509f-4fe9-a000-23514cc0b9fb | jq

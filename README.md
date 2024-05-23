# GOLANG CLEAN ARCH CHALLENGE

Implement list orders by REST, GRPC and GRAPHQL

## Running the project
```bash
docker-compose up -d

migrate -path=internal/infra/database/migrations  -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up

go mod tidy

cd cmd/ordersystem

go run main.go wire_gen.go  
```

## To execute challenge requisites
### REST
The api rest run on port `:8000`

- Open [api/create_order](api/create_order.http)
- Send request
- Open [api/list_orders](api/list_orders.http)
- Send request


### GRPC
The grpc server run on port `:50051`

```bash
evans -r repl

call ListOrders
```

### GRAPHQL
The graphql server run on port `:8080`

- Go to http://localhost:8080/
- Add the query below
- Execute query

```graphql
query listOrders {
  orders {
    id
    Price
    Tax
    FinalPrice
  }
}
```
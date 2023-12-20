# go-hexagonal


This is a laboratory project to learn about hexagonal architecture in Go.

This application is a simple service to manage "products".

The project defines interfaces for the main entity Product, it's service and repository. The idea is to have different implementations of the repository (in memory, postgres, mongo, etc) and the service (sync, async, etc) without changing the business logic.

It currently has one repository implementation (in memory) and one service implementation (sync).

There are lots of things to refactor and improve, but the idea is to keep it simple and focus on the concept.

There are two interfaces to interact with the project: a simple CLI and a REST API.

To run the cli use:

```bash
go run main.go cli -h
```

It should give a simple explanation of the commands available.

To run the webserver for the REST API use:

```bash
go run main.go http
```

It will start a webserver on port 9000. You can use the following endpoints:

* POST http://localhost:9000/product
```json
{
    "name": "Product 1",
    "price": 10.5
}
```

* GET localhost:9000/product

* GET localhost:9000/product/{id}

* PUT localhost:9000/product/{id}/enable

* PUT localhost:9000/product/{id}/disable

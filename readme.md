# GO RESTful Inventory System

## Features

* Stores actual stock of products
* To store product that will be stored into the inventory.
* To store product, quantity, notes of the products going out of inventory
* Shows a report to help analyze and make decision. This report is related to total inventory value.
* Shows a report to help analyze and make decision. This report is related to omzet / selling / profit.

## Getting Started

If this is your first time encountering Go, please follow [the instructions](https://golang.org/doc/install) to
install Go on your computer. Requires Go 1.5 or above.

After installing Go, run the following commands to download and install this project:

```shell
# install the proejct
go get github.com/andrideng/inventory-system

# install dep
$ curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

# fetch the dependent packages
cd $GOPATH/andrideng/inventory-system
dep ensure
```

Start Application

```shell
go run server.go
```

The application runs as an HTTP server at port 8080. It provides the following RESTful endpoints:

* `GET /api/`: welcoming text api
* `GET /api/ping`: a ping service mainly provided for health check purpose
* `GET /api/products`: list all products
* `GET /api/products/:sku`: get product based on the sku
* `POST /api/products`: create product
* `GET /api/incoming-goods`: list all incoming-goods
* `POST /api/incoming-goods`: create incoming-goods
* `PUT /api/outgoing-goods/:id`: update outgoing-goods based on id
* `GET /api/outgoing-goods`: list all outgoing-goods
* `POST /api/incoming-goods`: create incoming-goods
* `GET /api/reports/value-of-goods`: generate and response value of goods reports
* `GET /api/reports/sales?start_date=2019-01-01&end_date=2019-01-31`: generate and response sales report
* `GET /api/import-products`: import product based on storages/products.csv.

API Documentation [https://documenter.getpostman.com/view/528724/RztitVGv]


For example, if you access the URL `http://localhost:8080/api/ping` in a browser, you should see the browser
displays something like `PONG!`.

## Project Structure

This project divides the whole project into four main packages:

* `models`: contains the data structures used for communication between different layers.
* `services`: contains the main business logic of the application.
* `daos`: contains the DAO (Data Access Object) layer that interacts with persistent storage.
* `apis`: contains the API layer that wires up the HTTP routes with the corresponding service APIs.

[Dependency inversion principle](https://en.wikipedia.org/wiki/Dependency_inversion_principle)
is followed to make these packages independent of each other and thus easier to test and maintain.

The rest of the packages in the project are used globally:
 
* `app`: contains routing middlewares and application-level configurations
* `errors`: contains error representation and handling
* `util`: contains utility code

The main entry of the application is in the `server.go` file. It does the following work:

* load external configuration
* establish database connection
* instantiate components and inject dependencies
* start the HTTP server
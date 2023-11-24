# Project features
The project represents the backend of a customer relationship management (CRM) web application. As users interact with the app via some user interface, your server will support all of the functionalities:

- [x] Getting a list of all customers
- [x] Getting data for a single customer
- [x] Adding a customer
- [x] Updating a customer's information
- [x] Removing a customer
- [x] Swagger API Documentation 

## CRM Endpoints:
### API Documentation
You can find all required information [here](http://localhost:3000/swagger/index.html)

### This CRM Backend offers following endpoints:
```shell
curl localhost:3000/
curl localhost:3000/swagger/index.html
curl -X GET localhost:3000/customers
curl -X GET localhost:3000/customers/{id}
curl -X POST localhost:3000/customers/{id}
curl -X DELETE localhost:3000/customers{id}
```
Responses come with JSON.

## CRM Functions:
- getCustomers()
- getCustomer()
- addCustomer()
- updateCustomer()
- deleteCustomer()

# Development Documentation

## Best Practices
--------------

| Success Criteria | Specifications |
| --- | --- |
|Set up a proper Go environment | The project requires only a simple `go run` command to launch the application.
|Document the project in the README |The project README contains a description of the project, and also includes instructions for installation, launch, and usage. |
|Organize and write clean code | <ul><li>Syntax and semantics of language features (e.g., functions, variables, loops, etc.) are well-formed and free from errors or warnings in the console </li><li>Data structures, handlers, routes, imports, and other assets are organized logically (e.g., grouped together) and are easy to find</li></ul> |
|Build an intuitive user experience |Users can interact with the application (i.e., make API requests) by simply using Postman or cURL. |

## Data
----

| Success Criteria | Specifications |
| --- | --- |
|Create a Customer struct |Each customer includes: <ul><li>ID</li><li>Name</li><li>Role</li><li>Email</li><li>Phone</li><li>Contacted (i.e., indication of whether or not the customer has been contacted)</li></ul> Data is mapped to logical, appropriate types (e.g., Name should not be a `bool`). |
| Create a mock "database" to store customer data |Customers are stored appropriately in a basic data structure (e.g., slice, map, etc.) that represents a "database." |
|Seed the database with initial customer data |The "database" data structure is non-empty. That is, prior to any CRUD operations performed by the user (e.g., adding a customer), the database includes at least three existing (i.e., hard-coded") customers. |
| Assign unique IDs to customers in the database |Customers in the database have unique ID values (i.e., no two customers have the same ID value). |

## Server
------

| Success Criteria | Specifications |
| --- | --- |
|Serve the API locally |The API can be accessed via `localhost`. |
|Create RESTful server endpoints for CRUD operations | The application handles the following 5 operations for customers in the "database": <ul><li>**Getting a single customer** through a `/customers/{id}` path</li><li>**Getting all customers** through a the `/customers` path</li><li>**Creating a customer** through a `/customers` path</li><li>**Updating a customer** through a `/customers/{id}` path</li><li>**Deleting a Customer** through a `/customers/{id}` path</li></ul> Each RESTful route is associated with the correct HTTP verb.|
|Return JSON in server responses |The application leverages the `encoding/json` package to return JSON values (i.e., not text, etc.) to the user. |
|Serve static HTML at the home ("/") route |The home route is a client API endpoint, and includes a brief overview of the API (e.g., available endpoints). Note: This is the only route that does *not* return a JSON response. |
|Set up and configure a router |The application uses a router (e.g., `gorilla/mux`, `http.ServeMux`, etc.) that supports HTTP method-based routing and variables in URL paths. |
|Create and assign handlers for requests |The Handler interface is used to handle HTTP requests sent to defined paths. There are five routes that return a JSON response, and are each is registered to a dedicated handler: <ul><li>`getCustomers()`</li><li>`getCustomer()`</li><li>`addCustomer()`</li><li>`updateCustomer()`</li><li>`deleteCustomer()`</li></ul>
|Includes basic error handling for non-existent customers  |If the user queries for a customer that doesn't exist (i.e., when getting a customer, updating a customer, or deleting a customer), the server response includes: <ul><li> A `404` status code in the header </li><li>`null` *or* an empty JSON object literal *or* an error message</li></ul>|
|Set headers to indicate the proper media type |An appropriate `Content-Type` header is sent in server responses. |
|Read request data |The application leverages the `io/ioutil` package to read I/O (e.g., request) data. |
|Parse JSON data |The applications leverages the `encoding/json` package to parse JSON data. |

### Suggestions to Make Your Project Stand Out

1.  Create an additional endpoint that updates customer values in a batch (i.e., rather than for a single customer).
2.  Upgrade the mock database to a real database (e.g., PostgreSQL).
3.  Deploy the API to the web.

# Credits to:
- [MitchDresdner](https://github.com/MitchDresdner/winestore/tree/master) for the Property and Parsing Code Example

# Sources I toke coding approaches and ideas from:
- stackoverflow
- [go tutorial (official)](https://github.com/golang/go/wiki/SQLInterface)
- https://www.soberkoder.com/go-rest-api-mysql-gorm/
- https://www.soberkoder.com/swagger-go-api-swaggo/
- https://www.soberkoder.com/go-rest-api-gorilla-mux/
- in case i missed somthing / someone please open a pr and i'll add you as inspirations


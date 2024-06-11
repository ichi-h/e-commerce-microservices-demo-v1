# e-commerce-microservices-demo-v1

This project is a demonstration of a microservices architecture implementation for an e-commerce application.  
The system design of this application is somewhat overstated and is not necessarily the correct representation for a real e-commerce application.

Currently, authentication or authorization mechanisms are not implemented.

## Technology Stack

- Container Technology
  - Docker
  - Docker Compose
- Container Orchestration
  - Kubernetes (minikube)
- Database
  - PostgreSQL
  - MongoDB
- Streaming Platform
  - Apache Kafka
- Server Side
  - Golang
    - Echo
  - Node.js
    - Hono
- API
  - GraphQL
  - gRPC
- Frontend
  - Coming soon...

## Server Side Architecture Diagram

```mermaid
flowchart BT
  Client["Client"]
  BFF["BFF\n(Hono)"]

  subgraph Product["Product (CRUD)"]
    ProductService["Service\n(Golang)"]
    ProductDB["DB\n(postgreSQL)"]
  end

  subgraph Order["Order (CRUD)"]
    OrderService["Service\n(Golang)"]
    OrderDB["DB\n(postgreSQL)"]
  end

  subgraph Cart["Cart (CQRS + ES)"]
    CartCommandService["CommandService\n(Golang)"]
    CartQueryService["QueryService\n(Golang)"]
    CartEventHandler["EventHandler\n(Golang)"]
    MessagingQueue["Messaging Queue\n(Apache Kafka)"]
    CartEventStore["Event Store\n(MongoDB)"]
    CartReadDB["Read DB\n(postgreSQL)"]
  end

  subgraph User["User (CQS + ES)"]
    UserService["Service\n(Golang)"]
    UserDB["DB\n(postgreSQL)"]
    subgraph UserDB["DB (postgreSQL)"]
      UserEventTable["Event Table"]
      UserStateTable["State Table"]
    end
  end

  Client <--"GraphQL"--> BFF
  BFF <--"gRPC"--> ProductService
  BFF <--"gRPC"--> CartCommandService
  BFF <--"gRPC"--> CartQueryService
  BFF <--"gRPC"--> UserService
  BFF <--"gRPC"--> OrderService
  ProductService <--> ProductDB
  CartCommandService <--> CartEventStore
  CartCommandService --> MessagingQueue --> CartEventHandler --> CartReadDB
  CartQueryService <--> CartReadDB
  UserService <--> UserDB
  OrderService <--> OrderDB
```

## Domain Modeling

### Class Diagram

```mermaid
classDiagram
  class User {
    <<Abstract>>
  }

  class Customer {
  }

  class Order {
  }

  class OrderDetail {
  }

  class Cart {
  }

  class CartItem {
  }

  class Product {
  }

  User <|-- Customer
  Customer "1" *-- "0..*" Order
  Customer "0..1" *-- "1" Cart
  note for Cart "A cart can exist without a customer, but \nwhen a customer deletes his/her account, \nthe cart associated with it also disappears."
  Order "1" *-- "1..*" OrderDetail
  Product "1" o-- "1..*" CartItem
  Product "1" o-- "1..*" OrderDetail
  Cart "1" *-- "0..*" CartItem
```

### Use Case Diagram

```mermaid
graph LR
  Customer(Customer)

  subgraph User
    Login["Login (not implemented)"]
    Logout["Logout (not implemented)"]
    ViewProfile
  end

  subgraph Product
    ViewProducts
  end

  subgraph Order
    Checkout
    CreateOrder
    ViewOrders
  end

  subgraph Cart
    AddProduct
    RemoveProduct
    ChangeProductQuantity
    ViewCart
    ResetCart
  end

  Customer --> ViewProfile
  Customer --> ViewProducts
  Customer --> AddProduct
  Customer --> RemoveProduct
  Customer --> ChangeProductQuantity
  Customer --> ViewCart
  Customer --> ResetCart
  Customer --> Checkout
  Customer --> ViewOrders
  Checkout -.-> ResetCart
  Checkout -.-> CreateOrder
```

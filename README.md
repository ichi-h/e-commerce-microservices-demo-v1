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
- Proxy Server
  - Envoy
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
  Client

  subgraph BFF
    BFFProxy["Proxy\n(Envoy)"]
    BFFService["Service\n(Hono)"]
    BFFProxy <--> BFFService
  end

  subgraph Product["Product (CRUD)"]
    ProductProxy["Proxy\n(Envoy)"]
    ProductService["Service\n(Golang)"]
    ProductDB["DB\n(postgreSQL)"]
    
    ProductProxy <--> ProductService 
  end

  subgraph Order["Order (CRUD)"]
    OrderProxy["Proxy\n(Envoy)"]
    OrderService["Service\n(Golang)"]
    OrderDB["DB\n(postgreSQL)"]
    OrderProxy <--> OrderService
  end

  subgraph Cart["Cart (CQRS + ES)"]
    CartCommandProxy["Proxy\n(Envoy)"]
    CartQueryProxy["Proxy\n(Envoy)"]
    CartCommandService["CommandService\n(Golang)"]
    CartQueryService["QueryService\n(Golang)"]
    CartEventHandler["EventHandler\n(Golang)"]
    MessagingQueue["Messaging Queue\n(Apache Kafka)"]
    CartEventStore["Event Store\n(MongoDB)"]
    CartReadDB["Read DB\n(postgreSQL)"]
    CartCommandProxy <--> CartCommandService
    CartQueryProxy <--> CartQueryService
    
  end

  subgraph User["User (CQS + ES)"]
    UserProxy["Proxy\n(Envoy)"]
    UserService["Service\n(Golang)"]
    UserDB["DB\n(postgreSQL)"]
    UserProxy <--> UserService
    subgraph UserDB["DB (postgreSQL)"]
      UserEventTable["Event Table"]
      UserStateTable["State Table"]
    end
  end

  %% Endpoint Discovery Service (planning to implement in the future)
  %% EDS
  %% BFFProxy <-...-> EDS
  %% ProductProxy <-...-> EDS
  %% CartCommandProxy <-...-> EDS
  %% CartQueryProxy <-...-> EDS
  %% UserProxy <-...-> EDS
  %% OrderProxy <-...-> EDS

  Client <--"GraphQL"--> BFFProxy
  BFFProxy <--"gRPC"---> ProductProxy
  BFFProxy <--"gRPC"---> CartCommandProxy
  BFFProxy <--"gRPC"---> CartQueryProxy
  BFFProxy <--"gRPC"---> UserProxy
  BFFProxy <--"gRPC"---> OrderProxy
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

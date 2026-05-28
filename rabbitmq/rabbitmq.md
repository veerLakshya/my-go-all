## Erlang - general purpose, concurrent, functional programming language

### Features

- Concurrency
- Fault tolerance
- Distributed
- Hot code upgrades

### Significance of Erlang in RabbitMQ

- Concurrency and Scalability
- Fault Tolerance and Reliability
- Distributed Systems
- Hot Code Upgrades (without downtime)
- Erlang Ecosystem and Libraries

## What is a Message Broker?

A message broker is a software component that facilitates communication btw diff applications. It decouples producers (senders) from consumers (receivers), allowing systems to communicate asynchronously and reliably even if they operate on different platforms or languages. Serves as a middleman mostly

### Key Functionalities -

- Message Queuing
- Message Routing based on predefined rules
- Message Transformation
- Handling Message Delivery

### Why to use message brokers?

- Decoupling
- Async communication
- Scalability
- fault tolerance

## RabbitMQ - developed my Rabbit Technologies

A message broker tool (not a framework/library) which allows async communication, exchange data and process tasks between applications and microservices by acting as the middleman.

- Handles Async communication
- enables loose coupling
- support scalability
- ensure fault tolerance
- integrate with other systems

### Problems before RabbitMQ -

- Tightly coupled systems
- Asynchronous Communication
- Reliability Issues

### Similar Technologies/tools before Rmq -

- JMS (Java Message Service)
- ActiveMQ (struggled in high memory usecases)
- ZeroMQ (p2p, doesnt require a central broker)
- IBM MQ (formerly WebSphere MQ)

### Features of RabbitMQ -

- Decouples Components
- Provides Asynchronous Messaging
- Enhances Scalability
- Ensures Fault Tolerance
- Supports Multiple Protocols

### What problems does it solve?

- message queuing
- message routing
- guaranteed message delivery
- load balancing across consumers
- real time communication

### A tool for modern architectures-

- microservices
- event driven applications
- real-time systems

### RabbitMQ concepts-

- Producer : an app/ervice that sends messages to RabbitMQ
- Consumer : an app/service that receives and processes messages from Rmq
- Queue : where Rmq stores messages until consumed(FIFO based)
- Exchange : routing mechanism. Types: Direct, Fanout, Topic, Headers
- Binding : link between an exchange and a queue
- Routing Key : a string that helps the exchange determine where to send a message
  - For Direct Exchanges: the routing key is compared to the binding key of queues -
  - For Topic Exchanges: the routing key acts as a pattern-matching string
- Message: it consists of two parts:
  - Body : the actual data or content of the message
  - Headers / Properties : Metadata that can be attached to the message
- Acknowledgement (Ack) : removes or reques a message depending on Ack
- Virtual Hosts (vhosts): a logical seperation withing the RabbitMQ broker
- Clustering : Allowing you to distribute the workload across multiple servers
- Management Plugin : Rmq provides management plugin that allows you to manage and monitor Rmq through a web-based interface

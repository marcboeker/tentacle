# Tentacle

Tentacle is a software-defined mesh network that uses the QUIC protocol as the underlying transport mechanism. It is designed to be used in a peer-to-peer fashion, where each peer can connect to any other peer in the network. The network is fully decentralized and does not rely on any central server for traffic passthrough.

The server is responsible for managing the network configuration and keeps track of the peers. It is also responsible for distributing the network configuration to the peers.

## Getting Started

There is a Docker Compose file included in the repository that defines the server and two clients. The server is configured to listen on port 1337 and the clients are configured to connect to the server on port 1337. The clients are also configured to listen on ports 4001 and 4002 respectively.

## Prerequisites

- Go 1.21 or later
- Docker with Docker Compose

## Installation

1. Clone the repository:
   `git clone https://github.com/marcboeker/tentacle.git`

2. Navigate to the project directory:
   `cd tentacle`

3. Build the Docker images:
   `docker-compose build`

## Running the Application

To start the server and the twp nodes, run:

`docker-compose up`

This will start the server on port 1337 and two nodes on ports 4001 and 4002 respectively.

## Project Structure

The project is structured as follows:

- `cert/`: Contains the code for handling certificates.
- `client/`: Contains the client code. This includes the code for setting up the TUN interface and the transport layer.
- `cmd/`: Contains the main entry points for the server and the clients.
- `protocol/`: Contains the protocol buffers definitions and the generated Go code.
- `server/`: Contains the server code.

## Code Overview

Here are some of the most important files and their purpose:

- `interface_darwin.go` and `interface_linux.go`: These files contain the code for setting up a TUN interface on macOS and Linux respectively. They define the `SetupInterface`, `AddIP`, and `AddRoute` functions.
- `transport_wrapper.go`: This file contains the `TransportWrapper` type which manages all incoming and outgoing peer connections. It defines the `AddStream`, `RemoveStream`, `GetOrConnect`, and `Serve` methods.
- `transport.go`: This file contains the `Transport` type which handles the transport layer. It defines the `NewTransport` function.
- `server.go`: This file contains the `Server` type which manages the network and the peers. It defines the `NewServer` function.

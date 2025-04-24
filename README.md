# Secure TCP Chat App in GO

## Goal
Build a client-server chat app over TCP, encrypted with SSL/TLS.

## How it works
- This project is a secure chat app using Transport Layer Security over Transmission Control Protocol.
- Main components are `server.go`: Listens for client connections, and `client.go`: Connects to the server and exchanges messages.
- In the app `localhost`(`127.0.0.1`for IPv4) is the IP address.
  - IP ensures that data goes to the correct machine, and in the app both client and server are on the same machine, so `localhost` points to the same device. 
- TCP creates a reliable, ordered, error-checked communication stream.
  - In this app, TCP handles reliable delivery.
  - TLS adds encryption and security.
  - `tls.Dial("tcp", "localhost:8080", ...)`, `tls.Listen("tcp", ":8080", ...)`
  - When the client connects to `localhost:8080`, TCP establishes a connection between the client and the server.
  - TCP splits messages into segments, ensures they arrive, and reassembles them in order.
- In this app, data is the messages exchanged between client and server.
  - This data is sent over the TLS connection, which wraps a TCP socket (`localhost:8080` = IP + port)
  - The data is protected by encrypted using TLS, and this ensures no one can see or tamper with messages.
```
You (client)        -->  TCP/IP  -->       Server
 "hi" (data)        Encrypted via TLS       "Hola..." (data)
```


### Repository Structure
```
secure-chat-app/
├── certs/                  # SSL/TLS certificates and keys
│   ├── server.crt          # Server's SSL certificate
│   ├── server.key          # Server's private key
├── server.go               # Server-side code (TCP server with TLS/SSL)
├── client.go               # Client-side code (TCP client with TLS/SSL)
├── go.mod                  
├── go.sum                  
└── README.md               # This doc
```

### Tools
net: to open TCP sockets.		
crypto/tls: to encrypt communication with TLS.		
bufio: for buffered input/output.		
Self-signed certificate (for local TLS).		
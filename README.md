# Secure TCP Chat App in GO

## Goal
Build a client-server chat app over TCP, encrypted with SSL/TLS.

### Repository Structure
secure-chat-app/
├── certs/                  # SSL/TLS certificates and keys
│   ├── server.crt          # Server's SSL certificate
│   ├── server.key          # Server's private key
├── server.go               # Server-side code (TCP server with TLS/SSL)
├── client.go               # Client-side code (TCP client with TLS/SSL)
├── go.mod                  
├── go.sum                  
└── README.md               # This doc
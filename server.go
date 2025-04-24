package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	// Send a welcome message to the client
	conn.Write([]byte("Welcome to the secure chat server!\n"))

	for {
		buffer := make([]byte, 1024)

		// Read the incoming messagae
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from client", err)
			return
		}

		message := string(buffer[:n])
		fmt.Printf("Client says:%s\n", message)

		message = strings.TrimSpace(message) // Trim any extra whitespace or newlines

		// Respond to the client based on their message
		if message == "hi" {
			conn.Write([]byte("Hola, cómo puedo ayudarte hoy?\n"))
		} else if message == "bye" {
			conn.Write([]byte("Ciao, Buonanotte!\n"))
			// Properly close the connection and break out of the loop
			conn.Close()
			break
		} else {
			_, err = conn.Write([]byte("Server received your message: " + message + "\n"))
			if err != nil {
				log.Println("Error sending message: ", err)
				return
			}
		}
	}
}

func main() {
	// Load the server's TLS certificate and private key
	cert, err := tls.LoadX509KeyPair("./certs/server.crt", "./certs/server.key")
	if err != nil {
		log.Fatalf("Error loading certificate and key: %v", err)
	}

	// Set up the TCP server with SSL/TLS
	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	listner, err := tls.Listen("tcp", ":8080", config)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
	defer listner.Close()

	// Print that the server is ready
	fmt.Println("Server is listening on port 8080...")

	// Loop to accept incoming connections and handle them concurrently
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		// Handle new client connection
		go handleClient(conn)
	}
}

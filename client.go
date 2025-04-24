package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"os"
)

func main() {
	// Connect to the server securely using TLS (disabling certificate verification)
	conn, err := tls.Dial("tcp", "localhost:8080", &tls.Config{
		InsecureSkipVerify: true, // Disable certificate verification
	})
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// Read and print the server's greeting message
	message, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Server says:", message)

	// Start chatting: send messages to the server and print responses
	for {
		fmt.Print("You: ")
		text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		conn.Write([]byte(text))

		// Read the response from the server
		response, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Server:", response)

		// Exit condition
		if text == "bye\n" {
			fmt.Println("Exiting chat...")
			break
		}
	}
}

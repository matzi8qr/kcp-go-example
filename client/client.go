package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
	"log"
)

const Raddr = "127.0.0.1:8095"

func main() {

	// Create a KCP connection to the server
	conn, err := kcp.DialWithOptions(Raddr, nil, 10, 3)
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}
	defer conn.Close()

	// Send a message to the server
	message := "Hello from client!"
	_, err = conn.Write([]byte(message))
	if err != nil {
		log.Println("Error writing to server:", err)
		return
	}

	// Receive and print the server's response
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Println("Error reading from server:", err)
		return
	}

	serverResponse := string(buffer[:n])
	fmt.Printf("Received from server: %s\n", serverResponse)
}

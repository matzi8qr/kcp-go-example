package main

import (
	"fmt"
	"github.com/xtaci/kcp-go"
	"log"
)

const Laddr = "127.0.0.1:8095"

func handleClient(conn *kcp.UDPSession) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			log.Println("Error reading from client:", err)
			return
		}

		clientMessage := string(buffer[:n])
		log.Printf("Received from client: %s\n", clientMessage)

		// Send a response to the client
		response := "Hello from server!"
		_, err = conn.Write([]byte(response))
		if err != nil {
			log.Println("Error writing to client:", err)
			return
		}
	}
}

func main() {

	// Create a KCP listener
	listener, err := kcp.ListenWithOptions(Laddr, nil, 10, 3)
	if err != nil {
		log.Fatal("Error creating KCP listener:", err)
	}
	defer listener.Close()

	fmt.Printf("Server listening on %s\n", Laddr)

	for {
		// Accept incoming connections
		conn, err := listener.AcceptKCP()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the client in a separate goroutine
		go handleClient(conn)
	}
}

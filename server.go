
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	

	"github.com/xtaci/kcp-go"
)

var clients = make(map[string]*kcp.UDPSession)

func handleConnection(conn *kcp.UDPSession, addr *net.UDPAddr) {
	scanner := bufio.NewScanner(conn)
	for {
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()
		log.Println("Received from", addr, ":", message)

		// Broadcast the message to all connected clients
		for _, client := range clients {
			fmt.Fprintln(client, message)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from", addr, ":", err)
	}

	// Remove the client from the clients map
	delete(clients, addr.String())
	log.Println("Client", addr, "disconnected")
}

func main() {
	listener, err := kcp.ListenWithOptions(":8888", nil, 10, 3)
	if err != nil {
		log.Fatal("Failed to create listener:", err)
	}
	defer listener.Close()

	fmt.Println("KCP chat server started on :8888")

	for {
		conn, err := listener.AcceptKCP()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		addr := conn.RemoteAddr().(*net.UDPAddr)
		clients[addr.String()] = conn

		log.Println("Client", addr, "connected")

		go handleConnection(conn, addr)
	}
}

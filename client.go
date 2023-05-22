
package main

import (
	"bufio"
	"fmt"
	"log"
	
	"os"

	"github.com/xtaci/kcp-go"
)

func main() {
	fmt.Println("Enter the server IP:")
	var ip string
	fmt.Scanln(&ip) // Read the IP address from the user

	fmt.Println("Enter the port:")
	var port string
	fmt.Scanln(&port)
	fmt.Println(ip + ":" + port)
	conn, err := kcp.DialWithOptions(ip + ":" + port, nil, 10, 3)
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		for {
			if !scanner.Scan() {
				break
			}

			message := scanner.Text()
			fmt.Println("Received:", message)
		}

		if err := scanner.Err(); err != nil {
			log.Println("Error reading from server:", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}

		message := scanner.Text()

		// Send the message to the server
		_, err = fmt.Fprintln(conn, message)
		if err != nil {
			log.Println("Error sending message to server:", err)
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading from stdin:", err)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
)

type Client chan<- string

var (
	incomimgClients = make(chan Client)
	leavingClients  = make(chan Client)
	messages        = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

// Client -> Server -> HandleConnection(Client1)

func HandleConnection(conn net.Conn) {
	defer conn.Close()
	message := make(chan string)
	go MessageWrite(conn, message)
	// Client: 3560 platzi.com, 38
	// platzi.com:38
	clientName := conn.RemoteAddr().String()
	message <- fmt.Sprintf("Welcome to the server, your name is: %s\n", clientName)
	messages <- fmt.Sprintf("New client is here %s\n", clientName)
	incomimgClients <- message

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messages <- fmt.Sprintf("%s: %s", clientName, inputMessage.Text())
	}

	leavingClients <- message
	messages <- fmt.Sprintf("%s: Said goodbye!", clientName)
}

func MessageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func Broadcast() {
	clients := make(map[Client]bool)
	for {
		select {
		case message := <-messages:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomimgClients:
			clients[newClient] = true
		case leavingClient := <-leavingClients:
			delete(clients, leavingClient)
			close(leavingClient)
		}
	}
}

func main() {
	flag.Parse()
	listner, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	go Broadcast()
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go HandleConnection(conn)
	}
}

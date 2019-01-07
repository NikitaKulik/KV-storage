/**

KV server with options for run
-p -> port were server will be run (default 9090)

*/

package main

import (
	"bufio"
	"flag"
	"io"
	"kv_storage/codebase/storage"
	"log"
	"net"
	"strings"
)

type command struct {
	fields []string
	result chan string
}

// Function for parse arguments and run KV server
func main() {
	serverPort := flag.String("p", "9090", "port")
	flag.Parse()

	handler, err := net.Listen("tcp", "localhost:"+*serverPort)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Server was started -> localhost:" + *serverPort)
	defer handler.Close()

	commands := make(chan command)
	go handleCommand(commands)

	for {
		connection, err := handler.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(commands, connection)
	}
}

// Function for handling connection from client to server
func handleConnection(commands chan command, connection net.Conn) {
	defer log.Println("Connection closed")
	defer connection.Close()

	scanner := bufio.NewScanner(connection)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)

		result := make(chan string)
		commands <- command{
			fields: fields,
			result: result,
		}
		_, err := io.WriteString(connection, <-result+"\n")
		if err != nil {
			log.Panic(err)
		}
	}

}

// Function for handling command from client to server
func handleCommand(commands chan command) {
	for cmd := range commands {

		if len(cmd.fields) < 1 {
			cmd.result <- ""
			continue
		}
		if len(cmd.fields) < 2 {
			cmd.result <- "Expected at least 2 arguments"
			continue
		}

		switch cmd.fields[0] {
		case "get":
			cmd.result <- storage.Get(cmd.fields[1], false)
		case "set":
			storage.Set(cmd.fields[1], cmd.fields[2], false)
			cmd.result <- ""
		case "del":
			storage.Del(cmd.fields[1], false)
			cmd.result <- ""
		default:
			cmd.result <- "Invalid command"
		}
	}
}

/**

KV server with options for run
-p -> port were server will be run (default 9090)

*/

package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
	"strings"
)

const JsonPath = "./server/storage.json"
const JsonTestPath = "./storage_test.json"

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
		case "GET":
			cmd.result <- Get(cmd.fields[1], false)
		case "SET":
			Set(cmd.fields[1], cmd.fields[2], false)
			cmd.result <- "OK"
		case "DEL":
			Del(cmd.fields[1], false)
			cmd.result <- "OK"
		default:
			cmd.result <- "Invalid command"
		}
	}
}

// Function for convert data from JSON to map
func getStorage(isTestMode bool) map[string]string {
	storagePath := JsonPath
	if isTestMode {
		storagePath = JsonTestPath
	}

	jsonFile, err := os.Open(storagePath)
	if err != nil {
		log.Panic(err)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Panic(err)
	}

	data := make(map[string]string)
	if len(bytes) > 0 {
		err = json.Unmarshal([]byte(bytes), &data)
		if err != nil {
			log.Panic(err)
		}
	}

	return data
}

// Function for "SET" handling
func Set(key string, value string, isTestMode bool) {
	storagePath := JsonPath
	if isTestMode {
		storagePath = JsonTestPath
	}

	storage := getStorage(isTestMode)
	storage[key] = value
	jsonData, _ := json.Marshal(storage)
	err := ioutil.WriteFile(storagePath, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}

// Function for "GET" handling
func Get(key string, isTestMode bool) string {
	storage := getStorage(isTestMode)
	value := storage[key]
	log.Printf(value)
	return value
}

// Function for "DEL" handling
func Del(key string, isTestMode bool) {
	storagePath := JsonPath
	if isTestMode {
		storagePath = JsonTestPath
	}

	storage := getStorage(isTestMode)
	delete(storage, key)
	jsonData, _ := json.Marshal(storage)
	err := ioutil.WriteFile(storagePath, jsonData, 0644)
	if err != nil {
		log.Panic(err)
	}
}

/**

KV client with options for server connection
-p -> port where server was run (default 9090)
-h -> host where server was run (default 127.0.0.1)

*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

//Function for parse client arguments and connect to server
func main() {
	serverPort := flag.String("p", "9090", "port")
	serverHost := flag.String("h", "127.0.0.1", "host")
	flag.Parse()

	conn, _ := net.Dial("tcp", *serverHost+":"+*serverPort)
	fmt.Printf("Client connect to -> " + *serverHost + ":" + *serverPort + "\n")
	fmt.Println("Please enter the command:")
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Print("Text to send: " + text)
		_, err := fmt.Fprintf(conn, text)
		if err != nil {
			log.Panic(err)
		}
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Server response: " + message)
	}
}

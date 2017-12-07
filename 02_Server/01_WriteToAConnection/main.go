package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080") //Create a Listener
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept() //Create a connection if there is a request
		if err != nil {
			log.Fatalln(err)
		}
		io.WriteString(conn, "Hello")             //Writing to the connection
		fmt.Fprintln(conn, "Welcome to my world") //connection implements reader and writer interface and since fprintf requires a writer we can use it
		fmt.Fprintf(conn, "%v", "Everything is different now")
		conn.Close()
	}
}

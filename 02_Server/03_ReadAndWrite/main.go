package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {

	err := conn.SetDeadline(time.Now().Add(5 * time.Second)) //Setting a timeout for the connection
	if err != nil {
		log.Println(err)
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "I heard you say %s\n", ln)
	}
	defer conn.Close()

	fmt.Println("Control got here")
}

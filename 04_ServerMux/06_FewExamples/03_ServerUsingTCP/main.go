package main

import (
	"io"
	"log"
	"net"
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
		//write to a connection
		io.WriteString(conn, "I hear you loud and clear")
		conn.Close()
	}

}

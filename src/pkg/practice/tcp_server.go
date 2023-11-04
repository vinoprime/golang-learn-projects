package practice

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func StartServer() {

	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}

		// io.WriteString(conn, "\n Hello from tcp")
		// fmt.Println(conn, "How are ypu?")
		// conn.Close()

		go handle(conn)

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
}

package main

import (
	"bufio"
	"fmt"
	"net"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	conn, errc := net.Dial("tcp", "140.112.42.221:12015")
	check(errc)
	defer conn.Close()

	writer := bufio.NewWriter(conn)
	len, errw := writer.WriteString("Hello World!\n")
	check(errw)
	fmt.Printf("Send a string of %d bytes\n", len)
	writer.Flush()

	reader := bufio.NewReader(conn)
	message, errr := reader.ReadString('\n')
	check(errr)
	fmt.Printf("Server replies: %s", message)
}

package main

import "fmt"
import "bufio"
import "net"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12015")
	conn, _ := ln.Accept()
	defer ln.Close()
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	message := ""
	if scanner.Scan() {
		message = scanner.Text()
		fmt.Println(message)
	}
	// for scanner.Scan() {
	// 	message = scanner.Text()
		
	// 	fmt.Println(message)
	// 	if message == "." {
	// 		break
	// 	}

	// }

	writer := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received\n", len(message))
	_, errw := writer.WriteString(newline)
	check(errw)
	writer.Flush()
}
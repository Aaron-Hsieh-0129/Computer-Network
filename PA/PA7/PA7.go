package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func handleConnection(c net.Conn) {
	reader := bufio.NewReader(c)
	req, err := reader.ReadString('\n')
	check(err)

	tokens := strings.Split(req, " ")
	fileName := tokens[1]
	processFile(fileName)
	c.Close()
}

func processFile(file string) {
	if file, errStat := os.Stat(file[1:]); errStat != nil {
		fmt.Println("File not found")
	} else {
		fmt.Println("File size = " + strconv.Itoa(int(file.Size())))
	}
}

func main() {
	fmt.Println("Launching server...")
	ln, _ := net.Listen("tcp", ":12015")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		defer conn.Close()

		go handleConnection(conn)
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
)

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

	// output file
	outputFile := "whatever.txt"
	f, err := os.Create(outputFile)
	check(err)
	defer f.Close()

	// reader and writer
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(f)
	count := 1

	// get read bits length
	line, errr := reader.ReadString('\n')
	check(errr)
	bytesToRead, errL2 := strconv.Atoi(line[0 : len(line)-1])
	check(errL2)
	fileBytes := bytesToRead

	for bytesToRead != 0 {
		line, errr = reader.ReadString('\n')
		check(errr)

		// combine message
		num := strconv.Itoa(count)
		combine := num + " " + line

		// write message
		writer.WriteString(combine)
		count++

		bytesToRead -= len(line)
	}
	writer.Flush()

	fi, err2 := f.Stat()
	check(err2)

	fmt.Printf("Upload file size: %d\n", fileBytes)
	fmt.Printf("Output file size: %d\n", fi.Size())

	writer2 := bufio.NewWriter(conn)
	newline := fmt.Sprintf("%d bytes received, %d bytes file generated\n", fileBytes, fi.Size())
	_, errw := writer2.WriteString(newline)
	check(errw)
	writer2.Flush()
}

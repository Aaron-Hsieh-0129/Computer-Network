package main

import "fmt"
import "os"
import "bufio"
import "net"
import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// open input file
	inputFile := "" 
	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &inputFile)

	// connect to server
	conn, errc := net.Dial("tcp", "140.112.42.221:12000")
	check(errc)
	defer conn.Close()

	// scan file
	fInput, err := os.Open(inputFile)
	check(err)
	defer fInput.Close()

	info, errStat := os.Stat(inputFile)
	check(errStat)

	writer := bufio.NewWriter(conn)
	writer.WriteString(strconv.Itoa(int(info.Size())) + "\n")
	fmt.Printf("Send the file size first: %d\n", info.Size())

	inputScanner := bufio.NewScanner(fInput)
	for inputScanner.Scan() {
		writer.WriteString(inputScanner.Text() + "\n")
	}
	writer.Flush()

	scanner := bufio.NewScanner(conn)
	if scanner.Scan() {
		fmt.Println("Server says: " + scanner.Text())
	}

}
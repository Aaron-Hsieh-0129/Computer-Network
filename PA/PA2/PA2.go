package main

import "fmt"
import "os"
import "bufio"
import "strconv"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	inputFile, outputFile := "", ""

	fmt.Printf("Input filename: ")
	fmt.Scanf("%s", &inputFile)

	fmt.Printf("Output filename: ")
	fmt.Scanf("%s", &outputFile)

	// open input file & create output file
	f, err := os.Open(inputFile)
	check(err)
	defer f.Close()

	f2, err2 := os.Create(outputFile)
	check(err2)
	defer f2.Close()

	// reader & writer
	scanner := bufio.NewScanner(f)
	writer := bufio.NewWriter(f2)

	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		num := strconv.Itoa(count)
		combine := num + " " + line

		writer.WriteString(combine)
		writer.WriteString("\n")
		count++
	}
	writer.Flush()
}
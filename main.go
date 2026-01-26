package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	
	line := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			fmt.Println(err) //io.EOF
			break
		}
		data = data[:n]
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			line += string(data[:i])
			data = data[i+1:]
			fmt.Printf("read: %s \n", line)
			line = ""
		}
		line += string(data)
	}
	if len(line) != 0 {
		fmt.Printf("read: %s\n", line)
	}
}

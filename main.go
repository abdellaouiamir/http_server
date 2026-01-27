package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

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
				out<-line
				line = ""
			}
			line += string(data)
		}
		if len(line) != 0 {
			out<-line
		}
	}()

	return out
}

func main() {
	listner, err := net.Listen("tcp", ":5555")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatal("error", "error", err)
		}
		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("read: %s\n", line)
		}
	}
}

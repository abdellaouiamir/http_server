package udpsender

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:5555")
	if err != nil {
		log.Fatal("error", "error", err)
	}
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("error", "error", err)
	}
	defer conn.Close()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
		}
	}
}

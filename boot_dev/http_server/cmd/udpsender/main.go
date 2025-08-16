package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatal("Not able to resolve udp address", err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal("Failed to dail udp", err)
	}

	fmt.Println("Local addr: ", conn.LocalAddr(), "Remote addr: ", conn.RemoteAddr())

	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		//Prompt
		fmt.Print("> ")

		// read user input
		line, err := reader.ReadString('\n') //input to fn is delimiter
		if err != nil {
			log.Printf("Not able to read stdin line, %v", err)
			continue
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Printf("Error writing to connection, %v", err)
			continue
		}
	}
}

// This is a dial udp, where we dial an udp connection. Oppposite to listen
// This way, we send data from our go program
// we can listen from `nc -u -l 42069`

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		str := ""
		for {
			data := make([]byte, 8)
			n, err := f.Read(data)
			if err != nil {
				break
			}

			data = data[:n]
			if i := bytes.IndexByte(data, '\n'); i != -1 {
				str += string(data[:i])
				data = data[i+1:]
				out <- str
				str = ""
			}

			str += string(data)
		}

		// when file doesnot end with a new line
		if len(str) != 0 {
			out <- str
		}
	}()

	return out
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Not able to create a listener", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Not able to create a listener", err)
		}
		for line := range getLinesChannel(conn) {
			fmt.Println(line)
		}
	}

}

// Notes

// https://www.youtube.com/watch?v=FknTw9bJsXM

// 1. f is a file pointer. When ever you do f.Read() it goes to next from prev position. Thats how it advances
// 2. second task is to still read data of 8 bytes but print every newline occurence. So we read until we find a newline using bytes.IndexByte() function. Print and clean
// 3. <-chan string : is a receive only channel of strings. You create a seperate fucntion with channels and goroutines and send each line to out channel
// 4. Now instead of reading from a file , we read from a tcp connection
// 5. To test this connection, run the program and in seperate pain, run `nc -v localhost 42069`. nc - netcat command to work with tcp connections
// 6. U can also send a curl request. For example: curl http://localhost:42069/coffee
//
// GET /coffee HTTP/1.1
// Host: localhost:42069
// User-Agent: curl/8.7.1
// Accept: */*
//
// 7. We get the plain request as output

// 8. curl -X POST -H "Content-Type: application/json" -d '{"flavor":"dark mode 2"}' http://localhost:42069/coffee
//
// POST /coffee HTTP/1.1
// Host: localhost:42069
// User-Agent: curl/8.7.1
// Accept: */*
// Content-Type: application/json
// Content-Length: 24

// {"flavor":"dark mode 2"}

// In POST, we get content-length field naem also
// In RFC, its called field-name. These are basically http headers

// Basic code

// start-line CRLF
// \*( field-line CRLF )
// CRLF
// [ message-body ]

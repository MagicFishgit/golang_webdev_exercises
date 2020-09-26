package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(c net.Conn) {
	defer c.Close()
	//Read from connection.
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" { //end of HTTP Request Headers
			break
		}
		fmt.Println(ln)
	}

	//Write to the connection
	io.WriteString(c, "Thanks for talking to me buddy!")
}

func main() {

	//I create a listener on port 8080 and specify to the net package that it should be tc.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error listening on port.", err)
	}
	defer listener.Close() //I need to remember to use defer to close things. Defer runs whatever you give it at the end of the funciton.

	//ive created the connection now I need to specify what it does with incomming connections.
	for {
		conn, err := listener.Accept() //accept the connection
		if err != nil {
			log.Println(err) //not using fatalln because I dont want the program to end if it got a failed connection.
			continue
		}

		//Serve the connection
		go serve(conn)

	}

}

package main

import(
	"fmt"
	"net"
	"bufio"
	"os"
)


func main(){
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil{
		fmt.Println("error connecting to server:", err)
		return
	}

	defer conn.Close()

	go receiveMessages(conn)

	for {
		message, _:= bufio.NewReader(os.Stdin).ReadString('\n')
		fmt.Fprintf(conn, message)
	}
}

func receiveMessages(conn net.Conn){
	for {
		message, _:= bufio.NewReader(conn).ReadString('\n')
		fmt.Print("message from server:", message)
	}
}
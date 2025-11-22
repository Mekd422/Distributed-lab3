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

	fmt.Print("enter message: ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	fmt.Fprintf(conn, text )

	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Print("response from server:", response)
}
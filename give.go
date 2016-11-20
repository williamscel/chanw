package main

import (
	"fmt"
	"net"
	"time"
)

/* conn, err := net.Dial("udp", "127.0.0.1:101")
if err != niil {
	fm.Println("Error:", err)
}*/

func main() {
	//connect to server
	conn, _ := net.Dial("udp", "127.0.0.1:8008")

	tim := time.Now()

	//convert to bytes and assigne to buf
	buf, _ := time.Now().MarshalBinary()

	//print out timestamp
	fmt.Printf("Packet sent at %s\n", tim)

	//write to socket
	conn.Write(buf)
	/*if buf != nil {
		fmt.Println("Server Message: Success Send")
	} else {
		fmt.Println("Server Message: Fail send")
	}*/

}

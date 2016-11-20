package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	add, _ := net.ResolveUDPAddr("udp", "8008")

	//listen on connections
	prt, _ := net.ListenUDP("udp", add)

	prt.Close()

	//accept connection on port
	//conn, _ := prt.Accept()

	//RUNS in a loop until cancelled
	for {
		var buf = make([]byte, 1024)
		n, _ := prt.Read(buf)

		t := new(time.Time)
		t.UnmarshalBinary(buf[:n])
		now := time.Now()

		tripTime := now.Sub(*t)

		fmt.Printf("Time Received is : %s\n", t)
		fmt.Printf("Traveling time: %s\n", tripTime)
		time.Sleep(time.Second * 1)

	}

	/*if buf != nil {
		fmt.Println("Client Message: Success Send")
	} else {
		fmt.Println("Client Message: Fail send")
	}*/

	//msg, _ := bufio.NewReader(conn).ReadString('\n')

}

//fmt.Print("Packet recieved at %s", buf)

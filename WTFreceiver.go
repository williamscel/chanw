package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"time"
)

type IntentionPacket struct {
	ClientId    string    // Identify the user sending the packets.
	PacketId    string    // Randomly generated string to id an intent packet.
	SentTime    time.Time // Time for latency analysis.
	ReceivedAt  time.Time //Store the time a packet is recieved at.
	PacketCount int       // Define how many packets to be generated.
	PayloadSize int       // Define the Size of the Packet Traffic Data.
	ListenTime  int       // integer Dictating the number of second the Receiver will listen for traffic packet.
}

type TrafficPacket struct {
	IntentPacketId string    // packetID of the intention packet associated with this traffic packet.
	SentTime       time.Time // Time for latency analysis.
	ReceivedAt     time.Time //Store the time a packet is recieved at.
	Data           []byte    // content of the packet.
}

func main() {
	//Listen for tcp connection
	listenForTCP()
}

//Method listening for tcp
func listenForTCP() {
	ConnTCP, err := net.Listen("tcp", ":3000")
	checkError(err)

	for {

		Conn, _ := ConnTCP.Accept()
		handleClientTCP(Conn)
		listenForUDP()
	}
}

//Method listening for tcp
func listenForUDP() {
	// listen for udp connection
	udpAddr, err := net.ResolveUDPAddr("udp", ":3000")
	checkError(err)
	for {
		ConnUDP, err := net.ListenUDP("udp", udpAddr)
		checkError(err)

		handleClientUDP(ConnUDP)
	}
}

//Deals with the traffic comming in through the socket
func handleClientUDP(conn *net.UDPConn) {
	defer conn.Close()
	decoder := gob.NewDecoder(conn)
	var p TrafficPacket

	//Ticker iterates over values incoming every 500 milliseconds for UPD packets
	//and stops receiving after 50000 milliseconds
	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {

		// Decode with reference to packet struct. If JSON format
		// and Packet struct format match accept the data.
		err := decoder.Decode(&p)
		checkError(err)
		fmt.Println(p)
	}
	time.Sleep(time.Millisecond * 50000)
	ticker.Stop()
}

//Deals with the traffic comming in through the socket
func handleClientTCP(conn net.Conn) {
	defer conn.Close()
	decoder := gob.NewDecoder(conn)
	var p IntentionPacket

	//Ticker iterates over values incoming every 500 milliseconds for TCP packets
	//and stops receiving after 50000 milliseconds
	ticker := time.NewTicker(time.Millisecond * 500)
	for range ticker.C {
		// Decode with reference to packet struct. If JSON format
		// and Packet struct format match accept the data.
		err := decoder.Decode(&p)
		checkError(err)
		fmt.Println(p)
	}
	time.Sleep(time.Millisecond * 50000)
	ticker.Stop()
}

// function to deal with error
func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal error ")
		os.Exit(1)
	}
}

func bytesToTime(b []byte) time.Time {
	var nsec int64
	for i := uint8(0); i < 8; i++ {
		nsec += int64(b[i]) << ((7 - i) * 8)
	}
	return time.Unix(nsec/1000000000, nsec%1000000000)
}

package main

import (
	"fmt"
	"os"
	"time"
	"net"
	"encoding/gob"
)

// intention packet structure
type IntentionPacket struct {
	ClientId string // Identify the user sending the packets.
	PacketId string  // Randomly generated string to id an intent packet.
	SentTime time.Time // Time for latency analysis.
	ReceivedAt time.Time //Store the time a packet is recieved at.
	ListenTime int // integer Dictating the number of second the Receiver will listen for traffic packet.
	PacketCount int // Define how many packets to be generated.
	PayloadSize int // Define the Size of the Packet Traffic Data.
}

// Traffic Packet  structure
type TrafficPacket struct {
	IntentPacketId string // packetID of the intention packet associated with this traffic packet.
	SentTime time.Time // Time for latency analysis.
	ReceivedAt time.Time //Store the time a packet is recieved at.
	Data  []byte // content of the packet.
}

func main() {

	udpAddr, err := net.ResolveUDPAddr("udp", ":2556")
	checkError(err)

	ConnUDP, err := net.ListenUDP("udp", udpAddr)
	checkError(err)

	time2live := handleClientTCP()

	ticker := time.NewTicker(time.Second * 1)
	go func(){
		for range ticker.C{
			handleClientUDP(ConnUDP)
		}
	}()
	time.Sleep(time.Duration(time2live) * time.Second)
	ticker.Stop()
	fmt.Println("ticker stopped")



	}

//Deals with the traffic comming in through the socket
func handleClientTCP() (int){
	Conn, err := net.Listen("tcp", ":2556")
	checkError(err)
	ConnTCP,_ := Conn.Accept()

 	fmt.Print("intension packet received", "\n")
  decoder := gob.NewDecoder(ConnTCP)

  var packet IntentionPacket

	decoder.Decode(&packet)

 	fmt.Print(packet, "\n","\n")

	return packet.ListenTime
}

//Deals with the traffic comming in through the socket
func handleClientUDP(conn *net.UDPConn) {

var packet TrafficPacket

fmt.Print("traffic packet received", "\n")
decoder := gob.NewDecoder(conn)

decoder.Decode(&packet)

fmt.Print(packet, "\n", "\n")

}

// function to deal with error
func checkError(e error) {
	if e != nil {
		fmt.Printf("Fatal error \n")
		os.Exit(1)
	}
}

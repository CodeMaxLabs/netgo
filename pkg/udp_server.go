package pkg


import (
	"fmt"
	"net"
)

func NewUDPServer(port int) {
	address, err := net.ResolveUDPAddr("udp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	conn, err := net.ListenUDP("udp", address)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP server listening on", address)

	buffer := make([]byte, 1024)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading:", err)
			return
		}

		fmt.Println("Received", string(buffer[:n]), "from", addr)

		_, err = conn.WriteToUDP([]byte("Message received"), addr)
		if err != nil {
			fmt.Println("Error sending:", err)
			return
		}
	}
}
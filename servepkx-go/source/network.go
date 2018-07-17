package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// getIP gets the the 3DS's IP address
func getIP() string {
	// startIP is the gateway, so if my ip address from getLen was 192.168.2.62
	// the startIP would be 192.168.2.1
	network := strings.Join(strings.Split(getLan(), ".")[0:3], ".")
	ip := scan(network)
	return ip
}

// scan scans the network for the 3DS's IP address
func scan(network string) string {
	dsIP := ""
	for i := 1; i < 254; i++ {
		// set the timeout for the connection checking
		timeout := time.Duration(100 * time.Millisecond)
		// check if port 8008 is open (if it is, then we ignore it
		// because that's a chromecast or a google home)
		// We do this because google devices also listen for port 9000
		// and this can cause the google device to be sent the pokemon files
		// instead of the 3ds
		c, err := net.DialTimeout("tcp", fmt.Sprintf("%s.%d:8008", network, i), timeout)

		if err == nil {
			c.Close()
			continue
		}
		// check for port 9000
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s.%d:9000", network, i), timeout)
		if err == nil {
			dsIP = fmt.Sprintf("%s.%d", network, i)
			conn.Close()
			break
		}
	}
	if dsIP == "" {
		fmt.Println("DS not found on network!")
		os.Exit(1)
	}
	fmt.Println("DS IP found! Please wait!")
	return dsIP
}

// getLan gets the LAN IP address
func getLan() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	// handle err...
	if err != nil {
		fmt.Println(err.Error())
	}

	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return fmt.Sprintf("%s", localAddr.IP)
}

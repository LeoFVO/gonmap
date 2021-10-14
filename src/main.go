package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	// get command line argument
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Using: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	target := os.Args[1] // target host

	activeRoutine := 0
	doneChannel := make(chan bool)

	for port := 0; port <= 65535; port++ {
		go testTCPConnection(target, port, doneChannel)
		activeRoutine++
	}
	// Wait for all Routine to finish
	for activeRoutine > 0 {
		<-doneChannel
		activeRoutine--
	}
}

func testTCPConnection(ip string, port int, doneChannel chan bool) {
	_, err := net.DialTimeout("tcp", ip+":"+strconv.Itoa(port),
		time.Second*10)
		// Uncomment next line to check all port scanned
		// fmt.Printf("Port %d: Closed\n", port)

	if err == nil {
		fmt.Printf("Port %d: Open\n", port)
	}
	doneChannel <- true
}

package main

import (
	"fmt"
	"net"
	"sync"
)

func ScanPort(port int, wg *sync.WaitGroup) {
	defer wg.Done()

	IP := "scanme.nmap.org"
	address := fmt.Sprintf(IP+":%d", port)
	connection, err := net.Dial("tcp", address)

	if err == nil {
		fmt.Printf("[+] Connection established to Port %v ... %v", port, connection.RemoteAddr().String())
		fmt.Print("\n")
		connection.Close()
	}
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go ScanPort(i, &wg)
	}

	wg.Wait()
}

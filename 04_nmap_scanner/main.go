package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Ullaakut/nmap"
)

func main() {
	target := "192.168.1.1/24"

	ctx, cancel_func := context.WithTimeout(context.Background(), 5*time.Minute)

	defer cancel_func()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets(target),
		nmap.WithPorts("22, 80, 443"),
		nmap.WithContext(ctx),
	)

	if err != nil {
		log.Fatal(err)
	}

	res, warning, err := scanner.Run()

	if err != nil {
		log.Fatal(err)
	}

	if warning != nil {
		log.Fatal(warning)
	}

	for _, host := range res.Hosts {
		if len(host.Ports) > 0 && len(host.Addresses) > 0 {
			fmt.Printf("IP: %q\n", host.Addresses[0])
			if len(host.Addresses) > 1 {
				fmt.Printf("MAC: %v\n", host.Addresses[1])
			}

			for _, port := range host.Ports {
				fmt.Printf("\tPort %d %s %s %s\n", port.ID, port.Protocol, port.State, port.Service)
			}

			fmt.Print("\n")
		}
	}

}

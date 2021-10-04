package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	iface      = "eth0"
	devFound   = false
	snaplength = int32(1600)
	timeout    = pcap.BlockForever
	filter     = "tcp and port 80"
	promisc    = false
)

// func find_devices() {
// 	devices, err := pcap.FindAllDevs()

// 	if err == nil {
// 		for index, dev := range devices {
// 			fmt.Println(index, ": ", dev.Name)

// 			for _, addr := range dev.Addresses {
// 				fmt.Println("\tIP:", addr.IP)
// 				fmt.Println("\tMASK:", addr.Netmask)
// 			}
// 		}
// 	}
// }

func capture_packets() {
	devices, err := pcap.FindAllDevs()

	if err == nil {
		for _, dev := range devices {
			if dev.Name == iface {
				devFound = true
			}

			if !devFound {
				log.Panicf("Device not found %s", iface)
			}

			handle, err := pcap.OpenLive(iface, snaplength, promisc, timeout)

			if err != nil {
				log.Fatalln(err)
			}

			defer handle.Close()

			err = handle.SetBPFFilter(filter)

			if err == nil {

				src := gopacket.NewPacketSource(handle, handle.LinkType())

				for pkt := range src.Packets() {
					applayer := pkt.ApplicationLayer()

					if applayer != nil {
						payload := applayer.Payload()
						search_arr := []string{"name", "username", "pass"}

						for _, term := range search_arr {
							index := strings.Index(string(payload), term)
							if index != -1 {
								fmt.Println(string(payload[index:]))
							}
						}
					}
				}
			}

		}
	}
}

func main() {
	capture_packets()
}

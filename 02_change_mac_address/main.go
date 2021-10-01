package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
)

func excuteCommand(command string, args_arr []string) (err error) {
	args := args_arr
	cmd_object := exec.Command(command, args...)

	cmd_object.Stderr = os.Stderr
	cmd_object.Stdout = os.Stdout
	cmd_object.Stdin = os.Stdin

	err = cmd_object.Run()

	if err != nil {
		log.Fatal(err)
		return
	}

	return nil
}

func main() {
	// ifconfig eth0 down
	// ifconfig eth0 hw ether 00:01:02:03:04:05
	// ifconfig eth0 up

	iface := flag.String("iface", "eth0", "Interface for which to change the MAC Address")
	newMAC := flag.String("newMAC", "", "The new MAC Address")

	flag.Parse()

	excuteCommand("sudo", []string{"ifconfig", *iface, "down"})
	excuteCommand("sudo", []string{"ifconfig", *iface, "hw", "ether", *newMAC})
	excuteCommand("sudo", []string{"ifconfig", *iface, "up"})
}

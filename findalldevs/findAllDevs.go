package findalldevs
// Package findalldevs provides a function to find all network devices
// and return the name of the device selected by the user.

import (

	"fmt"
	"github.com/google/gopacket/pcap"

)

func FindAllDevs() string {
devices, err := pcap.FindAllDevs()
	if err != nil {
		println("Error finding devices : ", err)
		// return "Error finding devices"
	}
	if len(devices) == 0 {
		println("No devices found")
		return ""
	}
	println("Available devices: ")
	for i, device := range devices {
		fmt.Printf("[%d]: %s | %d", i, device.Name, device.Flags)
		if device.Description != "" {
			fmt.Printf(" | %s", device.Description)
		}
		if len(device.Addresses) > 0 {
			fmt.Println(" Addresses: ")
			for _, address := range device.Addresses {
				fmt.Printf(" %s |", address.IP)
				fmt.Printf(" %s\n\n", address.Netmask)
			}

		}
	}
	// Take in the network interface name from the user using the number of the device

	fmt.Println("\nEnter the number of the device you want to sniff: ")
	var choice int
	_, err = fmt.Scanln(&choice)
	if err != nil {
		fmt.Println("Error reading choice: ", err)
		return ""
	}
	if choice < 0 || choice >= len(devices) {
		fmt.Println("Invalid choice")
		// make them pick a valid choice
		return ""
	}
	// print out the device name picked
	fmt.Printf("You picked device: %s\n", devices[choice].Name)
	return devices[choice].Name

}
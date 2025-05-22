package packetsniffer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"github.com/Taofik01/go-sdn-lab/findalldevs"
)

func PacketSniffer() string {
	// A go cli tool that can sniff a packet on a netweork interface
	// displaying the list of live network interfaces
	// and return the log including it's protocol, Source IP and port, Destination IP and port, payload length.
	// Output them to both terminal and a file.
	// Open a live packet capture on the specified network interface


	// get the list of all network devices
		
	
	deviceName := findalldevs.FindAllDevs()

	// open the device for live packet capture

	handle, err := pcap.OpenLive(deviceName, 262144, true, pcap.BlockForever)
    _ = handle.SetBPFFilter("tcp or udp")
	// log the packets to a file 
	// fmt.Printf("Opening device %s for live packet capture...\n", devices[choice].Name)
	
	
	if err != nil {
		fmt.Println("Error opening device: ", err)

	}
	packets := gopacket.NewPacketSource(handle, handle.LinkType())

	

	


	


	// write the packets to the log file
	fmt.Println("Do you want to save the packets to a file? (y/n)")
	var saveChoice string
	reader := bufio.NewReader(os.Stdin)
	saveChoice, _= reader.ReadString('\n')
	saveChoice = strings.TrimSpace(saveChoice)
	

	if saveChoice == "y" || saveChoice == "Y"{
	logFile, err := os.OpenFile("Packets_sniffed.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error while opening log file: ", err)
	}
	defer logFile.Close()
	writer := bufio.NewWriter(logFile)
	defer writer.Flush()



		writer.WriteString("Packet Sniffer Log\n")
		writer.WriteString("--------------------------------------------------\n")
		// iterate over the packets and print out the protocol, source IP and port, destination IP and port, payload length
		count := 0
		for packet := range packets.Packets() {
			count++
			if count > 200 {
				break
			}
			networklayer := packet.NetworkLayer()
			transportlayer := packet.TransportLayer()
			appLayer := packet.ApplicationLayer()
			if networklayer != nil && transportlayer != nil && appLayer != nil {
				writer.WriteString(fmt.Sprintf("Protocol: %s | Src: %s:%s -> Dst: %s:%s | Payload: %d bytes \n", 
					networklayer.LayerType(),
					networklayer.NetworkFlow().Src(),
					transportlayer.TransportFlow().Src(),
					networklayer.NetworkFlow().Dst(),
					transportlayer.TransportFlow().Dst(),
					len(appLayer.Payload()),
				))
				writer.WriteString("--------------------------------------------------\n")
			}
			
		}
	} else {
		count := 0
		for packet := range packets.Packets() {
		
		count++
		if count > 200 {
			break
		}

		networklayer := packet.NetworkLayer()
		transportlayer := packet.TransportLayer()
		appLayer := packet.ApplicationLayer()
		if networklayer != nil && transportlayer != nil && appLayer != nil {
			fmt.Printf("Protocol: %s | Src: %s:%s -> Dst: %s:%s | Payload: %d bytes \n", 
				networklayer.LayerType(),
				networklayer.NetworkFlow().Src(),
				transportlayer.TransportFlow().Src(),
				networklayer.NetworkFlow().Dst(),
				transportlayer.TransportFlow().Dst(),
				len(appLayer.Payload()),
			)	
			fmt.Println("--------------------------------------------------")
		}
		
	}
		
	
	}

	return "Live packet capture successfully on.. "
}

func FindAllDevs() {
	panic("unimplemented")
}

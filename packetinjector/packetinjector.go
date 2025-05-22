package packetinjector

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Taofik01/go-sdn-lab/findalldevs"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func PacketInjector() string {
	var protocol string
	var payload string
	// prompt the user for input 
	fmt.Println("Enter the destination IP address: ")
	reader := bufio.NewReader(os.Stdin)
	DstIP, _ := reader.ReadString('\n')
	DstIP = strings.TrimSpace(DstIP)  // remove newline character
	
	fmt.Println("Enter the destination port: ")
	portStr, _ := reader.ReadString('\n')
	portStr = strings.TrimSpace(portStr)  // remove newline character
	port, _ := strconv.Atoi(portStr)
	

	fmt.Println("Enter the protocol (tcp/udp): ")
	protocol, _ = reader.ReadString('\n')
	protocol = strings.TrimSpace(protocol)

	fmt.Println("Enter the payload lenght: ") 
	payload, _ = reader.ReadString('\n')
	payload = strings.TrimSpace(payload)  // remove newline character

	fmt.Println("Packet message to be injected: ")
	yrMessage, _ := reader.ReadString('\n')
	yrMessage = strings.TrimSpace(yrMessage)  // remove newline character

	println("Which interface do you want to inject the packet into? ")
	// get the list of all network devices
	deviceName := findalldevs.FindAllDevs()

	// create a packet with the given parameters
	message := fmt.Sprintf("internet face: %s | Protocol: %s | Dst: %s:%d | payload: %s | message: %s", deviceName, protocol, DstIP, port, payload, yrMessage)

	// create a packet with the given parameters
	gopacketMessage := gopacket.NewPacket([]byte(message), gopacket.LayerTypePayload, gopacket.Default)
	if gopacketMessage == nil {
		fmt.Println("Error creating packet")
		return "Error creating packet"
	}
	// send the packet to the destination IP and port
	handle, err := pcap.OpenLive(deviceName, 262144, true, pcap.BlockForever)
	if err != nil {
		return fmt.Sprintf("Error opening pcap handle: %v", err)
	}
	defer handle.Close()

	// construct the layer for the packet 
	ethLayer := &layers.Ethernet{
		SrcMAC:    getMacAddress(deviceName),
		DstMAC:    getMacAddress(DstIP),
		EthernetType: layers.EthernetTypeIPv4,

	}
	ipLayer := &layers.IPv4{
		SrcIP:    net.ParseIP("172.20.10.5"),
		DstIP:    net.ParseIP(DstIP),
		Protocol: layers.IPProtocolTCP,
	}
	

	// create the payload layer
	payloadLayer := gopacket.Payload([]byte(yrMessage)) // set the payload to the packet


	// serialize the layers
	buffer := gopacket.NewSerializeBuffer()
	options := gopacket.SerializeOptions{
		ComputeChecksums: true,
		FixLengths:       true,
	}
	if protocol == "udp" {
		udpLayer := &layers.UDP{
			SrcPort: layers.UDPPort(54321),
			DstPort: layers.UDPPort(port),
		}
		udpLayer.SetNetworkLayerForChecksum(ipLayer)
		gopacket.SerializeLayers(buffer, options, ethLayer, ipLayer, udpLayer, payloadLayer)
		
	} else if protocol == "tcp" {
		tcplayer := &layers.TCP{
			SrcPort: layers.TCPPort(54321),
			DstPort: layers.TCPPort(port),
			SYN:   true,
		}
		tcplayer.SetNetworkLayerForChecksum(ipLayer)
		gopacket.SerializeLayers(buffer, options, ethLayer, ipLayer, tcplayer, payloadLayer)
	} else {
		return "Invalid protocol"
	}
	// write the packet to the handle
	err = handle.WritePacketData(buffer.Bytes())
	if err != nil {
		return fmt.Sprintf("Error writing packet data: %v", err)
	}

	// capturing response after sending the packet
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		appLayer := packet.ApplicationLayer()
		if appLayer != nil {
			fmt.Printf("Received packet: %s\n", appLayer.Payload())
		} else {
			fmt.Println("No application layer found in the packet")
		}
		// print out the packet details to log
		networkLayer := packet.NetworkLayer()
		transportLayer := packet.TransportLayer()
		if networkLayer != nil && transportLayer != nil {
			logFile, err := os.OpenFile("packets_injected.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
			if err == nil {
				defer logFile.Close()
				writer := bufio.NewWriter(logFile)
				defer writer.Flush()
				payloadLen := 0
				if appLayer != nil {
					payloadLen = len(appLayer.Payload())
				}
				writer.WriteString(fmt.Sprintf("Protocol: %s | Src: %s:%s -> Dst: %s:%s | Payload: %d bytes \n",
					networkLayer.LayerType(),
					networkLayer.NetworkFlow().Src(),
					transportLayer.TransportFlow().Src(),
					networkLayer.NetworkFlow().Dst(),
					transportLayer.TransportFlow().Dst(),
					payloadLen,
				))
				writer.WriteString("--------------------------------------------------\n")
			} else {
				fmt.Println("Error opening log file:", err)
			}
		} else {
			fmt.Println("No network or transport layer found in the packet")
		}
		time.Sleep(1 * time.Second) // Avoid infnite loop if no packet is received
	}
	fmt.Printf("packet sent successfully to %s:%d\n", DstIP, port)

	return "Packet sent successfully"	
}



// getMacAddress returns the MAC address of the given network interface name or IP address.
func getMacAddress(deviceName string) net.HardwareAddr {
	iface, err := net.InterfaceByName(deviceName)
	if err == nil && iface != nil {
		return iface.HardwareAddr
	}
	addrs, err := iface.Addrs()
	if err != nil || len(addrs) == 0 {
		fmt.Println("No addresses found for the interface", err)
		return nil
	} 

	// Could not find MAC address, return nil
	return nil
}
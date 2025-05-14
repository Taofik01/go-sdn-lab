package packetanalyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Packet struct {
	SrcIP    string
	DstIP    string
	Protocol string
}

func PacketAnalyzer() string {
	var packets []Packet

	reader := bufio.NewReader(os.Stdin)
	stat, _ := os.Stdin.Stat()

	if stat.Mode()&os.ModeCharDevice == 1 {

		for i := 1; i <= 3; i++ {
			fmt.Printf("Enter your source IP for the %d packet: ", i)
			srcIP, _ := reader.ReadString('\n')
			srcIP = strings.TrimSpace(srcIP)

			fmt.Printf("Enter your Destination IP for the %d packet: ", i)
			dstIP, _ := reader.ReadString('\n')
			dstIP = strings.TrimSpace(dstIP)

			fmt.Printf("Enter your Protocol for the %d packet (tcp/udp): ", i)
			protocol, _ := reader.ReadString('\n')
			protocol = strings.TrimSpace(protocol)

			packets = append(packets, Packet{SrcIP: srcIP, DstIP: dstIP, Protocol: protocol})
		}
	} else {
		// read from a file passed during execution
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			// Each line contains SrcIp, DstIP, and Protocol separated by commas or space
			line := scanner.Text()
			fields := strings.FieldsFunc(line, func(r rune) bool {
				return r == ',' || r == ' '
			})
			if len(fields) != 3 {
				fmt.Println("Invalid line format: ", line)
				continue
			}

			// Trim spaces and create a packet
			srcIP := strings.TrimSpace(fields[0])
			dstIP := strings.TrimSpace(fields[1])
			protocol := strings.TrimSpace(fields[2])
			packets = append(packets, Packet{SrcIP: srcIP, DstIP: dstIP, Protocol: protocol})
		}

		if err := scanner.Err(); err != nil {
			fmt.Println("Error reading file: ", err)
			return "Error occured"
		}
	}

	// Analyze packets
	analyzePackets(packets)

	if len(packets) > 1 {
		fmt.Printf("%s", packets[1].SrcIP)
	}

	return "packet successfully analyzed"
}

func analyzePackets(packets []Packet) {
	// TO DO: implement packet analysis logic
	// Going through the packet and displaying the values with their frequency in a packet.txt file with log


	// Create a frequency map
	frequency := make(map[string]int)

	// Open a file to write the packet data
	logFile, err := os.OpenFile("packets.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error while opening log file: ", err)
		return
	}
	defer logFile.Close()

	// Count the frequency of each field value
	for _, packet := range packets {
		frequency[packet.SrcIP]++
		frequency[packet.DstIP]++
		frequency[packet.Protocol]++
	}

	// Create the tabular part for the file with the header
	header := fmt.Sprintf("%-20s %-20s %-20s\n", "Field", "Value", "Frequency")
	logFile.WriteString(header)
	logFile.WriteString(strings.Repeat("-", len(header)) + "\n")

	// Write only the unique packet data to the file
	for _, packet := range packets {
		// Write SrcIP
		row := fmt.Sprintf("%-20s %-20s %-20d\n", "SrcIP", packet.SrcIP, frequency[packet.SrcIP])
		logFile.WriteString(row)

		// Write DstIP
		row = fmt.Sprintf("%-20s %-20s %-20d\n", "DstIP", packet.DstIP, frequency[packet.DstIP])
		logFile.WriteString(row)

		// Write Protocol
		row = fmt.Sprintf("%-20s %-20s %-20d\n", "Protocol", packet.Protocol, frequency[packet.Protocol])
		logFile.WriteString(row)
	}

	fmt.Println("Packet data written to packets.txt file")
}

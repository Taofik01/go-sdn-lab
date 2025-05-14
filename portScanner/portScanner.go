package portscanner

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"strconv"
)

func PortScanner() string {
	var input string
	var rangeStart, rangeEnd int

	if len(os.Args) < 2 {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter an IP address: ...")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)

		fmt.Println("Enter the start of the port range: ...")
		rangeStartInput, _ := reader.ReadString('\n')
		rangeStartInput = strings.TrimSpace(rangeStartInput)
		rangeStart, _ = strconv.Atoi(rangeStartInput)
		
		

		fmt.Println("Enter the end of the port range: ...")
		rangeEndInput, _ := reader.ReadString('\n')
		rangeEndInput = strings.TrimSpace(rangeEndInput)
		rangeEnd, _ = strconv.Atoi(rangeEndInput)
		

	} else {
		input = os.Args[1]
		rangeStart, _ = strconv.Atoi(os.Args[2])
		rangeEnd, _ = strconv.Atoi(os.Args[3])
		
	}

	
	
	
	
	
	fmt.Printf("Scanning IP address %s for open ports in the range %d-%d...\n", input, rangeStart, rangeEnd)
	

	
	// Define the protocol to use for the connection
	
	fmt.Println("Enter the protocol to use (tcp/udp): ...")
	reader := bufio.NewReader(os.Stdin)
	protocol, _ := reader.ReadString('\n')
	protocol = strings.TrimSpace(protocol)
	if protocol != "tcp" && protocol != "udp" {
		fmt.Println("Invalid protocol. Please enter either 'tcp' or 'udp'.")
		return ""
	}
	
	 
	for port := rangeStart; port <= rangeEnd; port++ {
		address := fmt.Sprintf("%s:%v", input, port)
		conn, err := net.Dial(protocol, address)
		if err != nil {
			// fmt.Printf("port %s:%d is closed or unreachable\n", input, port)
			continue
		}

		fmt.Printf("The Port %s:%d is open\n", input, port)

		conn.Close()

		
	}
	

	
	return "Port Scanning completed successfully!"
	// close the connection

}

// This Go program scans a given IP address for open ports in a specified range.
// make use of the "net" and "fmt" packages to create a TCP connection to each port in the range.
// it should take Input: localhost, range: 1-1024 through os.Args also known as command line arguments
// It should return the open and closed ports in the specified range
// Display the ports that are open

package dnsresolver

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func DnsResolver() string {
	// This is a placeholder for a function that would resolve a DNS query.
	// Take a domain as input
	// Resolve and display
	// IPv4 addresses (A records)
	// IPv6 addresses (AAAA records)
	// Mail servers (MX records)
	// Name servers (NS records)
	// Print result to the terminal and optionally log them to dns_log.txt

	var input string
	reader := bufio.NewReader(os.Stdin)
	if len(os.Args) > 1 {
		input = os.Args[1]
	}
	
	// var ipv4, ipv6 []string

	// if user didnt add input from os Get user input
	if input == "" {
		
		fmt.Print("Enter a domain name to resolve: .... ")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
	}

	ips, err := net.LookupIP(input)
	if err != nil {
		fmt.Println("Error: ", err)

	}

	fmt.Printf("IPv4 addresses for %s\n", input)
	for _, ip := range ips {
		if ipv4 := ip.To4(); ipv4 != nil {
			fmt.Println(ipv4)
			// ipv4 = append(ipv4, ipv4.String()) // Save the IPv4 address in the slice
		}
	}

	fmt.Printf("IPv6 addresses for %s\n", input)
	for _, ip := range ips {
		if ipv6 := ip.To16(); ipv6 != nil {
			fmt.Println(ipv6)
			// ipv6 = append(ipv6, ipv6.String()) // Save the IPv6 address in the slice
		}
	}

	//To print mail servers (MX records)

	mxRecords, err := net.LookupMX(input)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Printf("Mail servers for %s:\n", input)
		for _, mx := range mxRecords {
			fmt.Printf("Host: %s, Priority: %d\n", mx.Host, mx.Pref)
			// ok displaing all values in the map
			for k, v := range mxRecords {
				fmt.Printf("%d: %v\n", k, v)
			}
		}
	}

	// To lookup name servers (NS records)

	nsRecord, err := net.LookupNS(input)
	if err != nil {
		fmt.Println("Error resolving name server: ", err)
	} else {
		fmt.Printf("Name Severs for %s\n: ", input)
		for k, v := range nsRecord {
			fmt.Printf("%d: %v\n", k, v)
		}
		
	}

	
	fmt.Print("Do you want the resolved information in a file: (Y/N)")
	
	var response string
	reader = bufio.NewReader(os.Stdin)
	response, _ = reader.ReadString('\n')
	if strings.TrimSpace(response) == "Y" || strings.TrimSpace(response) == "y" {
		// Open the file in write mode
		fileName := fmt.Sprintf("%s_dns.txt", input)
		file, err := os.Create(fileName)
		if err != nil {
			log.Fatal("There's an error resolving in the file contact administrator and try again later")
		} else {
			defer file.Close()
			// Write the resolved information to the file
			for _, ip := range ips {

				_, err = file.WriteString(fmt.Sprintf("IPv4 address for %s: %s\n", input, ip.To4().String()))
				if err != nil {
					log.Println("Error writing Ipv4 address to file", err)
				}

				_, err = file.WriteString(fmt.Sprintf("IPv6 address for %s: %s\n", input, ip.To16().String()))

				if err != nil {
					log.Println("Error writing IPv6 address to file", err)
				}

			}			

			for _, mx := range mxRecords {
				_, err = file.WriteString(fmt.Sprintf("MX records for "+input+" is: Host: %s , Priority : %d\n", mx.Host, mx.Pref))
				if err != nil {
				log.Println("Error writing MX records to file:", err)
			}

			}

			for _, ns := range nsRecord {
				_, err = file.WriteString(fmt.Sprintf("NS records for " + input + "is : Host: %s\n", ns.Host))
				if err != nil {
				log.Println("Error writing name servers to file:", err)
			}
			}
						

		}

		return fmt.Sprintf("Dns resolved information saved to %s", fileName)
	}

	return "Dns resolved information saved to console"

}

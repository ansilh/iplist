package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	//Get List of interfaces
	interfaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//Take one interface from interfaces list
	for _, iface := range interfaces {
		//Get interface name
		ifaceName, err := net.InterfaceByName(iface.Name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//Get IPs assigned to the interface name
		cidrs, _ := ifaceName.Addrs()
		//Loop through each IP/netmask
		for _, cidr := range cidrs {
			//Get IP and mask in separate variables - net.ParseCIDR will only accept string
			ip, ipNet, _ := net.ParseCIDR(cidr.String())
			//Convert Mask in bytes to decimal by casting mask to IP
			mask := net.IP(ipNet.Mask)
			//Print details
			fmt.Printf("%v: \n\tIP %v \n\tMask %v\n", iface.Name, ip, mask)
		}

	}
}

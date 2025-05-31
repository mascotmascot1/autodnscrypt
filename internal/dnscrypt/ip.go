// internal/dnscrypt/ip.go
package dnscrypt

import (
	"log"
	"net"
)

// GetIPv4FromInterface returns the first IPv4 address found on the network
// interface with the specified name. If an error occurs or no IPv4 address
// is found, it logs the error and exits the program.
func GetIPv4FromInterface(ifaceName string) string {
	interfaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}

	for _, iface := range interfaces {
		if iface.Name == ifaceName {
			addrs, err := iface.Addrs()
			if err != nil {
				log.Fatal(err)
			}
			for _, addr := range addrs {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				if ipv4 := ip.To4(); ipv4 != nil {
					return ipv4.String()
				}
			}
		}
	}
	return ""
}

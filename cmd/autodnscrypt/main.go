package main

import (
	"log"

	"github.com/mascotmascot1/autodnscrypt/internal/dnscrypt"
)

func main() {
	cfg := dnscrypt.LoadConfig()

	ip := dnscrypt.GetIPv4FromInterface(cfg.InterfaceName)
	if ip == "" {
		log.Fatalf("No IPv4 address found for interface: %s", cfg.InterfaceName)
	}

	dnscrypt.UpdateConfig(cfg.DNSCryptConfigPath, ip)
	dnscrypt.RunProxy(cfg.DNSCryptExePath)
	dnscrypt.SetDNS(cfg.InterfaceName, ip)
}

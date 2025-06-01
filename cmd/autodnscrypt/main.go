package main

import (
	"log"
	"os"

	"github.com/mascotmascot1/autodnscrypt/internal/dnscrypt"
)

func main() {
	const logFileName = "log.txt"
	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	cfg := dnscrypt.LoadConfig()

	ip := dnscrypt.GetIPv4FromInterface(cfg.InterfaceName)
	if ip == "" {
		log.Fatalf("No IPv4 address found for interface: %s", cfg.InterfaceName)
	}

	dnscrypt.UpdateConfig(cfg.DNSCryptConfigPath, ip)
	dnscrypt.RunProxy(cfg.DNSCryptExePath)
	dnscrypt.SetDNS(cfg.InterfaceName, ip)
}

// internal/dnscrypt/powershell.go
package dnscrypt

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// RunProxy starts the dnscrypt-proxy executable located at the given path.
// It inherits standard input, output, and error streams. Logs and exits on failure.
func RunProxy(exePath string) {
	cmd := exec.Command(exePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		log.Fatalf("Error running dnscrypt-proxy.exe: %v", err)
	}
	log.Println("dnscrypt-proxy.exe started, now configuring DNS via PowerShell...")
}

// SetDNS sets the DNS server address for the specified network interface
// using PowerShell. It runs the Set-DnsClientServerAddress command.
// Logs and exits on failure.
func SetDNS(ifaceName string, ip string) {
	psCmd := exec.Command("powershell", "-Command",
		fmt.Sprintf(`Set-DnsClientServerAddress -InterfaceAlias "%s" -ServerAddresses %s`, ifaceName, ip))
	psCmd.Stdout = os.Stdout
	psCmd.Stderr = os.Stderr
	psCmd.Stdin = os.Stdin

	if err := psCmd.Run(); err != nil {
		log.Fatalf("Failed to set DNS server via PowerShell: %v", err)
	}
	log.Printf(`Executed PowerShell command: Set-DnsClientServerAddress -InterfaceAlias "%s" -ServerAddresses %s`, ifaceName, ip)
}

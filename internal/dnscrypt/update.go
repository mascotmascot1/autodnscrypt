// internal/dnscrypt/update.go
package dnscrypt

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// UpdateConfig updates the dnscrypt-proxy configuration file at the given path
// by replacing the line starting with "listen_addresses" with one that contains
// the provided IP address. Logs and exits on failure.
func UpdateConfig(path string, ip string) {
	const listenKey = "listen_addresses" // dnscrypt-proxy config key

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(strings.TrimSpace(line), listenKey) {
			newLine := fmt.Sprintf("%s = ['%s:53']", listenKey, ip)
			lines = append(lines, newLine)
		} else {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Configuration file successfully updated.")
}

// Package dnscrypt provides utilities to configure and run dnscrypt-proxy,
// including config file updates, interface IP detection, and DNS setup via PowerShell.
package dnscrypt

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	DNSCryptConfigPath string `yaml:"dnscryptConfigPath"`
	DNSCryptExePath    string `yaml:"dnscryptExePath"`
	InterfaceName      string `yaml:"interfaceName"`
}

func LoadConfig() Config {
	const configFileName = "config.yaml"

	data, err := os.ReadFile(configFileName)
	if err != nil {
		log.Fatalf("Failed to read %s: %v", configFileName, err)
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		log.Fatalf("Failed to parse %s: %v", configFileName, err)
	}

	return cfg
}

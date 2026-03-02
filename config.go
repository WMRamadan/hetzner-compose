package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Network  Network  `yaml:"network"`
	Firewall Firewall `yaml:"firewall"`
	Servers  []Server `yaml:"servers"`
}

type NetworkSubnetConfig struct {
	IPRange string `yaml:"ip_range"`
	Zone    string `yaml:"zone"`
}

type Network struct {
	Name    string              `yaml:"name"`
	IPRange string              `yaml:"ip_range"`
	Subnet  NetworkSubnetConfig `yaml:"subnet"`
}

type Firewall struct {
	Name  string         `yaml:"name"`
	Rules []FirewallRule `yaml:"rules"`
}

type FirewallRule struct {
	Protocol  string   `yaml:"protocol"`
	Port      string   `yaml:"port"`
	SourceIPs []string `yaml:"source_ips"`
}

type Server struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Image    string   `yaml:"image"`
	Location string   `yaml:"location"`
	SSHKeys  []string `yaml:"ssh_keys"`
}

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

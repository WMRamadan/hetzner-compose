package hetzner

import (
	"context"
	"fmt"
	"hetzner-compose/config"
	"net"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func parseIPNet(cidr string) (*net.IPNet, error) {
	_, ipnet, err := net.ParseCIDR(cidr)
	return ipnet, err
}

func parseIPNets(cidrs []string) ([]net.IPNet, error) {
	var result []net.IPNet
	for _, c := range cidrs {
		_, ipnet, err := net.ParseCIDR(c)
		if err != nil {
			return nil, err
		}
		result = append(result, *ipnet)
	}
	return result, nil
}

func Up(client *hcloud.Client, cfg *config.Config) error {
	ctx := context.Background()

	labels := map[string]string{
		"managed-by": "hetzner-compose",
	}

	// -----------------------
	// Create Network
	// -----------------------
	fmt.Println("Creating network...")

	ipRange, err := parseIPNet(cfg.Network.IPRange)
	if err != nil {
		return err
	}

	netResult, _, err := client.Network.Create(ctx, hcloud.NetworkCreateOpts{
		Name:    cfg.Network.Name,
		IPRange: ipRange,
		Labels:  labels,
	})
	if err != nil {
		return err
	}

	// Add subnet (required for cloud servers)
	subnetRange, err := parseIPNet(cfg.Network.Subnet.IPRange)
	if err != nil {
		return err
	}

	zone := hcloud.NetworkZone(cfg.Network.Subnet.Zone)

	_, _, err = client.Network.AddSubnet(ctx, netResult, hcloud.NetworkAddSubnetOpts{
		Subnet: hcloud.NetworkSubnet{
			Type:        hcloud.NetworkSubnetTypeCloud,
			NetworkZone: zone,
			IPRange:     subnetRange,
		},
	})
	if err != nil {
		return err
	}

	// -----------------------
	// Create Firewall
	// -----------------------
	fmt.Println("Creating firewall...")

	var rules []hcloud.FirewallRule

	for _, r := range cfg.Firewall.Rules {
		port := r.Port

		sourceIPs, err := parseIPNets(r.SourceIPs)
		if err != nil {
			return err
		}

		rules = append(rules, hcloud.FirewallRule{
			Direction: hcloud.FirewallRuleDirectionIn,
			Protocol:  hcloud.FirewallRuleProtocol(r.Protocol),
			Port:      &port,
			SourceIPs: sourceIPs,
		})
	}

	fwResult, _, err := client.Firewall.Create(ctx, hcloud.FirewallCreateOpts{
		Name:   cfg.Firewall.Name,
		Rules:  rules,
		Labels: labels,
	})
	if err != nil {
		return err
	}

	// -----------------------
	// Create Servers
	// -----------------------
	for _, s := range cfg.Servers {
		fmt.Printf("Creating server %s...\n", s.Name)

		serverType, _, err := client.ServerType.GetByName(ctx, s.Type)
		if err != nil {
			return err
		}
		if serverType == nil {
			return fmt.Errorf("server type not found: %s", s.Type)
		}

		image, _, err := client.Image.GetByName(ctx, s.Image)
		if err != nil {
			return err
		}

		location, _, err := client.Location.GetByName(ctx, s.Location)
		if err != nil {
			return err
		}

		sshKeys, err := LoadOrCreateSSHKeys(ctx, client, s.SSHKeys)
		if err != nil {
			return err
		}

		var sshPointers []*hcloud.SSHKey
		for _, k := range sshKeys {
			sshPointers = append(sshPointers, k)
		}

		createResult, _, err := client.Server.Create(ctx, hcloud.ServerCreateOpts{
			Name:       s.Name,
			ServerType: serverType,
			Image:      image,
			Location:   location,
			SSHKeys:    sshPointers,
			Networks: []*hcloud.Network{
				netResult,
			},
			Firewalls: []*hcloud.ServerCreateFirewall{
				{
					Firewall: *fwResult.Firewall,
				},
			},
			Labels: labels,
		})
		if err != nil {
			return err
		}

		if createResult.Action != nil {
			if err := client.Action.WaitFor(ctx, createResult.Action); err != nil {
				return err
			}
		}

		if createResult.Server != nil {
			fmt.Println("Server IP:", createResult.Server.PublicNet.IPv4.IP.String())
		}
	}

	fmt.Println("Infrastructure created.")
	return nil
}

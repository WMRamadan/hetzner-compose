package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func Down(client *hcloud.Client, cfg *Config) error {
	ctx := context.Background()

	fmt.Println("Deleting servers...")
	servers, _, err := client.Server.List(ctx, hcloud.ServerListOpts{
		ListOpts: hcloud.ListOpts{
			LabelSelector: "managed-by=hetzner-compose",
		},
	})
	if err != nil {
		return err
	}

	for _, s := range servers {
		fmt.Printf("Deleting %s...\n", s.Name)
		_, err := client.Server.Delete(ctx, s)
		if err != nil {
			return err
		}
		if err := waitForServerDeletion(ctx, client, 120*time.Second); err != nil {
			return err
		}
	}

	fmt.Println("Deleting firewalls...")
	firewalls, _, _ := client.Firewall.List(ctx, hcloud.FirewallListOpts{
		ListOpts: hcloud.ListOpts{
			LabelSelector: "managed-by=hetzner-compose",
		},
	})
	if err != nil {
		return err
	}

	for _, f := range firewalls {
		fmt.Println("Deleting firewall", f.Name)

		for {
			_, err := client.Firewall.Delete(ctx, f)
			if err == nil {
				break
			}

			// Retry if resource is still attached
			time.Sleep(2 * time.Second)
		}
	}

	fmt.Println("Deleting networks...")
	networks, _, _ := client.Network.List(ctx, hcloud.NetworkListOpts{
		ListOpts: hcloud.ListOpts{
			LabelSelector: "managed-by=hetzner-compose",
		},
	})
	for _, n := range networks {
		client.Network.Delete(ctx, n)
	}

	fmt.Println("Deleting SSH keys...")

	sshKeys, _, err := client.SSHKey.List(ctx, hcloud.SSHKeyListOpts{
		ListOpts: hcloud.ListOpts{
			LabelSelector: "managed-by=hetzner-compose",
		},
	})
	if err != nil {
		return err
	}

	for _, k := range sshKeys {
		fmt.Println("Deleting SSH key", k.Name)

		_, err := client.SSHKey.Delete(ctx, k)
		if err != nil {
			return err
		}

		// Short grace wait for backend propagation
		time.Sleep(2 * time.Second)
	}

	fmt.Println("Infrastructure deleted.")
	return nil
}

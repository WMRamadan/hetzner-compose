package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func waitForServerDeletion(ctx context.Context, client *hcloud.Client, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)

	for {
		if time.Now().After(deadline) {
			return fmt.Errorf("timeout waiting for servers to be deleted")
		}

		servers, _, err := client.Server.List(ctx, hcloud.ServerListOpts{
			ListOpts: hcloud.ListOpts{
				LabelSelector: "managed-by=hetzner-compose",
			},
		})
		if err != nil {
			return err
		}

		if len(servers) == 0 {
			return nil
		}

		time.Sleep(3 * time.Second)
	}
}

package hetzner

import (
	"os"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func NewClient() *hcloud.Client {
	token := os.Getenv("HCLOUD_TOKEN")
	if token == "" {
		panic("HCLOUD_TOKEN not set")
	}

	return hcloud.NewClient(
		hcloud.WithToken(token),
	)
}

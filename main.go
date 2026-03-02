package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: hetzner-compose [up|down]")
	}

	token := os.Getenv("HCLOUD_TOKEN")
	if token == "" {
		log.Fatal("HCLOUD_TOKEN not set")
	}

	client := hcloud.NewClient(hcloud.WithToken(token))

	cfg, err := LoadConfig("hetzner-compose.yml")
	if err != nil {
		log.Fatal(err)
	}

	switch os.Args[1] {
	case "up":
		if err := Up(client, cfg); err != nil {
			log.Fatal(err)
		}
	case "down":
		if err := Down(client, cfg); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("unknown command")
	}
}

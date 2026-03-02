package main

import (
	"log"
	"os"

	"hetzner-compose/config"
	"hetzner-compose/provider/hetzner"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: hetzner-compose up|down")
	}

	cfg, err := config.Load("hetzner-compose.yml")
	if err != nil {
		log.Fatal(err)
	}

	client := hetzner.NewClient()

	switch os.Args[1] {
	case "up":
		if err := hetzner.Up(client, cfg); err != nil {
			log.Fatal(err)
		}

	case "down":
		if err := hetzner.Down(client, cfg); err != nil {
			log.Fatal(err)
		}

	default:
		log.Fatal("unknown command")
	}
}

# Hetzner Compose

A Docker Compose like tool for Hetzner Cloud which uses a YAML file to build your infrastructure.

## Requirements
You will need an `hcloud token` once you have your token export it:
```
export HCLOUD_TOKEN=your_token_here
```

## Build
To build the cli:
```
go build
```

## Execution
You can use the tool like you do docker compose, make sure that your `hetzner-compose.yml` file is in the same directory as the cli.

## Commands
```
./hetzner-compose up
./hetzner-compose down
```
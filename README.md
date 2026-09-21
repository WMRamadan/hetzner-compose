# Hetzner Compose

### Docker Compose for Hetzner Cloud Infrastructre.

Define your Hetzner Cloud infrastructure in a simple YAML file and create or destroy it with a single command.

## Requirements
You will need an `hcloud token` once you have your token export it:
```
export HCLOUD_TOKEN=your_token_here
```

## Build
To build the cli:
```
go build ./cmd/hetzner-compose
```

## Execution
You can use the tool like you do docker compose, make sure that your `hetzner-compose.yml` file is in the same directory as the cli.

## Commands
```
./hetzner-compose up
./hetzner-compose down
```

## Why Hetzner Compose?

Managing cloud infrastructure can become unnecessarily complicated for small projects.

Terraform is powerful, but sometimes you just want to describe a few servers in a YAML file and create them.

Hetzner Compose provides a simple, Docker Compose like interface for Hetzner Cloud infrastructure.

## Features

- Declarative Hetzner Cloud Infrastructure
- YAML based configuration
- Create & Destroy Infrastructure
- Lightweight single binary

## Use Cases

Hetzner Compose is useful for developers and infrastructure engineers who want a simple way to create repeatable Hetzner Cloud environments.

- Small production environments
- Compose a fleet of servers with one command
- Can be used in CI Deployment pipelines

## Hetzner Compose vs Terraform

Hetzner Compose is not intended to replace Terraform for every use case.

Terraform provides a much larger infrastructure-as-code ecosystem and is better suited to complex multi-provider environments.

Hetzner Compose focuses on a simpler problem:

|     | Hetzner Compose | Terraform |
| --- | --------------- | --------- |
| Configuration | YAML | HCL |
| Primary focus | Hetzner Cloud | Multi-provider infrastructure |
| Workflow | Declarative YAML file | IaC with state |
| Learning curve | Low | High |
| Multi-cloud | No | Yes |
| Best suited for | Small Hetzner environments | Complex infrastructure |

## How It Works
                hetzner-compose.yml
                         │
                         ▼
                 hetzner-compose
                         │
                         ▼
                 Hetzner Cloud API
                         │
             ┌───────────┼───────────┐
             ▼           ▼           ▼
          Servers      Networks    Volumes

The CLI reads the YAML configuration and uses the Hetzner Cloud API to create and manage the resources described by it.

## Roadmap

Planned improvements may include:

[] More Hetzner Cloud resources

[] Resource dependencies

[] Update/reconciliation support

[] Better infrastructure state handling

[] Configuration validation

[] Multiple configuration files

[] Variable interpolation

[] Dry-run support

[] Import existing Hetzner resources

[] Kubernetes cluster provisioning

[] Additional networking features

Have an idea? Open an issue or start a discussion.

## Contributing

Contributions are welcome.

If you have an idea for a feature, improvement, or bug fix, please open an issue before submitting a large change.

Pull requests are welcome.

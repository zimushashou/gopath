# Ihomeweb Service

This is the Ihomeweb service

Generated with

```
micro new ihome/ihomeweb --namespace=go.micro --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.web.ihomeweb
- Type: web
- Alias: ihomeweb

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./ihomeweb-web
```

Build a docker image
```
make docker
```
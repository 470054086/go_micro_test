# Hello Service

This is the Hello service

Generated with

```
micro new hello --namespace=com.foo --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: com.foo.srv.hello
- Type: srv
- Alias: hello

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
./hello-srv
```

Build a docker image
```
make docker
```
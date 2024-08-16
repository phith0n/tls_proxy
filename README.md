# TLS Traffic Forwarder

A lightweight reverse proxy server that converts TLS traffic to TCP, allowing secure communication between clients and upstream servers.

## Features

- Accepts TLS connections from clients
- Forwards decrypted traffic to upstream TCP servers
- Easy to configure and use
- Supports custom TLS certificates

## Usage

```shell
./tproxy -l [local_address]:[port] -r [upstream_address]:[port] -c [cert_file] -k [key_file]
```

### Options

- `-l`: Local address and port to listen on (e.g., 127.0.0.1:8443)
- `-r`: Upstream server address and port to forward traffic to (e.g., 127.0.0.1:8080)
- `-c`: Path to the TLS certificate file (e.g., cert.pem)
- `-k`: Path to the TLS private key file (e.g., key.pem)

### Example

```shell
./main -l 127.0.0.1:8443 -r 127.0.0.1:8080 -c cert.pem -k key.pem
```

This command starts the proxy server, listening on 127.0.0.1:8443 for TLS connections, and forwarding decrypted traffic to 127.0.0.1:8080 using the specified certificate and key files.

## Building from Source

```shell
go build -o tproxy tproxy.go
```

## License

MIT © [phith0n](https://www.leavesongs.com)

# HTTP/1.1 vs HTTP/2 Diagnostic Server + Client in Go #

## Goal: ##

Create a Go-based diagnostic tool that includes:
### 1. An HTTP server supporting both HTTP/1.1 and HTTP/2 ###
### 2. A custom client that: ###

    - Sends custom headers
    - Benchmarks header size
    - Tests multiple concurrent requests
    - Shows whether multiplexing is happening

### 3. Logs detailed request info: ###

    - Protocol version
    - Header sizes
    - Arrival times
    - Client IP

## Features: ##

### Server (server/main.go) ###

- TLS-enabled Go HTTP server
- Supports both HTTP/1.1 and HTTP/2 automatically
- Logs:
    - Client IP
    - Header key-values
    - Protocol used (HTTP/1.1 or HTTP/2)
    - Time of request

### Client (client/main.go) ###
- Sends custom headers
- Configurable to force HTTP/1.1 or HTTP/2
- Measures response time
- Sends multiple concurrent requests to test multiplexing
- Prints protocol used in response


## Project Structure: ##

```bash
GohttpDiagnostics/
│
├── server/
│   └── main.go        # HTTP/1.1 + HTTP/2 enabled server
│
├── client/
│   └── main.go        # Custom diagnostic client
│
├── go.mod
│
├── README.md
```
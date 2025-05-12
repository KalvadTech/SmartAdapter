# SmartAdapter

A simple HTTP proxy adapter written in Go. Forwards requests to a target domain, preserving query parameters, POST bodies, and headers.

## Download and Run (Binary)

Download the latest release for your platform from the [GitHub Releases page](https://github.com/KalvadTech/SmartAdapter/releases) and make it executable:

```
chmod +x smartadapter
./smartadapter -target https://sandboxapi.customerpulse.gov.ae/ -port 8080 -H "X-Integration-Apikey: your-api-key" -H "Auth-GSB-example: another-value"
```

You can add any number of custom headers to be forwarded to the target by repeating the `-H` flag:

```
./smartadapter -target https://example.com/ -H "Authorization: Bearer token" -H "Auth-GSB-example: value"
```

## Run with Docker

You can use the Docker image directly from the GitHub repository, or build it yourself.

### Dynamic Port Exposure

The Dockerfile exposes the port dynamically based on the `PORT` environment variable (default: 8080). To run the container on a custom port, set the `PORT` environment variable and map the same port:

```
docker run -e PORT=9090 -p 9090:9090 ghcr.io/kalvadtech/smartadapter:latest -target https://sandboxapi.customerpulse.gov.ae/ -port 9090 -H "X-Integration-Apikey: your-api-key"
```

If you do not set `PORT`, it defaults to 8080:


## Pull from GitHub Container Registry :

```
docker pull ghcr.io/kalvadtech/smartadapter:latest
docker run -p 8080:8080 ghcr.io/kalvadtech/smartadapter:latest -target https://sandboxapi.customerpulse.gov.ae/ -port 8080 -H "X-Integration-Apikey: your-api-key" -H "X-Another-Header: another-value"
```

## Example docker-compose.yml

```
version: '3.8'
services:
  smartadapter:
    build: .
    ports:
      - "8080:8080"
    command: ["-target", "https://sandboxapi.customerpulse.gov.ae/", "-port", "8080", "-H", "X-Integration-Apikey: your-api-key", "-H", "X-Another-Header: another-value"]
```

## Arguments
- `-target` (required): The target domain to forward requests to (e.g., https://sandboxapi.customerpulse.gov.ae/)
- `-port`: Port to listen on (default: 8080)
- `-H`: Custom header to add (repeatable, format: 'Key: Value'). **You can specify as many custom headers as you need by repeating this flag.**

## License
MIT

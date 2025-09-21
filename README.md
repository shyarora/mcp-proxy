# MCP Proxy Server

An MCP proxy that aggregates multiple MCP servers behind a single HTTP entrypoint.

**This fork has been enhanced to work both as a standalone CLI application and as a Go library for integration into other projects.**

## Features

- Proxy multiple MCP clients: aggregate tools, prompts, and resources from many servers.
- SSE and streamable HTTP: serve via Server‑Sent Events or streamable HTTP.
- Flexible config: supports `stdio`, `sse`, and `streamable-http` client types.
- **Library Support**: Can be imported and used as a Go library in other projects.

## Documentation

- Configuration: [docs/configuration.md](docs/CONFIGURATION.md)
- Usage: [docs/usage.md](docs/USAGE.md)
- Deployment: [docs/deployment.md](docs/DEPLOYMENT.md)
- **Library Usage**: [LIBRARY_USAGE.md](LIBRARY_USAGE.md)
- **Integration Guide**: [INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md)
- Claude config converter: https://tbxark.github.io/mcp-proxy

## Quick Start

### As a Standalone CLI

#### Build from source

```bash
git clone https://github.com/shyarora/mcp-proxy.git
cd mcp-proxy
go build .
./mcp-proxy --config path/to/config.json
```

#### Install via Go

```bash
go install github.com/shyarora/mcp-proxy@latest
mcp-proxy --config path/to/config.json
```

### As a Go Library

```bash
go get github.com/shyarora/mcp-proxy
```

```go
import mcpproxy "github.com/shyarora/mcp-proxy"

opts := mcpproxy.ProxyOptions{
    ConfigPath:  "config.json",
    ExpandEnv:   true,
    HTTPTimeout: 10,
}
err := mcpproxy.RunProxy(opts)
```

See [LIBRARY_USAGE.md](LIBRARY_USAGE.md) for detailed library documentation and [INTEGRATION_GUIDE.md](INTEGRATION_GUIDE.md) for Cobra CLI integration.

### Docker

The image includes support for launching MCP servers via `npx` and `uvx`.

```bash
docker run -d -p 9090:9090 -v /path/to/config.json:/config/config.json ghcr.io/tbxark/mcp-proxy:latest
# or provide a remote config
docker run -d -p 9090:9090 ghcr.io/tbxark/mcp-proxy:latest --config https://example.com/config.json
```

More deployment options (including docker‑compose) are in [docs/deployment.md](docs/DEPLOYMENT.md).

## Configuration

See full configuration reference and examples in [docs/configuration.md](docs/CONFIGURATION.md).
An online Claude config converter is available at: https://tbxark.github.io/mcp-proxy

## Usage

Command‑line flags, endpoints, and auth examples are documented in [docs/usage.md](docs/USAGE.md).

## Thanks

- This project was inspired by the [adamwattis/mcp-proxy-server](https://github.com/adamwattis/mcp-proxy-server) project
- If you have any questions about deployment, you can refer to [《在 Docker 沙箱中运行 MCP Server》](https://miantiao.me/posts/guide-to-running-mcp-server-in-a-sandbox/)([@ccbikai](https://github.com/ccbikai))

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

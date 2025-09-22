# MCP Proxy Server

An MCP proxy that aggregates multiple MCP servers behind a single HTTP entrypoint.

**This is a Go library for integration into other projects.**

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
- Claude config converter: https://tbxark.github.io/mcp-proxy

## Quick Start

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

See [LIBRARY_USAGE.md](LIBRARY_USAGE.md) for detailed library documentation.

## Configuration

See full configuration reference and examples in [docs/configuration.md](docs/CONFIGURATION.md).
An online Claude config converter is available at: https://tbxark.github.io/mcp-proxy

## Usage

Library usage examples and API documentation are documented in [LIBRARY_USAGE.md](LIBRARY_USAGE.md).

## Thanks

- This project was inspired by the [adamwattis/mcp-proxy-server](https://github.com/adamwattis/mcp-proxy-server) project
- If you have any questions about deployment, you can refer to [《在 Docker 沙箱中运行 MCP Server》](https://miantiao.me/posts/guide-to-running-mcp-server-in-a-sandbox/)([@ccbikai](https://github.com/ccbikai))

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

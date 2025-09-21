# Using mcp-proxy as a Library

This project can be used both as a standalone CLI application and as a Go library in your own projects.

## As a CLI Application

```bash
# Build and run
go build .
./mcp-proxy --config config.json

# Or install and run
go install github.com/shyarora/mcp-proxy@latest
mcp-proxy --config config.json
```

## As a Library

### Installation

```bash
go get github.com/shyarora/mcp-proxy
```

### Basic Usage

```go
package main

import (
    "log"
    mcpproxy "github.com/shyarora/mcp-proxy"
)

func main() {
    // Option 1: Use the convenience function
    opts := mcpproxy.ProxyOptions{
        ConfigPath:  "config.json",
        Insecure:    false,
        ExpandEnv:   true,
        HTTPHeaders: "",
        HTTPTimeout: 10,
    }

    if err := mcpproxy.RunProxy(opts); err != nil {
        log.Fatal(err)
    }
}
```

### Advanced Usage

```go
package main

import (
    "log"
    mcpproxy "github.com/shyarora/mcp-proxy"
)

func main() {
    // Load configuration
    config, err := mcpproxy.LoadConfig("config.json", false, true, "", 10)
    if err != nil {
        log.Fatalf("Failed to load config: %v", err)
    }

    // Start server
    if err := mcpproxy.StartServer(config); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
```

### Integration with Cobra CLI

```go
package main

import (
    "github.com/spf13/cobra"
    mcpproxy "github.com/shyarora/mcp-proxy"
)

func main() {
    var rootCmd = &cobra.Command{Use: "myapp"}

    var mcpProxyCmd = &cobra.Command{
        Use:   "mcp-proxy",
        Short: "Start MCP proxy server",
        RunE: func(cmd *cobra.Command, args []string) error {
            configPath, _ := cmd.Flags().GetString("config")
            insecure, _ := cmd.Flags().GetBool("insecure")
            expandEnv, _ := cmd.Flags().GetBool("expand-env")
            httpHeaders, _ := cmd.Flags().GetString("http-headers")
            httpTimeout, _ := cmd.Flags().GetInt("http-timeout")

            opts := mcpproxy.ProxyOptions{
                ConfigPath:  configPath,
                Insecure:    insecure,
                ExpandEnv:   expandEnv,
                HTTPHeaders: httpHeaders,
                HTTPTimeout: httpTimeout,
            }

            return mcpproxy.RunProxy(opts)
        },
    }

    // Add the same flags as the original mcp-proxy
    mcpProxyCmd.Flags().StringP("config", "c", "config.json", "path to config file or a http(s) url")
    mcpProxyCmd.Flags().Bool("insecure", false, "allow insecure HTTPS connections")
    mcpProxyCmd.Flags().Bool("expand-env", true, "expand environment variables in config file")
    mcpProxyCmd.Flags().String("http-headers", "", "optional HTTP headers for config URL")
    mcpProxyCmd.Flags().Int("http-timeout", 10, "HTTP timeout in seconds")

    rootCmd.AddCommand(mcpProxyCmd)
    rootCmd.Execute()
}
```

## API Reference

### Types

- `ProxyConfig`: The main configuration struct
- `ProxyOptions`: Options for running the proxy (equivalent to CLI flags)

### Functions

- `LoadConfig(configPath, insecure, expandEnv, httpHeaders, httpTimeout)`: Load configuration from file or URL
- `StartServer(config)`: Start the HTTP server with the given configuration
- `StartServerWithContext(ctx, config)`: Start server with context for lifecycle control
- `RunProxy(opts)`: Convenience function that combines LoadConfig and StartServer
- `NewDefaultProxyOptions()`: Create ProxyOptions with default values

### Default Values

- ConfigPath: `"config.json"`
- Insecure: `false`
- ExpandEnv: `true`
- HTTPHeaders: `""`
- HTTPTimeout: `10`

## Example for OneDevX CLI

```bash
# In your onedevx project
go get github.com/shyarora/mcp-proxy

# Then use it as:
onedevx apps mcp-proxy --config config.json --insecure=false
```

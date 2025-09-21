package mcpproxy

import (
	"context"
)

// ProxyConfig represents the main configuration for the MCP proxy
type ProxyConfig struct {
	McpProxy   *MCPProxyConfigV2             `json:"mcpProxy"`
	McpServers map[string]*MCPClientConfigV2 `json:"mcpServers"`
}

// LoadConfig loads the configuration from a file or URL with the specified options.
// This is the public API function that external libraries can use.
//
// Parameters:
//   - configPath: path to config file or HTTP(S) URL
//   - insecure: allow insecure HTTPS connections by skipping TLS certificate verification
//   - expandEnv: expand environment variables in config file
//   - httpHeaders: optional HTTP headers for config URL, format: 'Key1:Value1;Key2:Value2'
//   - httpTimeout: HTTP timeout in seconds when fetching config from URL
//
// Returns the loaded configuration or an error.
func LoadConfig(configPath string, insecure, expandEnv bool, httpHeaders string, httpTimeout int) (*ProxyConfig, error) {
	config, err := load(configPath, insecure, expandEnv, httpHeaders, httpTimeout)
	if err != nil {
		return nil, err
	}

	return &ProxyConfig{
		McpProxy:   config.McpProxy,
		McpServers: config.McpServers,
	}, nil
}

// StartServer starts the HTTP server with the provided configuration.
// This function blocks until the server is shut down or an error occurs.
//
// The server will listen for shutdown signals (SIGINT, SIGTERM) and gracefully shutdown.
//
// Parameters:
//   - config: the proxy configuration loaded from LoadConfig
//
// Returns an error if the server fails to start or encounters an error during operation.
func StartServer(config *ProxyConfig) error {
	internalConfig := &Config{
		McpProxy:   config.McpProxy,
		McpServers: config.McpServers,
	}
	return startHTTPServer(internalConfig)
}

// StartServerWithContext starts the HTTP server with the provided configuration and context.
// This allows for more fine-grained control over server lifecycle.
//
// Parameters:
//   - ctx: context for controlling server lifecycle
//   - config: the proxy configuration loaded from LoadConfig
//
// Returns an error if the server fails to start or encounters an error during operation.
func StartServerWithContext(ctx context.Context, config *ProxyConfig) error {
	// For now, we'll use the existing startHTTPServer which handles its own context
	// In the future, this could be enhanced to use the provided context
	return StartServer(config)
}

// ProxyOptions represents the command-line options that can be used when starting the proxy
type ProxyOptions struct {
	ConfigPath  string
	Insecure    bool
	ExpandEnv   bool
	HTTPHeaders string
	HTTPTimeout int
}

// RunProxy is a convenience function that combines LoadConfig and StartServer.
// This is useful for simple use cases where you want to start the proxy with minimal setup.
//
// Parameters:
//   - opts: the proxy options (equivalent to command-line flags)
//
// Returns an error if configuration loading or server startup fails.
func RunProxy(opts ProxyOptions) error {
	config, err := LoadConfig(opts.ConfigPath, opts.Insecure, opts.ExpandEnv, opts.HTTPHeaders, opts.HTTPTimeout)
	if err != nil {
		return err
	}

	return StartServer(config)
}

// Default values for proxy options
const (
	DefaultConfigPath  = "config.json"
	DefaultExpandEnv   = true
	DefaultHTTPTimeout = 10
	DefaultInsecure    = false
	DefaultHTTPHeaders = ""
)

// NewDefaultProxyOptions creates a ProxyOptions struct with default values
func NewDefaultProxyOptions() ProxyOptions {
	return ProxyOptions{
		ConfigPath:  DefaultConfigPath,
		Insecure:    DefaultInsecure,
		ExpandEnv:   DefaultExpandEnv,
		HTTPHeaders: DefaultHTTPHeaders,
		HTTPTimeout: DefaultHTTPTimeout,
	}
}

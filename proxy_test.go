package mcpproxy

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Test loading config with default values
	opts := NewDefaultProxyOptions()

	if opts.ConfigPath != DefaultConfigPath {
		t.Errorf("Expected default config path %s, got %s", DefaultConfigPath, opts.ConfigPath)
	}

	if opts.ExpandEnv != DefaultExpandEnv {
		t.Errorf("Expected default expand env %v, got %v", DefaultExpandEnv, opts.ExpandEnv)
	}

	if opts.HTTPTimeout != DefaultHTTPTimeout {
		t.Errorf("Expected default HTTP timeout %d, got %d", DefaultHTTPTimeout, opts.HTTPTimeout)
	}
}

func TestProxyOptions(t *testing.T) {
	opts := ProxyOptions{
		ConfigPath:  "test-config.json",
		Insecure:    true,
		ExpandEnv:   false,
		HTTPHeaders: "Auth:Bearer token",
		HTTPTimeout: 30,
	}

	if opts.ConfigPath != "test-config.json" {
		t.Errorf("ConfigPath not set correctly")
	}

	if !opts.Insecure {
		t.Errorf("Insecure flag not set correctly")
	}
}

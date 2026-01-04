package otel

import "time"

// Option configures the Phoenix OTEL integration.
type Option func(*Config)

// WithEndpoint sets the Phoenix collector endpoint.
//
// Example endpoints:
//   - http://localhost:6006 (local Phoenix)
//   - https://app.phoenix.arize.com (Phoenix Cloud - requires SpaceID)
func WithEndpoint(endpoint string) Option {
	return func(c *Config) {
		c.Endpoint = endpoint
	}
}

// WithSpaceID sets the space identifier for Phoenix Cloud.
// When using Phoenix Cloud (app.phoenix.arize.com), set this to your space ID.
// The endpoint will be constructed as {Endpoint}/s/{SpaceID}.
func WithSpaceID(spaceID string) Option {
	return func(c *Config) {
		c.SpaceID = spaceID
	}
}

// WithProjectName sets the project name for traces.
func WithProjectName(name string) Option {
	return func(c *Config) {
		c.ProjectName = name
	}
}

// WithAPIKey sets the API key for authentication.
// This is required for Arize cloud.
func WithAPIKey(key string) Option {
	return func(c *Config) {
		c.APIKey = key
	}
}

// WithHeaders sets additional headers to send with requests.
func WithHeaders(headers map[string]string) Option {
	return func(c *Config) {
		c.Headers = headers
	}
}

// WithProtocol sets the transport protocol.
// Use ProtocolHTTP, ProtocolGRPC, or ProtocolInfer.
func WithProtocol(protocol Protocol) Option {
	return func(c *Config) {
		c.Protocol = protocol
	}
}

// WithBatch enables batch span processing.
// Recommended for production environments.
func WithBatch(batch bool) Option {
	return func(c *Config) {
		c.Batch = batch
	}
}

// WithBatchTimeout sets the maximum time to wait before exporting a batch.
func WithBatchTimeout(timeout time.Duration) Option {
	return func(c *Config) {
		c.BatchTimeout = timeout
	}
}

// WithBatchSize sets the maximum number of spans to batch.
func WithBatchSize(size int) Option {
	return func(c *Config) {
		c.BatchSize = size
	}
}

// WithGlobalProvider sets whether to register as the global tracer provider.
// Defaults to true.
func WithGlobalProvider(global bool) Option {
	return func(c *Config) {
		c.SetGlobalProvider = global
	}
}

// WithServiceName sets the service name for the resource.
func WithServiceName(name string) Option {
	return func(c *Config) {
		c.ServiceName = name
	}
}

// WithServiceVersion sets the service version for the resource.
func WithServiceVersion(version string) Option {
	return func(c *Config) {
		c.ServiceVersion = version
	}
}

// WithInsecure disables TLS for gRPC connections.
func WithInsecure(insecure bool) Option {
	return func(c *Config) {
		c.Insecure = insecure
	}
}

package phoenix

import (
	"net/http"
	"time"
)

// Option is a functional option for configuring the Client.
type Option func(*clientOptions)

// clientOptions holds the options for creating a Client.
type clientOptions struct {
	config     *Config
	httpClient *http.Client
	timeout    time.Duration
}

func defaultClientOptions() *clientOptions {
	return &clientOptions{
		config:  LoadConfig(),
		timeout: 60 * time.Second,
	}
}

// WithURL sets the API URL.
func WithURL(url string) Option {
	return func(o *clientOptions) {
		o.config.URL = url
	}
}

// WithAPIKey sets the API key for authentication.
func WithAPIKey(apiKey string) Option {
	return func(o *clientOptions) {
		o.config.APIKey = apiKey
	}
}

// WithProjectName sets the default project name.
func WithProjectName(projectName string) Option {
	return func(o *clientOptions) {
		o.config.ProjectName = projectName
	}
}

// WithSpaceID sets the space identifier for Phoenix Cloud.
// When using Phoenix Cloud (app.phoenix.arize.com), set this to your space ID.
// The URL will be constructed as {BaseURL}/s/{SpaceID}.
func WithSpaceID(spaceID string) Option {
	return func(o *clientOptions) {
		o.config.SpaceID = spaceID
	}
}

// WithConfig sets the entire configuration.
func WithConfig(config *Config) Option {
	return func(o *clientOptions) {
		o.config = config
	}
}

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(client *http.Client) Option {
	return func(o *clientOptions) {
		o.httpClient = client
	}
}

// WithTimeout sets the request timeout.
func WithTimeout(timeout time.Duration) Option {
	return func(o *clientOptions) {
		o.timeout = timeout
	}
}

// ListOption is a functional option for list operations.
type ListOption func(*listOptions)

type listOptions struct {
	cursor string
	limit  int
}

func defaultListOptions() *listOptions {
	return &listOptions{
		limit: 100,
	}
}

// WithCursor sets the pagination cursor.
func WithCursor(cursor string) ListOption {
	return func(o *listOptions) {
		o.cursor = cursor
	}
}

// WithLimit sets the maximum number of items to return.
func WithLimit(limit int) ListOption {
	return func(o *listOptions) {
		o.limit = limit
	}
}

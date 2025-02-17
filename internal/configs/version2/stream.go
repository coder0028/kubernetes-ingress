package version2

// TransportServerConfig holds NGINX configuration for a TransportServer.
type TransportServerConfig struct {
	Server         StreamServer
	Upstreams      []StreamUpstream
	StreamSnippets []string
}

// StreamUpstream defines a stream upstream.
type StreamUpstream struct {
	Name           string
	Servers        []StreamUpstreamServer
	UpstreamLabels UpstreamLabels
}

// StreamUpstreamServer defines a stream upstream server.
type StreamUpstreamServer struct {
	Address     string
	MaxFails    int
	FailTimeout string
}

// StreamServer defines a server in the stream module.
type StreamServer struct {
	TLSPassthrough           bool
	UnixSocket               string
	Port                     int
	UDP                      bool
	StatusZone               string
	ProxyRequests            *int
	ProxyResponses           *int
	ProxyPass                string
	Name                     string
	Namespace                string
	ProxyTimeout             string
	ProxyConnectTimeout      string
	ProxyNextUpstream        bool
	ProxyNextUpstreamTimeout string
	ProxyNextUpstreamTries   int
	HealthCheck              *StreamHealthCheck
	ServerSnippets           []string
}

// StreamHealthCheck defines a health check for a StreamUpstream in a StreamServer.
type StreamHealthCheck struct {
	Enabled  bool
	Interval string
	Port     int
	Passes   int
	Jitter   string
	Fails    int
	Timeout  string
}

// TLSPassthroughHostsConfig defines a mapping between TLS Passthrough hosts and the corresponding unix sockets.
type TLSPassthroughHostsConfig map[string]string

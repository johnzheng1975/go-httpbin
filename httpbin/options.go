package httpbin

import "time"

// OptionFunc uses the "functional options" pattern to customize an HTTPBin
// instance
type OptionFunc func(*HTTPBin)

// WithDefaultParams sets the default params handlers will use
func WithDefaultParams(defaultParams DefaultParams) OptionFunc {
	return func(h *HTTPBin) {
		h.DefaultParams = defaultParams
	}
}

// WithMaxBodySize sets the maximum amount of memory
func WithMaxBodySize(m int64) OptionFunc {
	return func(h *HTTPBin) {
		h.MaxBodySize = m
	}
}

// WithMaxDuration sets the maximum amount of time httpbin may take to respond
func WithMaxDuration(d time.Duration) OptionFunc {
	return func(h *HTTPBin) {
		h.MaxDuration = d
	}
}

// WithHostname sets the hostname to return via the /hostname endpoint.
func WithHostname(s string) OptionFunc {
	return func(h *HTTPBin) {
		h.hostname = s
	}
}

// WithServerEnv sets the server_env to return via the /get endpoint.
func WithServerEnv(s string) OptionFunc {
	return func(h *HTTPBin) {
		h.server_env = s
	}
}

// WithObserver sets the request observer callback
func WithObserver(o Observer) OptionFunc {
	return func(h *HTTPBin) {
		h.Observer = o
	}
}

// WithAllowedRedirectDomains limits the domains to which the /redirect-to
// endpoint will redirect traffic.
func WithAllowedRedirectDomains(hosts []string) OptionFunc {
	return func(h *HTTPBin) {
		hostSet := make(map[string]struct{}, len(hosts))
		for _, host := range hosts {
			hostSet[host] = struct{}{}
		}
		h.AllowedRedirectDomains = hostSet
	}
}

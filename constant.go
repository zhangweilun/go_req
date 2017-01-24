package gor

import (
	"time"
	"errors"
)

/**
* 
* @author willian
* @created 2017-01-24 13:38
* @email 18702515157@163.com  
**/
const (
	localUserAgent = "GRequests/0.10"

	// Default value for net.Dialer Timeout
	dialTimeout = 30 * time.Second

	// Default value for net.Dialer KeepAlive
	dialKeepAlive = 30 * time.Second

	// Default value for http.Transport TLSHandshakeTimeout
	tlsHandshakeTimeout = 10 * time.Second

	// Default value for Request Timeout
	requestTimeout = 90 * time.Second
)

var (
	// ErrRedirectLimitExceeded is the error returned when the request responded
	// with too many redirects
	ErrRedirectLimitExceeded = errors.New("grequests: Request exceeded redirect count")

	// RedirectLimit is a tunable variable that specifies how many times we can
	// redirect in response to a redirect. This is the global variable, if you
	// wish to set this on a request by request basis, set it within the
	// `RequestOptions` structure
	Redirect_limit = 30

	// SensitiveHTTPHeaders is a map of sensitive HTTP headers that a user
	// doesn't want passed on a redirect. This is the global variable, if you
	// wish to set this on a request by request basis, set it within the
	// `RequestOptions` structure
	SensitiveHTTPHeaders = map[string]struct{}{
		"Www-Authenticate":    {},
		"Authorization":       {},
		"Proxy-Authorization": {},
	}
)
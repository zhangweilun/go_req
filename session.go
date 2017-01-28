package gor

import "net/http"

/**
* 
* @author willian
* @created 2017-01-28 20:19
* @email 18702515157@163.com  
**/

// Session allows a user to make use of persistent cookies in between
// HTTP requests
type Session struct {
	Request *Request_options
	HttpClient *http.Client
}


// NewSession returns a session struct which enables can be used to maintain establish a persistent state with the
// server
// This function will set UseCookieJar to true as that is the purpose of using the session
func NewSession(r *Request_options) *Session {
	if r ==nil{
		r = &Request_options{}
	}
	r.Use_cookieJar = true
	return &Session{Request:r,HttpClient:BuildHTTPClient(*r)}
}

// Combine session options and request options
// 1. UserAgent
// 2. Host
// 3. Auth
// 4. Headers
func (s *Session) combineRequestOptions(ro *Request_options) *Request_options {
	if ro == nil {
		ro = &Request_options{}
	}

	if ro.UserAgent == "" && s.Request.UserAgent != "" {
		ro.UserAgent = s.Request.UserAgent
	}

	if ro.Host == "" && s.Request.Host != "" {
		ro.Host = s.Request.Host
	}

	if ro.Auth == nil && s.Request.Auth != nil {
		ro.Auth = s.Request.Auth
	}

	if len(s.Request.Headers) > 0 || len(ro.Headers) > 0 {
		headers := make(map[string]string)
		for k, v := range s.Request.Headers {
			headers[k] = v
		}
		for k, v := range ro.Headers {
			headers[k] = v
		}
		ro.Headers = headers
	}
	return ro
}

// Get takes 2 parameters and returns a Response Struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Get(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("GET", url, ro, s.HttpClient)
}

// Put takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Put(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("PUT", url, ro, s.HttpClient)
}

// Patch takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Patch(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("PATCH", url, ro, s.HttpClient)
}

// Delete takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Delete(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("DELETE", url, ro, s.HttpClient)
}

// Post takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Post(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("POST", url, ro, s.HttpClient)
}

// Head takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Head(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("HEAD", url, ro, s.HttpClient)
}

// Options takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
// A new session is created by calling NewSession with a request options struct
func (s *Session) Options(url string, ro *Request_options) (*Response, error) {
	ro = s.combineRequestOptions(ro)
	return doSessionRequest("OPTIONS", url, ro, s.HttpClient)
}

// CloseIdleConnections closes the idle connections that a session client may make use of
func (s *Session) CloseIdleConnections() {
	s.HttpClient.Transport.(*http.Transport).CloseIdleConnections()
}
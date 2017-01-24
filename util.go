package gor

import (
	"net/http"
	"runtime"
	"io"
)

/**
* 
* @author willian
* @created 2017-01-24 13:52
* @email 18702515157@163.com  
**/

// EnsureTransporterFinalized will ensure that when the HTTP client is GCed
// the runtime will close the idle connections (so that they won't leak)
// this function was adopted from Hashicorp's go-cleanhttp package
func EnsureTransporterFinalized(httpTransport *http.Transport) {
	runtime.SetFinalizer(&httpTransport, func(transportInt **http.Transport) {
		(*transportInt).CloseIdleConnections()
	})
}


// XMLCharDecoder is a helper type that takes a stream of bytes (not encoded in
// UTF-8) and returns a reader that encodes the bytes into UTF-8. This is done
// because Go's XML library only supports XML encoded in UTF-8
type XMLCharDecoder func(charset string, input io.Reader) (io.Reader, error)


func addRedirectFunctionality(client *http.Client, ro *Request_options) {
	if client.CheckRedirect != nil {
		return
	}
	client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		if ro.Redirect_limit == 0 {
			ro.Redirect_limit = Redirect_limit
		}

		if len(via) >= ro.Redirect_limit {
			return ErrRedirectLimitExceeded
		}

		if ro.SensitiveHTTPHeaders == nil {
			ro.SensitiveHTTPHeaders = SensitiveHTTPHeaders
		}

		for k, vv := range via[0].Header {
			// Is this a sensitive header?
			if _, found := ro.SensitiveHTTPHeaders[k]; found {
				continue
			}

			for _, v := range vv {
				req.Header.Add(k, v)
			}
		}

		return nil
	}
}
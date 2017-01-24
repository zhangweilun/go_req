package gor

/**
* 
* @author willian
* @created 2017-01-24 18:43
* @email 18702515157@163.com  
**/

// Get takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A Request_options struct
// If you do not intend to use the `Request_options` you can just pass nil
func Get(url string, ro *Request_options) (*Response,error) {
	return reg("GET",url,ro)
}

// Put takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
func Put(url string, ro *Request_options) (*Response, error) {
	return reg("PUT", url, ro)
}

// Patch takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
func Patch(url string, ro *Request_options) (*Response, error) {
	return reg("PATCH", url, ro)
}

// Delete takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
func Delete(url string, ro *Request_options) (*Response, error) {
	return reg("DELETE", url, ro)
}

// Post takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
func Post(url string, ro *Request_options) (*Response, error) {
	return reg("POST", url, ro)
}

// Head takes 2 parameters and returns a Response channel. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
func Head(url string, ro *Request_options) (*Response, error) {
	return reg("HEAD", url, ro)
}

// Options takes 2 parameters and returns a Response struct. These two options are:
// 	1. A URL
// 	2. A RequestOptions struct
// If you do not intend to use the `RequestOptions` you can just pass nil
func Options(url string, ro *Request_options) (*Response, error) {
	return reg("OPTIONS", url, ro)
}
package gor

import (
	"net/http"
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"encoding/xml"
)

type Response struct {
	// Ok is a boolean flag that validates that the server returned a 200 code
	Ok bool

	// This is the Go error flag – if something went wrong within the request, this flag will be set.
	Error error

	// We want to abstract (at least at the moment) the Go http.Response object away. So we are going to make use of it
	// internal but not give the user access
	RawResponse *http.Response

	// Status is the HTTP Status Code returned by the HTTP Response. Taken from resp.StatusCode
	Status int

	// Header is a net/http/Header structure
	Header http.Header

	internalByteBuffer *bytes.Buffer

}

func build_response(res *http.Response, err error) (*Response, error) {
	if err != nil {
		return &Response{Error:err},err
	}else {
		response := &Response{
			Ok:                res.StatusCode >= 200 && res.StatusCode < 300,
			Error:             nil,
			RawResponse:       res,
			Status:            res.StatusCode,
			Header:            res.Header,
			internalByteBuffer:bytes.NewBuffer([]byte{}),
		}
		return response,nil
	}
}

func (res *Response) Read(p []byte) (n int, err error)  {
	if res.Error != nil {
		return -1,res.Error
	}else {
		return res.RawResponse.Body.Read(p)
	}
}

// Close is part of our ability to support io.ReadCloser if someone wants to make use of the raw body
func (res *Response) Close() error {

	if res.Error != nil {
		return res.Error
	}

	io.Copy(ioutil.Discard, res)

	return res.RawResponse.Body.Close()
}

// DownloadToFile allows you to download the contents of the response to a file
func (res *Response) DownloadToFile(fileName string) error {

	if res.Error != nil {
		return res.Error
	}

	fd, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer res.Close() // This is a noop if we use the internal ByteBuffer
	defer fd.Close()

	if _, err := io.Copy(fd, res.getInternalReader()); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// getInternalReader because we implement io.ReadCloser and optionally hold a large buffer of the response (created by
// the user's request)
func (res *Response) getInternalReader() io.Reader {

	if res.internalByteBuffer.Len() != 0 {
		return res.internalByteBuffer
	}
	return res
}

// XML is a method that will populate a struct that is provided `userStruct` with the XML returned within the
// response body
func (res *Response) XML(userStruct interface{}, charsetReader XMLCharDecoder) error {

	if res.Error != nil {
		return res.Error
	}

	xmlDecoder := xml.NewDecoder(res.getInternalReader())

	if charsetReader != nil {
		xmlDecoder.CharsetReader = charsetReader
	}

	defer res.Close()

	if err := xmlDecoder.Decode(&userStruct); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// populateResponseByteBuffer is a utility method that will populate the internal byte reader – this is largely used for .String()
// and .Bytes()
func (res *Response) populateResponseByteBuffer() {

	// Have I done this already?
	if res.internalByteBuffer.Len() != 0 {
		return
	}

	defer res.Close()

	// Is there any content?
	if res.RawResponse.ContentLength == 0 {
		return
	}

	// Did the server tell us how big the response is going to be?
	if res.RawResponse.ContentLength > 0 {
		res.internalByteBuffer.Grow(int(res.RawResponse.ContentLength))
	}

	if _, err := io.Copy(res.internalByteBuffer, res); err != nil && err != io.EOF {
		res.Error = err
		res.RawResponse.Body.Close()
	}

}

// Bytes returns the response as a byte array
func (res *Response) Bytes() []byte {

	if res.Error != nil {
		return nil
	}

	res.populateResponseByteBuffer()

	// Are we still empty?
	if res.internalByteBuffer.Len() == 0 {
		return nil
	}
	return res.internalByteBuffer.Bytes()

}

// String returns the response as a string
func (res *Response) String() string {
	if res.Error != nil {
		return ""
	}

	res.populateResponseByteBuffer()

	return res.internalByteBuffer.String()
}

// ClearInternalBuffer is a function that will clear the internal buffer that we use to hold the .String() and .Bytes()
// data. Once you have used these functions – you may want to free up the memory.
func (res *Response) ClearInternalBuffer() {

	if res == nil || res.internalByteBuffer == nil {
		return
	}

	res.internalByteBuffer.Reset()
}
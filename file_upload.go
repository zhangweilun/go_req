package gor

import "io"

/**
* 
* @author willian
* @created 2017-01-24 13:23
* @email 18702515157@163.com  
**/

type File_upload struct {
	// Filename is the name of the file that you wish to upload. We use this to guess the mimetype as well as pass it onto the server
	FileName string

	// FileContents is happy as long as you pass it a io.ReadCloser (which most file use anyways)
	FileContents io.ReadCloser

	// FieldName is form field name
	FieldName string
}
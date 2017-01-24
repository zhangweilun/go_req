package gor

import "testing"

/**
* 
* @author willian
* @created 2017-01-24 18:36
* @email 18702515157@163.com  
**/

func TestAddQueryParams(t *testing.T) {
	url, err := buildURLParams("http://www.genshuixue.com/i-cxy/p/7648250", map[string]string{"1":"2", "3":"4"})
	if err != nil {
		t.Error("url parse error",err)
	}
	t.Log(url)
}

func TestAddQueryStringParamsExistingParam(t *testing.T) {
	userURL, err := buildURLParams("https://www.google.com/?5=6", map[string]string{"3": "4", "1": "2"})

	if err != nil {
		t.Error("URL Parse Error: ", err)
	}

	if userURL != "https://www.google.com/?1=2&3=4&5=6" {
		t.Error("URL params not properly built and sorted", userURL)
	}
}
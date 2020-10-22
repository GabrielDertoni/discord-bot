package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func simplePOST(url URL, data interface{}, headers map[string]string) (response []byte, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return
	}
	return rawPOST(url, jsonStr, headers)
}

func rawPOST(url URL, jsonStr []byte, headers map[string]string) (response []byte, err error) {
	req, err := http.NewRequest("POST", string(url), bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}

func simpleGET(url URL, headers map[string]string) (response []byte, err error) {
	req, err := http.NewRequest("GET", string(url), nil)
	if err != nil {
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}

func simplePATCH(url URL, data interface{}, headers map[string]string) (response []byte, err error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest("PATCH", string(url), bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return body, nil
}

func (a URL) Concat(b URL) URL {
	bytesA := []byte(a)
	bytesB := []byte(b)
	if bytesA[len(bytesA)-1] == '/' {
		if bytesB[0] == '/' {
			return URL(append(bytesA, bytesB[1:]...))
		}
		return URL(append(bytesA, bytesB...))
	} else if bytesB[0] == '/' {
		return URL(append(bytesA, bytesB...))
	}
	return URL(fmt.Sprintf("%s/%s", bytesA, bytesB))
}

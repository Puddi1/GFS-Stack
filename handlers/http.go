package handlers

import (
	"bytes"
	"io"
	"net/http"
)

type RequestHTTP struct {
	MethodHTTP string
	Url        string
	Body       []byte
	Headers    [][2]string
}

// HandleRequestHTTP takes an handlers.RequestHTTP reference and executes a HTTP request
func HandleRequestHTTP(r *RequestHTTP) (*http.Response, error) {
	// create client
	hc := http.Client{}
	// create request
	req, errRequest := http.NewRequest(r.MethodHTTP, r.Url, bytes.NewReader(r.Body))
	if errRequest != nil {
		return nil, errRequest
	}
	// add headers to request
	for i := 0; i < len(r.Headers); i++ {
		req.Header.Add(r.Headers[i][0], r.Headers[i][1])
	}
	// Make request
	resp, errResponse := hc.Do(req)
	if errResponse != nil {
		return nil, errResponse
	}
	return resp, nil
}

// HandleResponseBodyToString takes a http.Response reference and returns the body as string
// without incurring memory leaks.
func HandleResponseBodyToString(r *http.Response) (string, error) {
	// read response body
	stringBody, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	// close response body
	r.Body.Close()
	// print response body
	return string(stringBody), nil
}

package obsgen

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"fmt"
)

func TestLogEvent(t *testing.T) {
	// create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check that the request is correct
		if req.Method != "POST" {
			t.Errorf("unexpected HTTP method: %s", req.Method)
			return
		}

		if req.URL.Path != "/v0/base/table" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
			return
		}

		if req.Header.Get("Authorization") != "Bearer key" {
			t.Errorf("unexpected Authorization header: %s", req.Header.Get("Authorization"))
			return
		}

		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected Content-Type header: %s", req.Header.Get("Content-Type"))
			return
		}

		// read the request body
		var requestBody map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			t.Errorf("error decoding request body: %s", err)
			return
		}

		// check that the request body is correct
		expectedBody := map[string]interface{}{
			"records": []map[string]interface{}{
				{"fields": map[string]interface{}{
					"foo": "bar",
				}},
			},
		}

		requestJSON, _ := json.Marshal(requestBody)
		expectedJSON, _ := json.Marshal(expectedBody)

		if string(requestJSON) != string(expectedJSON) {
			t.Errorf("unexpected request body: %s", string(requestJSON))
			return
		}

		// write the response
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	// create a new ObsGenClient with the mock endpoint
	client := &ObsGenClient{
		apiKey:   "key",
		baseID:   "base",
		table:    "table",
		client:   server.Client(),
		endpoint: server.URL + "/v0/base/table",
	}

	// call the LogEvent method with test data
	c := client.LogEvent(map[string]interface{}{
		"foo": "bar",
	})

	err := <-c

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}

func TestLogEventNoErr(t *testing.T) {
	// create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check that the request is correct
		if req.Method != "POST" {
			t.Errorf("unexpected HTTP method: %s", req.Method)
			return
		}

		if req.URL.Path != "/v0/base/table" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
			return
		}

		if req.Header.Get("Authorization") != "Bearer key" {
			t.Errorf("unexpected Authorization header: %s", req.Header.Get("Authorization"))
			return
		}

		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected Content-Type header: %s", req.Header.Get("Content-Type"))
			return
		}

		// read the request body
		var requestBody map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			t.Errorf("error decoding request body: %s", err)
			return
		}

		// check that the request body is correct
		expectedBody := map[string]interface{}{
			"records": []map[string]interface{}{
				{"fields": map[string]interface{}{
					"foo": "bar",
				}},
			},
		}

		requestJSON, _ := json.Marshal(requestBody)
		expectedJSON, _ := json.Marshal(expectedBody)

		if string(requestJSON) != string(expectedJSON) {
			t.Errorf("unexpected request body: %s", string(requestJSON))
			return
		}

		// write the response
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	// create a new ObsGenClient with the mock endpoint
	client := &ObsGenClient{
		apiKey:   "key",
		baseID:   "base",
		table:    "table",
		client:   server.Client(),
		endpoint: server.URL + "/v0/base/table",
	}

	fmt.Println("hi1")

	// call the LogEvent method with test data
	client.LogEvent(map[string]interface{}{
		"foo": "bar",
	})

	fmt.Println("hi2")

	client.LogEvent(map[string]interface{}{
		"foo": "bar",
	})

	fmt.Println("hi3")
}

func TestLogEventHang(t *testing.T) {
	// create a mock HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// check that the request is correct
		if req.Method != "POST" {
			t.Errorf("unexpected HTTP method: %s", req.Method)
			return
		}

		if req.URL.Path != "/v0/base/table" {
			t.Errorf("unexpected URL path: %s", req.URL.Path)
			return
		}

		if req.Header.Get("Authorization") != "Bearer key" {
			t.Errorf("unexpected Authorization header: %s", req.Header.Get("Authorization"))
			return
		}

		if req.Header.Get("Content-Type") != "application/json" {
			t.Errorf("unexpected Content-Type header: %s", req.Header.Get("Content-Type"))
			return
		}

		// read the request body
		var requestBody map[string]interface{}
		err := json.NewDecoder(req.Body).Decode(&requestBody)
		if err != nil {
			t.Errorf("error decoding request body: %s", err)
			return
		}

		// check that the request body is correct
		expectedBody := map[string]interface{}{
			"records": []map[string]interface{}{
				{"fields": map[string]interface{}{
					"foo": "bar",
				}},
			},
		}

		requestJSON, _ := json.Marshal(requestBody)
		expectedJSON, _ := json.Marshal(expectedBody)

		if string(requestJSON) != string(expectedJSON) {
			t.Errorf("unexpected request body: %s", string(requestJSON))
			return
		}

		// write the response
		rw.WriteHeader(http.StatusOK)
	}))

	defer server.Close()

	// create a new ObsGenClient with the mock endpoint
	client := &ObsGenClient{
		apiKey:   "key",
		baseID:   "base",
		table:    "table",
		client:   server.Client(),
		endpoint: server.URL + "/v0/base/table",
	}

	// call the LogEvent method with test data
	c := client.LogEvent(map[string]interface{}{
		"foo": "bar",
	})

	err := <-c
	err2 := <-c

	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
//go:build unit
// +build unit

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// Based on the blog post at https://blog.alexellis.io/golang-writing-unit-tests/
func Test_GetAstronauts_NoInputs_ReturnsAstronautData(t *testing.T) {

	client := GetFakeHttpClient()
	result := GetAstronauts(client)

	AssertDataReceived(t, result)
	AssertAstronautCountIsOne(t, result)
}

func GetFakeHttpClient() *http.Client {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"people": [{"craft": "The-Hab", "name": "Mark Watney"}]}`))
	}))

	defer ts.Close()

	client := ts.Client()
	return client
}

func AssertDataReceived(t *testing.T, data people) {
	if data == (people{}) {
		t.Error("No data received")
	}
}

func AssertAstronautCountIsOne(t *testing.T, data people) {
	if data.Number != 1 {
		t.Errorf("Incorrect Astronaut data received.  Count was %d", data.Number)

	}
}

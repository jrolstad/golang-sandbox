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

	testServer := GetTestServer()
	defer testServer.Close()

	client := testServer.Client()
	result := GetAstronauts(client, testServer.URL)

	AssertDataReceived(t, result)
	AssertAstronautCountIsOne(t, result)
}

func GetTestServer() *httptest.Server {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(GetTestData()))
	}))

	return testServer
}

func GetTestData() string {
	return `{"people": [{"craft": "The-Hab", "name": "Mark Watney"}],"number": 1}`
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

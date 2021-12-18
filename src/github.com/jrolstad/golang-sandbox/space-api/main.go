package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type people struct {
	Number int `json:number`
}

func main() {
	url := "http://api.open-notify.org/astros.json"

	req := CreateRequest(url)
	res := ExecuteRequest(req)

	body := ReadResponse(res)
	data := ParseJson(body)

	fmt.Println(data.Number)
}

func CreateRequest(url string) *http.Request {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal((err))
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")
	return req
}

func ExecuteRequest(req *http.Request) *http.Response {
	httpClient := http.Client{
		Timeout: time.Second * 2,
	}
	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	return res
}

func ReadResponse(res *http.Response) []byte {
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	return body
}

func ParseJson(body []byte) people {
	people1 := people{}
	jsonErr := json.Unmarshal(body, &people1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return people1
}

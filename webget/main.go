package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: time.Second * 1,
	}
	// direct verb call
	res, err := client.Get("https://google.com.br")
	if err != nil {
		fmt.Println("[main] error fetching google's web page", err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("[main] error reading google's page content", err.Error())
			return
		}
		fmt.Println(string(body))
	} else {
		fmt.Println("[main] fetching google's page returned status = ", res.StatusCode)
		return
	}

	// request from scratch
	request, err := http.NewRequest(http.MethodGet, "https://google.com.br", nil /* content as io.Reader */)
	if err != nil {
		fmt.Println("[main] error creating google page request", err.Error())
		return
	}
	// manipulating
	request.SetBasicAuth("user", "password")
	res, err = client.Do(request)
	// treatment is the same
	if err != nil {
		fmt.Println("[main] error fetching google's web page", err.Error())
		return
	}
	defer res.Body.Close()

	if res.StatusCode == 200 {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println("[main] error reading google's page content", err.Error())
			return
		}
		fmt.Println("\n-- new request --")
		fmt.Println(string(body))
	} else {
		fmt.Println("[main] fetching google's page returned status = ", res.StatusCode)
		return
	}
}

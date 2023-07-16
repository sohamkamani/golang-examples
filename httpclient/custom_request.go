/*
 * This example shows how to make an HTTP request using the custom Request type
 * from the net/http package. This is useful if you want to set custom methods, headers,
 * or otherwise control the request parameters.
 */

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// we need to use a *url.URL type instead of a string
	u, err := url.Parse("http://localhost:3000/some/path")
	if err != nil {
		log.Fatal(err)
	}
	// create request body
	body := ioutil.NopCloser(strings.NewReader("This is the request body."))

	req := &http.Request{
		// we can set a custom method here
		Method: http.MethodPut, // or http.MethodDelete, http.MethodPatch, etc.
		URL:    u,
		Body:   body,
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}

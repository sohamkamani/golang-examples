/*
 * This example shows how to make an HTTP request using the custom Request type
 * from the net/http package. This is useful if you want to set custom methods, headers,
 * or otherwise control the request parameters.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	// declare the request url and body
	url := "http://localhost:3000/some/path"
	body := strings.NewReader("This is the request body.")

	// we can set a custom method here, like http.MethodPut
	// or http.MethodDelete, http.MethodPatch, etc.
	req, err := http.NewRequest(http.MethodPut, url, body)
	if err != nil {
		log.Fatal(err)
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

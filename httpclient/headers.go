/*
 * This example shows how to work with headers when sending a request.
 * which includes setting headers in the request, and reading headers from the response.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "http://localhost:3000"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// add a custom header to the request
	// here we specify the header name and value as arguments
	req.Header.Add("X-Custom-Header", "custom-value")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print all the response headers
	fmt.Println("Response headers:")
	for k, v := range resp.Header {
		fmt.Printf("%s: %s\n", k, v)
	}
}

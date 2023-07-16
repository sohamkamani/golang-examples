/*
 * This example shows how to make a simple HTTP POST request using the net/http
 * package. We will send a POST request to http://www.example.com, along with
 * some data in the request body.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "http://www.example.com"

	// create post body
	body := strings.NewReader("This is the request body.")

	resp, err := http.Post(url, "text/plain", body)
	if err != nil {
		// we will get an error at this stage if the request fails, such as if the
		// requested URL is not found, or if the server is not reachable.
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}

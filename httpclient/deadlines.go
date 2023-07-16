/* This example shows us how to set a dealine for an http request to complete
 * This is useful if we want to ensure that our request
 * does not take too long to complete, and we want to cancel it if it does.
 */

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	url := "http://www.example.com"

	http.DefaultClient.Timeout = 50 * time.Millisecond
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// print the status code
	fmt.Println("Status:", resp.Status)
}

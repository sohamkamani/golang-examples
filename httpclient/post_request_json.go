/*
 * This example shows how to make a simple HTTP POST request using the net/http
 * package. We will send a POST request along with
 * some JSON data in the request body.
 */

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Person is a struct that represents the data we will send in the request body
type Person struct {
	Name string
	Age  int
}

func main() {
	url := "http://localhost:3000"

	// create post body using an instance of the Person struct
	p := Person{
		Name: "John Doe",
		Age:  25,
	}
	// convert p to JSON data
	jsonData, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	// We can set the content type here
	resp, err := http.Post(url, "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// parse the response

	// responsePerson := Person{}
	// err = json.NewDecoder(resp.Body).Decode(&responsePerson)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Status:", resp.Status)
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "https://httpbin.org/ip"

func main() {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	// fail on a non http 200 status code
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}

	// read the body
	b, err := ioutil.ReadAll(resp.Body)

	// unmarshal the response into the ip struct
	result := ip{}
	if err := json.Unmarshal(b, &result); err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", result)
}

// ip handles the response from the /ip endpoint
type ip struct {
	Origin string `json:"origin"`
}

// String returns a friendly representation
func (i ip) String() string {
	return fmt.Sprintf("%s", i.Origin)
}

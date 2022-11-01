package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Item struct {
	Count int
	Description string
	Auth string
	Link string
}

type Response struct {
	Entries []struct {
		Item
	}
}

func main() {
	res, err := http.Get("https://api.publicapis.org/entries")

	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	if err != nil {
		log.Fatal(err)
	}

	r := new(Response)

	err = json.NewDecoder(res.Body).Decode(r)

	for _, child := range r.Entries {
		fmt.Println(child.Description)
	}
}

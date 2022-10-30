package main

import (
	"net/http"
	"log"
	"io"
	"os"
)

func main()  {
	res, err := http.Get("https://api.publicapis.org/entries")

	if err != nil {
		log.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}

	_, err = io.Copy(os.Stdout, res.Body)

	if err != nil {
		log.Fatal(err)
	}
}

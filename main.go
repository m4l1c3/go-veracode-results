package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	pass := os.Getenv("vcresultsapipass")
	user := os.Getenv("vcresultsapiuser")
	body := strings.NewReader("('include_user_info', 'true')")

	apiEndpoints := [1]string{
		"https://analysiscenter.veracode.com/api/5.0/getapplist.do",
	}

	req, err := http.NewRequest("GET", apiEndpoints[0], body)
	if err != nil {
		// handle err
	}

	req.SetBasicAuth(user, pass)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	var reader io.ReadCloser

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			fmt.Printf("Error decoding response: %s\n", err)
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	_, err = io.Copy(os.Stdout, reader)
	if err != nil {
		log.Fatal(err)
	}
}

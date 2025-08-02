package http_client

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
)

func PrintGetResult(path string) {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Printf("error returned is %v\n", err.Error())
		return
	}

	defer resp.Body.Close()
	fmt.Println("Response status : ", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// todo add json decoding
	// todo add redirects
}

func PrintPostResult(path string, contentType string, content string) {
	resp, err := http.Post(path, contentType, bytes.NewBuffer([]byte(content)))
	if err != nil {
		fmt.Printf("error returned is %v\n", err.Error())
		return
	}

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

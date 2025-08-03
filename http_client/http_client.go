package http_client

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/ishanchopra90/Practice/practice/http_data"
)

func GetPost(path string) string {
	resp, err := http.Get(path)
	if err != nil {
		fmt.Printf("http get error returned is %v\n", err.Error())
		return ""
	}

	defer resp.Body.Close()
	fmt.Println("Response status : ", resp.Status)
	scanner := bufio.NewScanner(resp.Body)
	finalText := ""
	for scanner.Scan() {
		/*
			fmt.Println("New text chunk")
			fmt.Println(scanner.Text())
		*/

		finalText += scanner.Text()
	}

	return finalText
}

func GetPostWithHeaders(path string, lastNameHeader string) string {
	httpRequest, err := http.NewRequest("GET", path, nil)
	if err != nil {
		fmt.Println("GetPostWithHeaders returned error ", err.Error())
	}
	httpRequest.Header.Add("LastName", lastNameHeader)
	resp, err := http.DefaultClient.Do(httpRequest)
	if err != nil {
		fmt.Println("GetPostWithHeaders returned error in client do ", err.Error())
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error in getting resp body ", err.Error())
	}

	return string(respBody)
}

func PostPost(path string, contentType string, content string) {
	fmt.Println("post to endpoint ", path)
	resp, err := http.Post(path, contentType, bytes.NewBuffer([]byte(content)))
	if err != nil {
		fmt.Printf("post error returned is %v\n", err.Error())
		return
	}
	/*
		fmt.Println("output from scanner.Scan")
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
	}*/
	posts := http_data.DecodeJson(resp.Body)
	for _, post := range posts {
		fmt.Println("Name ", post.Name)
		fmt.Println("Job ", post.Job)
	}
}

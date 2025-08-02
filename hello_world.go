package main

import (
	"fmt"

	"github.com/ishanchopra90/Practice/practice/http_client"
)

func main() {
	fmt.Println("Hello World")
	http_client.PrintGetResult("https://httpbin.org/get")
	http_client.PrintPostResult("https://httpbin.org/post", "application/json", `{"Name": "Ishan", "Job": "Engineer"}`)
}

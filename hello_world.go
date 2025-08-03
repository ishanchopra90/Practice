package main

import (
	"fmt"

	"github.com/ishanchopra90/Practice/practice/http_client"
	"github.com/ishanchopra90/Practice/practice/http_server"
)

func main() {
	fmt.Println("Hello World")

	/*
		fmt.Println(http_client.GetPost("https://httpbin.org/get"))
		http_client.PostPost("https://httpbin.org/post", "application/json", `{"Name": "Ishan", "Job": "Engineer"}`)
		http_client.PostPost("https://postman-echo.com/post", "application/json", `{"Name": "Ishan", "Job": "Engineer"}`)
	*/

	httpServer := http_server.MyServer{}
	go httpServer.Start()

	fmt.Println()
	fmt.Println("getting initial server cache. should be empty")
	fmt.Println(http_client.GetPost("http://localhost:8090/MyServer"))

	fmt.Println()
	fmt.Println("adding Name Ishan and Job Engineer")
	http_client.PostPost("http://localhost:8090/MyServer", "application/json", `{"Name": "Ishan", "Job": "Engineer"}`)

	fmt.Println()
	fmt.Println("getting server cache")
	fmt.Println(http_client.GetPost("http://localhost:8090/MyServer"))

	fmt.Println()
	fmt.Println("adding Name Snoopy and Job Pet")
	http_client.PostPost("http://localhost:8090/MyServer", "application/json", `{"Name": "Snoopy", "Job": "Pet"}`)

	fmt.Println()
	fmt.Println("getting server cache")
	fmt.Println(http_client.GetPost("http://localhost:8090/MyServer"))

	fmt.Println()
	fmt.Println("get requests with last name header")
	fmt.Println(http_client.GetPostWithHeaders("http://localhost:8090/MyServer", "Smith"))

	fmt.Println()
	fmt.Println("redirect to google")
	fmt.Println(http_client.GetPost("http://localhost:8090/MyServer/redirectToGoogle"))
}

package http_server

import (
	"net/http"

	"github.com/ishanchopra90/Practice/practice/http_data"
)

type MyServer struct {
	PostCache []*http_data.Post
}

func (myServer *MyServer) MyHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		posts := http_data.DecodeJson(req.Body)
		myServer.PostCache = append(myServer.PostCache, &posts[0])
	}

	// check for headers
	lastName := req.Header.Get("LastName")

	for _, post := range myServer.PostCache {
		if lastName != "" {
			post.Name = post.Name + " " + lastName
		}

		w.Write([]byte(post.String() + " "))
	}
}

func (myServer *MyServer) RedirectHandler(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req, "http://google.com", http.StatusMovedPermanently)
}

func (myServer *MyServer) Start() {
	http.HandleFunc("/MyServer", myServer.MyHandler)
	http.HandleFunc("/MyServer/redirectToGoogle", myServer.RedirectHandler)
	http.ListenAndServe(":8090", nil)
}

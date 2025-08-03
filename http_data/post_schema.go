package http_data

import (
	"encoding/json"
	"fmt"
	"io"
)

type Post struct {
	Job  string
	Name string
}

func (post *Post) String() string {
	return fmt.Sprintf("{\"Name\":\"%v\", \"Job\":\"%v\"}", post.Name, post.Job)
}

func DecodeJson(io io.Reader) []Post {
	//fmt.Println("output from json decoder")
	var posts []Post
	decoder := json.NewDecoder(io)
	for decoder.More() {
		post := Post{}
		err := decoder.Decode(&post)
		if err != nil {
			fmt.Println("error encountered while decoding json ", err.Error())
		}

		posts = append(posts, post)
	}

	return posts
}

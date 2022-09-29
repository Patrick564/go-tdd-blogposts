package main

import (
	"log"
	"os"

	blogposts "github.com/Patrick564/go-tdd-blogposts"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))

	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
}

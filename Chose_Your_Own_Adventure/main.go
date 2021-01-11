package main

import (
	"gitub/koas/cyoa/cyoa"
	"net/http"
)

var (
	story   *cyoa.Story
	service *cyoa.StoryService
	fs      http.Handler
)

func init() {
	jsonPath := "./cyoa/story.json"
	story = cyoa.NewStory()
	story.LoadStoryFromJsonFile(jsonPath)

	service = cyoa.NewStoryService(story)

	fs = http.FileServer(http.Dir("./static"))
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/", cyoa.NewStoryHandler(service))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	err := http.ListenAndServe(":5050", mux)

	if err != nil {
		panic(err)
	}
}

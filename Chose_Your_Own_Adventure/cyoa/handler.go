package cyoa

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("./cyoa/views/storyPage.html"))
}

type storyHandler struct {
	storyService *StoryService
}

func NewStoryHandler(s *StoryService) http.Handler {
	return storyHandler{storyService: s}
}

func (s storyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlPath := r.URL.Path[1:]
	newArc, err := s.storyService.GetStoryArc(urlPath)

	if err != nil {
		http.NotFound(w, r)
		return
	}

	tmpl.Execute(w, *newArc)
}

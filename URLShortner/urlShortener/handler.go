package urlshortener

import "net/http"

type Handler interface {
	GetPath(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service Service
}

func NewHanlder(urlService Service) Handler {
	return handler{service: urlService}
}

func (u handler) GetPath(w http.ResponseWriter, r *http.Request) {
	url := r.RequestURI

	path := u.service.GetPath(url)

	if path == "" {
		http.NotFound(w, r)
	}

	http.Redirect(w, r, path, http.StatusPermanentRedirect)
}

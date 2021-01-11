package urlshortener

type Reader interface {
	GetURL(string) string
}

type Repository interface {
	Reader
}

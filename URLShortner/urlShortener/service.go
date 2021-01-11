package urlshortener

type Service interface {
	GetPath(string) string
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo}
}

func (s *service) GetPath(urlPath string) string {
	return s.repo.GetURL(urlPath)
}

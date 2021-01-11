package urlshortener

import (
	"github/Ko4s/urlShortener/utils"

	"gopkg.in/yaml.v2"
)

type inMemoryRepo struct {
	urlMap map[string]string
}

type pathToURL struct {
	Path string
	Url  string
}

func NewInMemoryRepo(yamlPath string) (*inMemoryRepo, error) {
	fileContent, err := utils.OpenAndReadFile(yamlPath)

	if err != nil {
		return nil, err
	}

	var urlPaths = []pathToURL{}

	err = yaml.Unmarshal(fileContent, &urlPaths)

	if err != nil {
		return nil, err
	}

	urlMap := createMapOfURLsPaths(urlPaths)

	return &inMemoryRepo{
		urlMap: urlMap,
	}, nil
}

func (r *inMemoryRepo) GetURL(url string) string {
	return r.urlMap[url]
}

func createMapOfURLsPaths(uPaths []pathToURL) map[string]string {
	m := make(map[string]string)

	for _, el := range uPaths {
		m[el.Path] = el.Url
	}

	return m
}

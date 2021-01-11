package cyoa

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type Story struct {
	story map[string]interface{}
}

func NewStory() *Story {
	return &Story{}
}

func (s *Story) LoadStoryFromJsonFile(filePath string) error {

	file, err := ioutil.ReadFile(filePath)

	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &s.story)

	if err != nil {
		return err
	}

	return nil
}

//GetArc returns ArcOfStory
func (s *Story) GetArc(arcName string) (*arc, error) {
	arc := s.story[arcName]

	if arc == nil {
		return nil, errors.New("No arc")
	}

	return newArc(arc.(map[string]interface{})), nil
}

//Arcs
type arc struct {
	Title      string
	StoryLines string
	Options    []arcOptions
}

//NewArc new arc constuctor
func newArc(mapedArc map[string]interface{}) *arc {

	storyLines := mapedArc["story"].([]interface{})

	newStoryLines := make([]string, len(storyLines))
	for i, el := range storyLines {
		newStoryLines[i] = el.(string)
	}

	return &arc{
		Title:      mapedArc["title"].(string),
		StoryLines: strings.Join(newStoryLines, "\n"),
		Options:    createArcOptions(mapedArc["options"].([]interface{})),
	}
}

type arcOptions struct {
	Text    string
	NextArc string
}

//NewArcOptions new arc opton Constructor
func newArcOptions(option map[string]interface{}) arcOptions {
	return arcOptions{
		Text:    option["text"].(string),
		NextArc: option["arc"].(string),
	}
}

func createArcOptions(aOptions []interface{}) []arcOptions {

	options := []arcOptions{}

	for _, el := range aOptions {
		options = append(options, newArcOptions(el.(map[string]interface{})))
	}

	return options
}

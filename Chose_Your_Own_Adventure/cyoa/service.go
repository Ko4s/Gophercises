package cyoa

type StoryService struct {
	story *Story
}

func NewStoryService(s *Story) *StoryService {
	return &StoryService{story: s}
}

func (s *StoryService) GetStoryArc(arcName string) (*arc, error) {

	newArc, err := s.story.GetArc(arcName)

	if err != nil {
		return nil, err
	}

	return newArc, nil
}

package quiz

import (
	"encoding/csv"
	"io"
	"math/rand"
	"os"
	"time"
)

//a quiz Question struct
type Question struct {
	Question string
	Answer   string
}

//loadQuestions load question from csv file [question,naswer]
func loadQuestions(path string) ([]Question, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	r := csv.NewReader(file)

	questions := []Question{}

	for {
		record, err := r.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		question := Question{
			Question: record[0],
			Answer:   record[1],
		}

		questions = append(questions, question)
	}

	return questions, nil
}

func shuffleQuestions(s []Question) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}

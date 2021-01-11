package main

import (
	"flag"
	"github/K04s/quiz-app/quiz"
	"log"
	"time"
)

func main() {

	limit := flag.Duration("limit", 30*time.Second, "time limit of the game")
	flag.Parse()

	path := "questions.csv"

	game, err := quiz.NewGame(path, *limit)

	if err != nil {
		log.Fatalln(err)
	}

	game.Play()
}

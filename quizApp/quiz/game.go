package quiz

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

//Game is a Game object
type Game struct {
	Questions         []Question
	correctAnswer     int
	AmountOfQuestions int
	gameLength        time.Duration
}

func NewGame(path string, gameLength time.Duration) (*Game, error) {

	q, err := loadQuestions(path)

	if err != nil {
		return nil, err
	}

	shuffleQuestions(q)
	return &Game{
		Questions:         q,
		correctAnswer:     0,
		AmountOfQuestions: len(q),
		gameLength:        gameLength,
	}, nil
}

func (g *Game) Play() {

	scanner := bufio.NewScanner(os.Stdin)
	startGame(scanner)

	fmt.Println("Game will start shortly...")
	fmt.Printf("You have %v seconds\n", g.gameLength)

	ctx, cancel := context.WithTimeout(context.Background(), g.gameLength)
	defer cancel()

	go g.gameLoop(ctx, scanner)
	<-ctx.Done()

	fmt.Printf("\nYou guesed  %v/%v correct", g.correctAnswer, g.AmountOfQuestions)

}

func startGame(scanner *bufio.Scanner) {
	for {
		fmt.Println("Please press enter to play the game...")
		scanner.Scan()
		break
	}
}

func (g *Game) gameLoop(ctx context.Context, scanner *bufio.Scanner) {

	for _, question := range g.Questions {
		fmt.Printf("%v: ", question.Question)
		answerCh := make(chan string)
		go func() {
			scanner.Scan()
			answerCh <- scanner.Text()
		}()

		select {
		case <-ctx.Done():
			return
		case answer := <-answerCh:
			if answer == question.Answer {
				fmt.Println("Correct")
				g.correctAnswer++
			} else {
				fmt.Println("Sry wrong :(")
			}
		}
	}
}

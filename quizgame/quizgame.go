/* Creates the quiz game from gophercises.com */

package main

import (
	"math/rand"
	"fmt"
	"encoding/csv"
	"os"
	"bufio"
	"time"
	"flag"
)

func main(){
	//import csv
	timeLimit := flag.Int("time", 30, "quiz time limit")
	shuffleOption := flag.Bool("shuffle", false, "shuffle the order")
	flag.Parse()

	file, err := os.Open("problems.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	quizGameReader := csv.NewReader(file)
	quizQA, err := quizGameReader.ReadAll()
	if err != nil{
		fmt.Printf("Uh oh!")
	}

	// Shuffle the deck
	if *shuffleOption == true{
		rand.Shuffle(len(quizQA), func(i, j int){
			quizQA[i], quizQA[j] =
			quizQA[j], quizQA[i]
		})
	}

	fmt.Println("Press 'Enter' to start the timer...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	c := make(chan bool)
	var numberCorrect int32 = 0
	go quizGame(quizQA, &numberCorrect, c)
	select {
		case <-c:

		case <- time.After(time.Duration(*timeLimit) * time.Second):
		}
	fmt.Printf("\nYou got %v right out of %v",
	numberCorrect, len(quizQA))
	return
}

func quizGame(qa [][]string, numberCorrect *int32, c chan bool){
	scanner := bufio.NewScanner(os.Stdin)
	for questionNum, question := range(qa){
		fmt.Printf("Problem #%v: %s = ",
		questionNum + 1, question[0])
		scanner.Scan()
		if scanner.Text() == question[1]{
			*numberCorrect++
			}
		}
	c <- true
	close(c)
}




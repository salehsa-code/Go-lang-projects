// Stefan Nilsson 2013-03-13

// This program implements an ELIZA-like oracle (en.wikipedia.org/wiki/ELIZA).
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	star   = "Pythia"
	venue  = "Delphi"
	prompt = "> "
)

func main() {
	fmt.Printf("Welcome to %s, the oracle at %s.\n", star, venue)
	fmt.Println("Your questions will be answered in due time.")

	oracle := Oracle()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		fmt.Printf("%s heard: %s\n", star, line)
		oracle <- line // The channel doesn't block.
	}
}

// Oracle returns a channel on which you can send your questions to the oracle.
// You may send as many questions as you like on this channel, it never blocks.
// The answers arrive on stdout, but only when the oracle so decides.
// The oracle also prints sporadic prophecies to stdout even without being asked.
func Oracle() chan<- string {
	questions := make(chan string)
	answers := make(chan string)
	//A go-routine that receives all questions,
	go func () {
		for question:= range questions{
			go prophecy(question, answers)
		}
		
	}()

	go func(){
		for{prophecy("", answers)}

	}()
	
	go func(){
		for{
			for _,letter := range strings.Split(<-answers, "") {
				time.Sleep(time.Millisecond * 25)
				fmt.Print(letter)
			}
			fmt.Print("\n", prompt)
		}
	}()

	// TODO: Answer questions.
	// TODO: Make prophecies.
	// TODO: Print answers.
	return questions
}

// This is the oracle's secret algorithm.
// It waits for a while and then sends a message on the answer channel.
// TODO: make it better.
func prophecy(question string, answer chan<- string) {
	// Keep them waiting. Pythia, the original oracle at Delphi,
	// only gave prophecies on the seventh day of each month.
	time.Sleep(time.Duration(20+rand.Intn(10)) * time.Second)

	// Find the longest word.
	longestWord := ""
	words := strings.Fields(question) // Fields extracts the words into a slice.
	for _, w := range words {
		if len(w) > len(longestWord) {
			longestWord = w
		}
	}

	// Cook up some pointless nonsense.
	nonsense := []string{
			
		// "The moon is dark.",
		// "The sun is bright.",







		"Cells",
		"Have you ever been in an institution? Cells",
		"When you're not performing your duties do they keep you in a little box? Cells",
		"Interlinked",
		"Do you dream about being interlinked... ?",
		"Did they teach you how to feel finger to finger? Interlinked",
		"Do you long for having your heart interlinked? Interlinked",
		"Do you feel that there's a part of you that's missing? Interlinked",
		"Dreadfully",
		"What’s it like to be filled with dread? Dreadfully",
		"Dreadfully distinct",
		"Within one stem.",
		"Against the dark", 
		"A tall white fountain played",
		"A blood black nothingness began to spin.",
		"What does it feel like to be part of the system. System.",
		"Is there anything in your body that wants to resist the system? System.",
		"Have they created you to be a part of the system? System.",
		"Do you long for having your heart interlinked? Interlinked.",
		"Do you have a heart? Within.",



		}
	answer <- longestWord + "... " + nonsense[rand.Intn(len(nonsense))]
}

func init() { // Functions called "init" are executed before the main function.
	// Use new pseudo random numbers every time.
	rand.Seed(time.Now().Unix())
}

package main

import "fmt"

func main() {
	fmt.Println("Let's simulate an election!")

	electoralVoteFile := "data/electoralVotes.csv"
	pollFile := "data/earlyPolls.csv"

	// now, read them in and store as maps
	electoralVotes := ReadElectoralVotes(electoralVoteFile)
	polls := ReadPollingData(pollFile)

	fmt.Println(electoralVotes)
	fmt.Println(polls)
}

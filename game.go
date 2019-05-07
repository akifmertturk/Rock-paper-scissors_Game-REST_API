package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var sessionId = 23
var round = 1
var i = 1 							  // to control the rounds
var selections = []string{"rock", "paper", "scissors"}
var scores = []int{0, 0}              //first one myScore, second one is yourScore
var allMySelections = []string{" "}   //stores my selections
var allYourSelections = []string{" "} //stores your selections

func determine(yourSelection string, mySelection string) int {
	switch yourSelection {
	case "rock":
		if mySelection == "paper" {
			return 1
		} else {
			return 0
		}
	case "paper":
		if mySelection == "scissors" {
			return 1
		} else {
			return 0
		}
	case "scissors":
		if mySelection == "rock" {
			return 1
		} else {
			return 0
		}
	default:
		return 2
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path[1:] {
	case "newGame":
		queryValues := r.URL.Query()  //I used .Query() to get the parameters
		s := queryValues.Get("round") //round is string here
		round, _ = strconv.Atoi(s)    //we need to convert round (string) to integer
		fmt.Fprintf(w, "\nNew Rock-Paper-Scissors game started\nSession ID = 23\n\n")
		fmt.Fprintln(w, "to play Rock: http://localhost:8080/play?choose=rock&id=23")
		fmt.Fprintln(w, "to play Paper: http://localhost:8080/play?choose=paper&id=23")
		fmt.Fprintln(w, "to play Scissors: http://localhost:8080/play?choose=scissors&id=23")
		fmt.Fprintf(w, "\n")
	case "play":
		if i <= round {
			queryValues := r.URL.Query()
			s := queryValues.Get("choose") //No need to convert types
			strSid := queryValues.Get("id")
			sid, _ := strconv.Atoi(strSid)
			if sid != sessionId {
				fmt.Fprintf(w, "Invalıd Session ID, Please retry by fixing Session ID\n")
			} else {
				rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
				mySelection := selections[rand.Intn(len(selections))]
				allMySelections = append(allMySelections, strings.Title(mySelection))
				allYourSelections = append(allYourSelections, strings.Title(s)) //strings.Title() makes words first letter uppercase
				message := ""
				if s == mySelection {
					message = "WE ARE TIED!!"
				} else {
					res := determine(s, mySelection)
					if res == 0 {
						message = "YOU WON THIS ROUND!!"
						scores[1]++
					} else {
						message = "I WON THIS ROUND!!"
						scores[0]++
					}
				}
				fmt.Fprintf(w, "\n-> ROUND %d\n\n", i)
				fmt.Fprintf(w, "me: %s\n", strings.ToUpper(mySelection))
				fmt.Fprintf(w, "you: %s\n", strings.ToUpper(s))
				fmt.Fprintf(w, "\n%s\n", message)
				if i < round {
					fmt.Fprintf(w, "\nThere is %d more round.\n", (round - i))
				} else {
					fmt.Fprintf(w, "\n->GAME COMPLETED\n")
					i = 1
					for i <= round {
						fmt.Fprintf(w, "Round %d: %s vs %s\n", i, allMySelections[i], allYourSelections[i])
						i++
					}
					fmt.Fprintf(w, "\n%d vs %d\n", scores[0], scores[1])
					if scores[0] < scores[1] {
						fmt.Fprintf(w, "YOU WON !!\n\n")
					} else if scores[0] > scores[1] {
						fmt.Fprintf(w, "I WON !!\n\n")
					} else {
						fmt.Fprintf(w, "WE ARE TIED !! DOSTLUK KAZANDI :) \n\n")
					}
				}
				i++
			}
		} else {
		}
	}

	//fmt.Fprintf(w, "Homepage end point, a: %s", i)

}

func handleRequests() {
	http.HandleFunc("/", homepage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {

	handleRequests()
}

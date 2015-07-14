/*
 * DAILY PROGRAMMER - 5/25/2015
 * TEXAS HOLD 'EM PART 1 (DEALER)
 * http://www.reddit.com/r/dailyprogrammer/comments/378h44/20150525_challenge_216_easy_texas_hold_em_1_of_3/
 * Author: Muhan Guclu
 * Description: Generates a deck of 52 unique cards and deals a game of Texas Hold 'Em given number of players
 */

package main

import (
    "strconv"
    "fmt"
    "math/rand"
    "time"
)

// get rank (Ace - King) of card
func getRank(n int) string {
    r := n%13   // remainder is rank in suit (52/4 = 13)

    switch {
    case r == 0:
        return "Ace"
    case r == 12:
        return "King"
    case r == 11:
        return "Queen"
    case r == 10:
        return "Jack"
    default:
        return strconv.Itoa(r)
    }
}
func getSuit (n int) string {
    s := n/13

    switch {
    case s == 1:
        return "Hearts"
    case s == 2:
        return "Diamonds"
    case s == 3:
        return "Clubs"
    default:
        return "Spades"
    }
}
// translates int to name
func readCard (n int) string {
    rank := getRank(n)
    suit := getSuit(n)
    return rank + " of " + suit
}

// cards stored as ints
func initDeck (d []int) []int {
    for i := range d {
        d[i] = i
    }
    return d
}

func getPlayers () int {
    fmt.Println("How many players (2-8)?")

    var numHands int
    minPlayers := 2
    maxPlayers := 8
    
    _, err := fmt.Scanf("%d", &numHands)

    // restart on error or number out of range
    if err != nil {
        fmt.Println("Enter a number between 2 and 8...")
        getPlayers()
    } else if numHands < minPlayers || numHands > maxPlayers {
        fmt.Println("Enter a number between 2 and 8...")
        getPlayers()
    } else {
        // no error
        return numHands
    }

    return numHands
}

func dealHands (numPlayers int , deck []int) ([]int, [][]int) {
    hands := make([][]int, numPlayers)   

    for i := range hands {
        handSize := 2
        hands[i] = make([]int, handSize)

        for j := range hands[i] {
            deck, hands[i][j] = drawCard(deck)
        }
    }
    return deck, hands
}

func burnCard (deck []int) []int {
    deck = deck[:len(deck)-1]
    return deck
}

func drawCard (deck []int) ([]int, int) {
    ret := deck[len(deck)-1]
    deck = burnCard(deck)
    return deck, ret
}

func burnAndDraw (deck []int) ([]int, int) { 
    deck = burnCard(deck)
    var newCard int
    deck, newCard = drawCard(deck) 
    return deck, newCard
}

func dealFlop (deck []int) ([]int, []int) {
    flop := make([]int, 3)
    for i := range flop {
        deck, flop[i] = burnAndDraw(deck)
    }
    return deck, flop
}

func shuffle (deck []int) []int {
    shuffled := make([]int, len(deck))
    perm := rand.Perm(len(deck))
    for i, v := range perm {
        shuffled[v] = deck[i]
    }
    return shuffled
}

func playGame () {
    numPlayers := getPlayers()     // user input int

    deck := make([]int, 52)
    deck = initDeck(deck)
    deck = shuffle(deck)

    deck, hands := dealHands(numPlayers, deck)     // 2d slice of player hands [][]int

    for i := range hands {
        handString := readCard(hands[i][0]) + ", " + readCard(hands[i][1])
        if i == 0 {
            fmt.Println("PLAYER " + handString)
        } else {
            fmt.Println("CPU", i, handString)
        }

    }
    fmt.Println("-----------")
    deck, flop := dealFlop(deck)
    fmt.Println("FLOP: " + readCard(flop[0]) + ", " + readCard(flop[1]) + ", " + readCard(flop[2]))

    deck, river := burnAndDraw(deck)
    fmt.Println("RIVER: " + readCard(river))

    deck, turn := burnAndDraw(deck)
    fmt.Println("TURN: " + readCard(turn))
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    playGame()
}

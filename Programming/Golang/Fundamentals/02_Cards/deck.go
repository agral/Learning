package main

import (
    "fmt"
    "io/ioutil"
    "math/rand"
    "os"
    "strings"
    "time"
)

type deck []string

func newDeck() deck {
    cards := deck{}
    cardSuits := []string {"Spades", "Clubs", "Diamonds", "Hearts"}
    cardValues := []string {"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight",
        "Nine", "Ten", "Jack", "Queen", "King"}
    for _, suit := range cardSuits {
        for _, value := range cardValues {
            cards = append(cards, value + " of " + suit)
        }
    }

    return cards
}

func deal(d deck, handSize int) (deck, deck) {
    return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
    return strings.Join([]string(d), ",")
}

func (d deck) print() {
    for i, card := range d {
        fmt.Println(i, card)
    }
}

func (d deck) saveToFile(filename string) error {
    return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
    bs, err := ioutil.ReadFile(filename)
    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(1)
    }
    s := strings.Split(string(bs), ",")
    return deck(s)
}

func (d deck) shuffle() {
    randSource := rand.NewSource(time.Now().UnixNano())
    r := rand.New(randSource)
    for index := range d {
        newPosition := r.Intn(len(d) - 1)
        d[index], d[newPosition] = d[newPosition], d[index]
    }
}

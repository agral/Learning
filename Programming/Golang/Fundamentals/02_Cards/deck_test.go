package main

import "testing"
import "os"

func TestNewDeck(t *testing.T) {
    d := newDeck()
    if len(d) != 52 {
        t.Errorf("Expected deck length of 52, but got %v", len(d))
    }
    if d[0] != "Ace of Spades" {
        t.Errorf("Expected the first card in deck to be 'Ace of Spades', got %v instead", d[0])
    }
    if d[len(d) - 1] != "King of Hearts" {
        t.Errorf("Expected the last card in deck to be 'King of Hearts', got %v instead", d[len(d) - 1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    os.Remove("_decktesting")
    deck := newDeck()
    deck.saveToFile("_decktesting")

    loadedDeck := newDeckFromFile("_decktesting")

    if (len(loadedDeck) != 52) {
        t.Errorf("Expected deck length of 52, but got %v", len(deck))
    }

    os.Remove("_decktesting")
}

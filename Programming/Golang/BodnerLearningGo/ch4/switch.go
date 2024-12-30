package main

import "fmt"

func standardSwitch() {
	animals := []string{"aardvark", "ox", "eel", "sheep", "giraffe", "toucan"}
	for _, word := range animals {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Printf("%s is a short word.\n", word)
		case 5:
			fmt.Printf("%s is exactly the right length: %d\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf("%s is a long word.\n", word)
		}
	}
	// Note: unlike C, there's no fall-through!
}

func blankSwitch() {
	greetings := []string{"hi", "salutations", "hello"}
	for _, word := range greetings {
		switch wordLen := len(word); { // note the blank after `;`, this is a blank switch.
		case wordLen < 5:
			fmt.Printf("%s is a short word.\n", word)
		case wordLen > 10:
			fmt.Printf("%s is a long word.\n", word)
		default:
			fmt.Printf("%s is exactly the right length: %d\n", word, wordLen)
		}
	}

}

func main() {
	standardSwitch()
	blankSwitch()
}

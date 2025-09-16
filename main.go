package main

import (
	"fmt"
	"strings"
)

func chriffrementCesar(word string, shift int) string {
	var b strings.Builder
	b.Grow(len(word))

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			// c-'a' pour revenir à un unicode entre 0 et 25
			// rajoute = 'a' pour le remettre à la valeur de l'unicode
			b.WriteRune(((c-'a')+rune(shift))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			// c-'A' pour revenir à un unicode entre 0 et 25
			// rajoute = 'A' pour le remettre à la valeur de l'unicode
			b.WriteRune(((c-'A')+rune(shift))%26 + 'A')
		} else {
			b.WriteRune(c)
		}
	}
	return b.String()
}

func dechriffrementCesar(word string, shift int) string {
	var wordCipher string
	wordCipher = chriffrementCesar(word, 26-shift)
	return wordCipher
}

func main() {
	fmt.Println(chriffrementCesar("Hello World!", 3))
	fmt.Println(dechriffrementCesar("Khoor Zruog!", 3))
}

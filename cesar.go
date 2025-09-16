package main

import (
	"strings"
)

func cesarCypher(word string, shift int) string {
	// Créer un builder pour créer une chaine de caractères avec une longueur initiale
	var b strings.Builder
	b.Grow(len(word))

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			// c-'a' pour revenir à un unicode entre 0 et 25
			// trasforme le décalage en rune
			// faire un modulo 26 pour renvoyer à 0
			// rajoute + 'a' pour le remettre à la valeur de l'unicode
			b.WriteRune(((c-'a')+rune(shift))%26 + 'a')
		} else if c >= 'A' && c <= 'Z' {
			// c-'A' pour revenir à un unicode entre 0 et 25
			// trasforme le décalage en rune
			// faire un modulo 26 pour renvoyer à 0
			// rajoute = 'A' pour le remettre à la valeur de l'unicode
			b.WriteRune(((c-'A')+rune(shift))%26 + 'A')
		} else {
			// Si c'est pas une lettre on renvoie le caratère
			b.WriteRune(c)
		}
	}
	return b.String()
}

func cesarDecypher(word string, shift int) string {
	var wordCipher string
	// Pour déchiffrer on décale de 26 - shift
	wordCipher = cesarCypher(word, 26-shift)
	return wordCipher
}

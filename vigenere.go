package main

import (
	"strings"
)

func vigenereCypher(word string, key string) string {
	var b strings.Builder

	// j permettra d'itérer seulement sur les lettres de la clé et ne prendra pas en compte les espaces ou autres caractères
	j := 0 // compteur de clé

	for _, c := range word {
		if c >= 'a' && c <= 'z' {
			// récupérer la lettre à l'index i modulo la longueur de la clé
			k := key[j%len(key)]
			// récupérer le décalage de la lettre de la clé
			shift := k - 'A'
			b.WriteString(cesarCypher(string(c), int(shift)))
			j++ // n'incrémenter j que si on a utilisé une lettre
		} else if c >= 'A' && c <= 'Z' {
			// récupérer la lettre à l'index i modulo la longueur de la clé
			k := key[j%len(key)]
			// récupérer le décalage de la lettre de la clé
			shift := k - 'A'
			b.WriteString(cesarCypher(string(c), int(shift)))
			j++ // n'incrémenter j que si on a utilisé une lettre
		} else {
			// Si c'est pas une lettre on renvoie le caratère
			b.WriteRune(c)
		}
	}
	return b.String()
}

func vigenereDecypher(cipher string, key string) string {
	var b strings.Builder

	// j permettra d'itérer seulement sur les lettres de la clé et ne prendra pas en compte les espaces ou autres caractères
	j := 0 // compteur de clé

	for _, c := range cipher {
		if c >= 'a' && c <= 'z' {
			// récupérer la lettre à l'index i modulo la longueur de la clé
			k := key[j%len(key)]
			// récupérer le décalage de la lettre de la clé
			shift := k - 'A'
			b.WriteString(cesarDecypher(string(c), int(shift)))
			j++ // n'incrémenter j que si on a utilisé une lettre
		} else if c >= 'A' && c <= 'Z' {
			// récupérer la lettre à l'index i modulo la longueur de la clé
			k := key[j%len(key)]
			// récupérer le décalage de la lettre de la clé
			shift := k - 'A'
			b.WriteString(cesarDecypher(string(c), int(shift)))
			j++ // n'incrémenter j que si on a utilisé une lettre
		} else {
			// Si c'est pas une lettre on renvoie le caratère
			b.WriteRune(c)
		}
	}
	return b.String()
}

package main

import (
	"math/rand"
	"strings"
)

// dictionnaire homophonique : chaque lettre a plusieurs codes possibles
var homophonique = map[rune][]string{
	'A': {"11", "45", "78"},
	'B': {"12", "46"},
	'C': {"13", "47"},
	'D': {"14", "48"},
	'E': {"15", "49", "79", "90"},
	'F': {"16", "50"},
	'G': {"17", "51"},
	'H': {"18", "52"},
	'I': {"19", "53"},
	'J': {"20", "54"},
	'K': {"21", "55"},
	'L': {"22", "56"},
	'M': {"23", "57"},
	'N': {"24", "58"},
	'O': {"25", "59"},
	'P': {"26", "60"},
	'Q': {"27", "61"},
	'R': {"28", "62"},
	'S': {"29", "63"},
	'T': {"30", "64"},
	'U': {"31", "65"},
	'V': {"32", "66"},
	'W': {"33", "67"},
	'X': {"34", "68"},
	'Y': {"35", "69"},
	'Z': {"36", "70"},
}

func homophoniqueCypher(word string) string {
	var b strings.Builder

	for _, c := range strings.ToUpper(word) {
		// récupérer les codes associés à la lettre
		if codes, ok := homophonique[c]; ok {
			// choisir un code aléatoire parmi ceux disponibles de la lettre
			code := codes[rand.Intn(len(codes))]
			// ajouter le code au builder avec un espace pour séparer les codes
			b.WriteString(code + " ")
		} else {
			// si ce n'est pas une lettre, on laisse tel quel
			b.WriteRune(c)
		}
	}

	return strings.TrimSpace(b.String())
}

func homophoniqueDecypher(cipher string) string {
	var b strings.Builder

	// Separer les codes par les espaces
	codes := strings.Split(cipher, " ")

	// Pour chaque code, chercher la lettre correspondante
	for _, code := range codes {
		found := false
		// Si le code est une chaine vide, c'est un espace
		if code == "" {
			b.WriteRune(' ')
			continue
		}
		for letter, letterCode := range homophonique {
			// Vérifier si le code fait partie des codes de la lettre
			for _, c := range letterCode {
				if code == c {
					b.WriteString(string(letter))
					found = true
					break
				}
			}

			// Si on a trouvé la lettre, arreter la recherche
			if found {
				break
			}
		}
		// Si le code n'a pas été trouvé, on le laisse tel quel
		if !found {
			b.WriteString(code)
		}
	}
	return b.String()
}

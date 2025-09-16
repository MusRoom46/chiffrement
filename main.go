package main

func main() {
	println("Cesar Cipher :")
	println(cesarCypher("Hello World!", 3))
	println(cesarDecypher("Khoor Zruog!", 3))
	println("")
	println("Homophonique Cipher :")
	println(homophoniqueCypher("Hello World!"))
	println(homophoniqueDecypher("18 49 22 56 59  67 25 62 56 14 !"))
	println("")
	println("Vigenere Cipher :")
	println(vigenereCypher("Hello World!", "KEY"))
	println(vigenereDecypher("Rijvs Uyvjn!", "KEY"))
}

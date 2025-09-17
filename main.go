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
	println("")

	println("Hill Cipher :")
	// Test 1 : Multiplication matrice × vecteur
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	vector1 := []int{1, 2, 1}
	println("Test 1 : Multiplication matrice × vecteur")
	println("Matrice :", matrix1)
	println("Vecteur :", vector1)
	println("Résultat :", multiplyMatrixVector(matrix1, vector1))
	println()

	// Test 2 : Addition de matrices
	matrix2 := [][]int{
		{1, 2},
		{3, 4},
	}
	matrix3 := [][]int{
		{5, 6},
		{7, 8},
	}
	println("Test 2 : Somme de matrices")
	println("Matrice A :", matrix2)
	println("Matrice B :", matrix3)
	println("Résultat :", addMatrices(matrix2, matrix3))
	println()

	// Test 3 : Déterminant et inversibilité modulo n
	matrix4 := [][]int{
		{2, 3},
		{1, 4},
	}
	println("Test 3 : Déterminant et inversibilité modulo n")
	println("Matrice :", matrix4)
	det := determinant2x2(matrix4)
	println("Déterminant :", det)
	println("Inversible mod 5 :", isInvertibleModuloN(matrix4, 5))
	println("Inversible mod 6 :", isInvertibleModuloN(matrix4, 6))
}

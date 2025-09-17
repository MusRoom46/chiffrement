package main

// -------------------------
// 1. Matrice × Vecteur
// -------------------------

// multiplyMatrixVector prend une matrice ([][]int) et un vecteur ([]int)
// et retourne un nouveau vecteur ([]int) qui est le résultat du produit.
// Hypothèse :
// - la matrice est de taille m × n
// - le vecteur est de taille n
// - le résultat est un vecteur de taille m
func multiplyMatrixVector(matrix [][]int, vector []int) []int {
	// Vérification : le nombre de colonnes de la matrice doit être égal
	// à la longueur du vecteur sinon le produit n’a pas de sens
	cols := len(matrix[0])
	if cols != len(vector) {
		panic("La taille du vecteur ne correspond pas au nombre de colonnes de la matrice")
	}

	// On crée le résultat (vecteur) avec une longueur égale au nombre de lignes de la matrice
	result := make([]int, len(matrix))

	// On parcourt chaque ligne de la matrice
	for i := 0; i < len(matrix); i++ {
		sum := 0
		// On calcule le produit scalaire entre la ligne de la matrice et le vecteur
		for j := 0; j < cols; j++ {
			sum += matrix[i][j] * vector[j]
		}
		result[i] = sum
	}
	return result
}

// -------------------------
// 2. Somme de matrices
// -------------------------

// addMatrices prend une liste de matrices de même taille
// et retourne leur somme.
// Hypothèse :
// - toutes les matrices ont la même taille (m × n)
func addMatrices(matrices ...[][]int) [][]int {
	if len(matrices) == 0 {
		panic("Aucune matrice fournie")
	}

	rows := len(matrices[0])
	cols := len(matrices[0][0])

	// Création d’une matrice résultat remplie de zéros
	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	// Addition de toutes les matrices
	for _, m := range matrices {
		if len(m) != rows || len(m[0]) != cols {
			panic("Toutes les matrices doivent avoir la même taille")
		}
		for i := 0; i < rows; i++ {
			for j := 0; j < cols; j++ {
				result[i][j] += m[i][j]
			}
		}
	}
	return result
}

// -------------------------
// 3. Déterminant 2×2 et inversibilité modulo n
// -------------------------

// determinant2x2 calcule le déterminant d’une matrice 2×2
func determinant2x2(matrix [][]int) int {
	if len(matrix) != 2 || len(matrix[0]) != 2 {
		panic("La matrice doit être de taille 2×2")
	}
	return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0]
}

// gcd calcule le plus grand commun diviseur (algorithme d’Euclide)
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// isInvertibleModuloN vérifie si une matrice 2×2 est inversible modulo n
// Une matrice est inversible modulo n si et seulement si
// gcd(det, n) == 1
func isInvertibleModuloN(matrix [][]int, n int) bool {
	det := determinant2x2(matrix)
	return gcd(det, n) == 1
}

package main

import (
	"fmt"

	"github.com/01-edu/z01"
)

var (
	initVal int  = 0        // Valeur initiale pour le pointeur limit
	limit   *int = &initVal // Pointeur utilisé pour stocker la ligne limite dans la recherche
	count   int  = 1        // Compteur pour numéroter les solutions
)

func main() {
	EightQueens()
}

// EightQueens démarre le processus de résolution du problème des huit reines
func EightQueens() {
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			oneVal := ""
			*limit = row
			solver(oneVal, col, row)
		}
	}
}

// solver résout récursivement le problème des huit reines en plaçant les reines sur le plateau
func solver(placedQueens string, col int, row int) {
	// Création d'un tableau 8x8 représentant l'échiquier et remplissage avec les reines déjà placées
	board := [8][8]int{}
	if len(placedQueens) > 0 {
		for i, c := range strToTab(placedQueens) {
			if i < len(placedQueens) {
				board[i][c] = 1
			}
		}
	}
	// Placer la reine actuelle sur l'échiquier
	board[row][col] = 1

	// Mise à jour de la chaîne qui contient les positions des reines
	placedQueens += string(rune(col + '0'))

	// Si 8 reines sont placées et la solution est valide, afficher la solution
	if len(placedQueens) == 8 && checkResult(placedQueens) {
		fmt.Printf("Solution %d : ", count)
		for _, fi := range placedQueens {
			z01.PrintRune(rune(fi + 1))
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
		count++
	}

	// Condition d'arrêt : Si la ligne actuelle est la dernière, arrêter la récursion
	if row < 7 {
		// Si on est revenu à la première ligne, ne pas continuer
		if row == *limit && len(placedQueens) > 1 {
			return
		} else {
			// Trouver toutes les possibilités pour la ligne suivante
			possibilities := ""
			for i := 0; i < 8; i++ {
				if findPossibility(placedQueens, row+1, i) != -1 {
					possibilities += string(rune(i + '0'))
				}
			}
			// Passer à la ligne suivante
			row += 1

			// Pour chaque possibilité, appeler récursivement solver
			for _, p := range possibilities {
				solver(placedQueens, int(p-'0'), row)
			}
		}
	}
}

// findPossibility vérifie si une reine peut être placée à la position donnée
func findPossibility(placedQueens string, row int, col int) int {
	board := [8][8]int{}
	if len(placedQueens) > 0 {
		for i, c := range strToTab(placedQueens) {
			if i < len(placedQueens) {
				board[i][c] = 1
			}
		}
	}

	// Vérifie s'il y a déjà une reine dans la même ligne ou colonne
	for _, k := range board[row] {
		if k == 1 {
			return -1
		}
	}
	for _, lignes := range board {
		if lignes[col] == 1 {
			return -1
		}
	}

	// Vérifie les diagonales
	for diagCol, diagRow := col, row; diagCol < 8 && diagRow < 8; diagCol, diagRow = diagCol+1, diagRow+1 {
		if board[diagRow][diagCol] == 1 {
			return -1
		}
	}
	for diagCol, diagRow := col, row; diagCol >= 0 && diagRow >= 0; diagCol, diagRow = diagCol-1, diagRow-1 {
		if board[diagRow][diagCol] == 1 {
			return -1
		}
	}
	for diagCol, diagRow := col, row; diagCol >= 0 && diagRow < 8; diagCol, diagRow = diagCol-1, diagRow+1 {
		if board[diagRow][diagCol] == 1 {
			return -1
		}
	}

	return col
}

// checkResult vérifie si la solution actuelle est valide
func checkResult(placedQueens string) bool {
	for row, col := range placedQueens {
		if !resultCheck(placedQueens, row, int(col-'0')) {
			return false
		}
	}
	return true
}

// strToTab convertit une chaîne de caractères en tableau d'entiers
func strToTab(s string) [8]int {
	runeS := []rune(s)
	out := [8]int{}
	for i := 0; i < len(runeS); i++ {
		out[i] = int(runeS[i] - '0')
	}
	return out
}

// resultCheck vérifie si une reine peut être placée à une position donnée sans conflit
func resultCheck(placedQueens string, row int, col int) bool {
	board := [8][8]int{}
	if len(placedQueens) > 0 {
		outTab := strToTab(placedQueens)
		for i, c := range outTab {
			if i < len(placedQueens) {
				board[i][c] = 1
			}
		}
	}

	// Vérifie les conflits dans la même ligne ou colonne
	for i, k := range board[row] {
		if k == 1 && i != col {
			return false
		}
	}
	for i, lignes := range board {
		if lignes[col] == 1 && i != row {
			return false
		}
	}

	// Vérifie les conflits dans les diagonales
	for diagCol, diagRow := col, row; diagCol < 8 && diagRow < 8; diagCol, diagRow = diagCol+1, diagRow+1 {
		if board[diagRow][diagCol] == 1 && diagRow != row && diagCol != col {
			return false
		}
	}
	for diagCol, diagRow := col, row; diagCol >= 0 && diagRow >= 0; diagCol, diagRow = diagCol-1, diagRow-1 {
		if board[diagRow][diagCol] == 1 && diagRow != row && diagCol != col {
			return false
		}
	}
	for diagCol, diagRow := col, row; diagCol >= 0 && diagRow < 8; diagCol, diagRow = diagCol-1, diagRow+1 {
		if board[diagRow][diagCol] == 1 && diagRow != row && diagCol != col {
			return false
		}
	}

	return true
}

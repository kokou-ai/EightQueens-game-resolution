# 🏰 Eight Queens Solver in Go

Welcome to the **Eight Queens Solver**! This project provides a Go implementation of the classic Eight Queens puzzle, which is a famous problem in the world of algorithms and computer science.

## 📜 Overview

The Eight Queens puzzle challenges you to place eight queens on an 8x8 chessboard such that no two queens can attack each other. This means no two queens can share the same row, column, or diagonal.

This repository contains a Go program that finds and displays all possible solutions to this problem.

## 🚀 Getting Started

### Prerequisites

- Go (Golang) installed on your machine. You can download it from [here](https://golang.org/dl/).

### Installation

### 1. Clone the repository:

   ```bash```
   git clone https://github.com/yourusername/eight-queens-go.git
   cd eight-queens-go 

### 2. Build the project (optional, as Go can run the source file directly):

```bash```
go build eightqueens.go

Usage

To run the program, use the following command:

```bash```

go run eightqueens.go

The program will display all the solutions to the Eight Queens problem in the format:

Solution 1 : 1 5 8 6 3 7 2 4
Solution 2 : 1 6 8 3 7 4 2 5
...


Each number represents the column position of the queen in the respective row.
## 🛠️ How It Works

The solver uses a recursive backtracking approach to explore all possible configurations of placing queens on the board. The key functions include:

    EightQueens: Initializes the search process by iterating through each position on the board.
    solver: Recursively attempts to place queens on the board and checks for valid configurations.
    findPossibility: Determines if a queen can be safely placed at a given position.
    checkResult: Validates if the current configuration meets the conditions of the puzzle.

## 📈 Output

The program outputs the positions of queens for each valid solution, prefixed with the solution number, for easy reference.
🤝 Contributing

Contributions are welcome! If you have ideas for improving the algorithm or additional features, feel free to fork the repository and submit a pull request.

    Fork the project
    Create your feature branch (git checkout -b feature/AmazingFeature)
    Commit your changes (git commit -m 'Add some AmazingFeature')
    Push to the branch (git push origin feature/AmazingFeature)
    Open a pull request

## 📝 License

This project is licensed under the MIT License - see the LICENSE file for details.
📧 Contact

## If you have any questions, feel free to reach out via LinkedIn or GitHub.
You can send me email there: sitsopekokou@gmail.com

Happy coding! 😊

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

//
// ANSI color codes for terminal output
//
const (
	Green  = "\u001B[32m" // correct letter, correct position
	Yellow = "\u001B[33m" // correct letter, wrong position
	White  = "\u001B[0m"  // incorrect letter
	Reset  = "\u001B[0m"  // reset terminal color
)

//
// ENTRY POINT
//
func main() {

	// Load word list from file
	words, err := loadWords("wordle-words.txt")
	if err != nil || len(words) == 0 {
		fmt.Println("Word list not found or empty.")
		return
	}

	// Choose secret word based on command-line argument
	secret := chooseSecretWord(words)

	// Create scanner for all user input (required by assignment)
	scanner := bufio.NewScanner(os.Stdin)

	// Handle login
	username, ok := askUsername(scanner)
	if !ok {
		return
	}

	// Run the game loop
	win, attemptsUsed := playGame(scanner, secret)

	// Save results
	saveGameStats(username, secret, attemptsUsed, win)

	// Ask if user wants to see stats
	handleStatsView(scanner, username)
}

//
// Reads username from stdin
//
func askUsername(scanner *bufio.Scanner) (string, bool) {
	fmt.Println("Enter your username:")
	if !scanner.Scan() { // handles Ctrl+D
		return "", false
	}
	return strings.TrimSpace(scanner.Text()), true
}

//
// Determines the secret word using CLI argument
//
func chooseSecretWord(words []string) string {
	index := parseIndex(os.Args, len(words))
	return strings.ToUpper(words[index])
}

//
// Main Wordle game loop
//
func playGame(scanner *bufio.Scanner, secret string) (bool, int) {

	fmt.Println("Welcome to Wordle! Guess the 5-letter word.")

	attemptsRemaining := 6
	attemptsUsed := 0
	win := false

	// Track letters that are confirmed NOT in the word
	usedLetters := make(map[rune]bool)

	for attemptsRemaining > 0 {

		guess, ok := readGuess(scanner)
		if !ok {
			return false, attemptsUsed
		}

		attemptsUsed++

		colors := evaluateGuess(guess, secret)

		printFeedback(guess, colors)
		updateUsedLetters(guess, colors, usedLetters)

		fmt.Print("Remaining letters: ")
		fmt.Println(remainingLetters(usedLetters))

		attemptsRemaining--
		fmt.Println("Attempts remaining: ", attemptsRemaining)

		if guess == secret {
			win = true
			break
		}
	}

	if !win {
		fmt.Println("The word was:", secret)
	}

	return win, attemptsUsed
}

//
// Reads and validates a guess
//
func readGuess(scanner *bufio.Scanner) (string, bool) {
	fmt.Println("Enter your guess:")
	if !scanner.Scan() {
		return "", false
	}

	guess := strings.ToUpper(strings.TrimSpace(scanner.Text()))
	if len(guess) != 5 {
		return readGuess(scanner) // retry until valid
	}

	return guess, true
}

//
// Prints colored feedback for a guess
//
func printFeedback(guess string, colors []string) {
	fmt.Print("Feedback: ")
	for i, ch := range guess {
		fmt.Print(colors[i] + string(ch) + Reset)
	}
	fmt.Println()
}

//
// Updates map of incorrect letters
//
func updateUsedLetters(guess string, colors []string, used map[rune]bool) {
	for i, ch := range guess {
		if colors[i] == White {
			used[ch] = true
		}
	}
}

//
// Wordle letter evaluation logic
//
func evaluateGuess(guess, secret string) []string {

	colors := make([]string, 5)
	secretUsed := make([]bool, 5)

	// First pass: exact matches (green)
	for i := 0; i < 5; i++ {
		if guess[i] == secret[i] {
			colors[i] = Green
			secretUsed[i] = true
		}
	}

	// Second pass: misplaced letters (yellow) or incorrect (white)
	for i := 0; i < 5; i++ {
		if colors[i] == Green {
			continue
		}
		for j := 0; j < 5; j++ {
			if !secretUsed[j] && guess[i] == secret[j] {
				colors[i] = Yellow
				secretUsed[j] = true
				break
			}
		}
		if colors[i] == "" {
			colors[i] = White
		}
	}

	return colors
}

//
// Shows stats prompt and displays stats if requested
//
func handleStatsView(scanner *bufio.Scanner, username string) {

	fmt.Println("Do you want to see your stats? (yes/no):")
	if !scanner.Scan() {
		return
	}

	if strings.ToLower(scanner.Text()) == "yes" {
		showStats("stats.csv", username)
		fmt.Println("Press Enter to exit...")
		scanner.Scan()
	}
}

//
// Loads valid 5-letter words from file
//
func loadWords(path string) ([]string, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if len(word) == 5 {
			words = append(words, strings.ToLower(word))
		}
	}

	return words, nil
}

//
// Parses command-line argument safely
//
func parseIndex(args []string, max int) int {

	if len(args) < 2 {
		return 0
	}

	i, err := strconv.Atoi(args[1])
	if err != nil || i < 0 || i >= max {
		return 0
	}

	return i
}

//
// Returns remaining unused letters (A–Z)
//
func remainingLetters(used map[rune]bool) string {

	var letters []string
	for ch := 'A'; ch <= 'Z'; ch++ {
		if !used[ch] {
			letters = append(letters, string(ch))
		}
	}

	sort.Strings(letters)
	return strings.Join(letters, " ")
}

//
// Writes one game result to stats.csv
//
func saveGameStats(username, secret string, attempts int, win bool) {

	file, err := os.OpenFile("stats.csv", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	result := "loss"
	if win {
		result = "win"
	}

	writer.Write([]string{
		username,
		secret,
		strconv.Itoa(attempts),
		result,
	})
}

//
// Displays statistics for a given user
//
func showStats(path, username string) {

	file, err := os.Open(filepath.Clean(path))
	if err != nil {
		fmt.Println("No stats available.")
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return
	}

	games := 0
	wins := 0
	totalAttempts := 0

	for _, row := range records {
		if row[0] != username {
			continue
		}
		games++
		if row[3] == "win" {
			wins++
		}
		a, _ := strconv.Atoi(row[2])
		totalAttempts += a
	}

	avg := 0.0
	if games > 0 {
		avg = float64(totalAttempts) / float64(games)
	}

	fmt.Println("Stats for", username+":")
	fmt.Println("Games played:", games)
	fmt.Println("Games won:", wins)
	fmt.Printf("Average attempts per game: %.2f\n", avg)
}

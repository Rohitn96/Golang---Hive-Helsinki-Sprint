package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const(
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
	Reset = "\033[0m"
	Blue = "\033[34m"
)

func ValidateArgs() (string, bool) {
	if len(os.Args) !=2 {
		fmt.Println(Red, "The program requires exactly one command-line argument. Please do it again!", Reset)
		return "", false
	}
	collection := os.Args[1]
	if collection == "help" {
		return "", false
	}
	return collection, true
}

func PrintHelp() {
	fmt.Println(Yellow, "Usage: ./notestool [TAG]", Reset)
	fmt.Println(Yellow, "Write the tool name 'notestool' along with the TAG 'collection'", Reset)
}

func LoadOrCreateFile(filename string) {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
		fmt.Println(Red, "Cannot create file", Reset)
		}
		file.Close()
	}
}

func MenuLoop(filename string) {
	for {
		fmt.Println("\nSelect operation:")
		fmt.Println("1. Show notes.")
		fmt.Println("2. Add a note.")
		fmt.Println("3. Delete a note.")
		fmt.Println("4. Exit.")

		choice := ReadUserChoice()
		exit := HandleChoice(choice, filename)
		if exit {
			fmt.Println(Blue, "Goodbye!", Reset)
			break
		}
	}
}


func ReadUserChoice() int {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil {
		return -1
	}
	return choice
}

func HandleChoice(choice int, filename string) bool {
	switch choice {
	case 1:
		fmt.Print("\033[H\033[2J")
		ShowNotes(filename)
	case 2:
		fmt.Print("\033[H\033[2J")
		AddNote(filename)
	case 3:
		fmt.Print("\033[H\033[2J")
		DeleteNote(filename)
	case 4:
		return true
	default:
		fmt.Println(Red, "Invalid choice. Please enter a number between 1 and 4.", Reset)
	}
	return false
}

func ReadNotes(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var notes []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		notes = append(notes, scanner.Text())
	}
	return notes
}

func ShowNotes(filename string) {
	notes := ReadNotes(filename)

	if len(notes) == 0 {
		fmt.Println(Red ,"No notes found." , Reset)
		return
	}

	for i, note := range notes {
		fmt.Printf(Green + "%03d - %s\n" + Reset, i+1, note)
	}
}

func AddNote(filename string) {
	fmt.Println(Green, "\nEnter the note text:", Reset)

	reader := bufio.NewReader(os.Stdin)
	note, _ := reader.ReadString('\n')
	note = strings.TrimSpace(note)
	
	if len(note) == 0 {
		fmt.Println(Red, "The note cannot be blank..!", Reset)
		return
	}

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0664)
	if err != nil {
		fmt.Println(Red, "Error writing to file", Reset)
		return
	}
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	
	writer := bufio.NewWriter(file)
	writer.WriteString(fmt.Sprintf("[%s] %s", timestamp, note + "\n"))
	writer.Flush()
	
	fmt.Println(Green, "Note added.", Reset)
}

func DeleteNote(filename string) {
	notes := ReadNotes(filename)

	if len(notes) == 0  {
		fmt.Println(Red, "No notes to delete.",Reset)
		return
	}

	for i, note := range notes {  
	fmt.Printf(Green + "%03d - %s\n" + Reset, i+1, note)
	}

	fmt.Println(Yellow, "Enter number of note to remove or 0 to cancel :", Reset)
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, err := strconv.Atoi(input)
	if err != nil || index > len(notes) {
		fmt.Println(Red, "Invalid number.", Reset)
		return
	}
	if index == 0 {
		fmt.Println(Red, "Deleting canceled.", Reset)
		return
	}

	notes = append(notes[:index-1], notes[index:]...)
	WriteNotesToFile(filename, notes)
	fmt.Println(Green, "Note deleted.", Reset)
}

func WriteNotesToFile(filename string, notes []string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(Red, "Error rewriting file", Reset)
		return
	}
	defer file.Close()

	for _, note := range notes {
		file.WriteString(note + "\n")
	}
}

func main() {
	collection, ok := ValidateArgs()
	if !ok {
		PrintHelp()
		return
	}

	filename := collection
	LoadOrCreateFile(filename)

	fmt.Println(Blue, "Welcome to the notes tool!" , Reset)
	MenuLoop(filename)
}

# Notes Tool 

Notes Tool is a command-line application written in Go that allows users to manage a simple collection of notes stored in a text file.  
The application provides an interactive menu to view, add, and delete notes until the user chooses to exit.

---

## Tool Functionality

When the application is started with a valid argument, it:

1. Greets the user
2. Displays an interactive menu
3. Allows the user to perform the following operations:
   - Display notes from the collection
   	
   		Shows all notes stored in the selected collection
   - Add a note to the collection
   	
   		Prompts the user to enter a new note and saves it to the collection
   - Remove a note from the collection
   	
   		Displays all notes with numbers and allows the user to remove a selected note
   - Exit the program
4. Returns the user to the menu after each operation until they choose to exit

---

## Data Storage

The program requires **exactly one command-line argument**, which represents the file used to store notes. It takes the notes from the user and stores it in the '.txt' file that the user created while running the program.
Eg. If user runs 'go run notestool.go rough'. The program will create a file called rough.txt and store all the notes in that file. 
   - Each note is stored on a separate line in the file.
   - If the file does not exist, it is created automatically.
   - Notes persist between different runs of the application.

---

## Run the program

```bash
go run notestool.go 'Collection Name'

---

## Usage

~/notes % go run notestool.go rough
 Welcome to the notes tool!

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
Tuan
 Note added.

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
1
001 - Rohit
002 - Tuan

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
2

Enter the note text:
Festus
 Note added.

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
3
001 - Rohit
002 - Tuan
003 - Festus
 Enter number to delete:
3
 Note deleted.

Select operation:
1. Show notes.
2. Add a note.
3. Delete a note.
4. Exit.
4
 Goodbye!

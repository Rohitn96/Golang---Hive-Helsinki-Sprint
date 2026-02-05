package main

import "bufio"
import 	"os"
import 	"fmt"
import "strings"
import "strconv"


func main(){

		var toEncrypt bool
		var message string
		var encoding string
		var result string

		toEncrypt ,message, encoding  = GetInput()


			if toEncrypt {
				switch encoding {
					case "Rot13":
						result = Encrypt_rot13(message)

					case "Reverse":
						result = Encrypt_reverse(message)
					case "Caesar3":
						result = Encrypt_caesar(message)
					default:
						fmt.Println("Could not Encode")
				}
			fmt.Printf("Encrypted message using %v:\n%v\n", encoding, result)

			} else {
				switch encoding {
					case "Rot13":
						result = Decrypt_rot13(message)
					case "Reverse":
						result = Decrypt_reverse(message)
					case "Caesar3":
						result = Decrypt_caesar(message)
					default:
						fmt.Println("Could not Decode")
				}
			fmt.Printf("Decrypted message using %v:\n%v\n", encoding, result)

			}
}

func GetInput() (toEncrypt bool, encoding string, message string) {
	reader := bufio.NewReader(os.Stdin)


	// firstMenu
		for{
				fmt.Println("Welcome to our cypher!!!\n")
				fmt.Println("Select operation (1/2):\n1.Encrypt\n2.Decrypt")

				/*

				Read line and and new line

				we convert the string to int to check in the coditional --> strconv.Atoi()
				this would return two values the (int and error )

				this mean ( _ ) we dont use that value so we dont take it

				*/
				input, _ :=  reader.ReadString('\n')

				// this delete extra white space --> strings.TrimSpace()
				input = strings.TrimSpace(input)


				option, err := strconv.Atoi(input)

				// here we check if there is error (example the error that would return is that input is not int)
				if err != nil{
					fmt.Println("choose a number")
					continue
				}

				/*
				1 -> encrypt -> true
				2 -> encrypt -> false
				*/

				if option == 1 {
					toEncrypt = true
					// this is because we need to say the menu loop to break to start the next one\
				} else if option == 2 {
					toEncrypt = false
				} else {
					fmt.Println("That option is not valid.")
					// here we place continue so the loop menu one would keep asking for the right input
					continue
				}
				break

		}
		//cypher menu loop
		for{
				fmt.Println("Select cypher (1/2):\n1.Rot13.\n2.Reverse.\n3.Caesar")



				input_cypher, _ :=  reader.ReadString('\n')



				input_cypher = strings.TrimSpace(input_cypher)
				option, err := strconv.Atoi(input_cypher)

				if err != nil{
					fmt.Println("please enter a number")
					continue
				}

				switch option {
					case 1:
						encoding = "Rot13"
					case 2:
						encoding = "Reverse"
					case 3:
						encoding = "Caesar3"
					default:
						fmt.Println("Invalid selection")
						continue
					}
					fmt.Println("Your cypher is: "+ encoding)
					break

		}

		// message menu
		for{
				fmt.Println("Enter the message:")

				input_msg, _ :=  reader.ReadString('\n')

				input_msg = strings.TrimSpace(input_msg)

				if input_msg == ""{
					fmt.Println("Cannot be empty string")
					continue
				}

				message = input_msg
				break
		}


		return toEncrypt , message , encoding

}




// Encrypt the message with rot13
func Encrypt_rot13(s string) string {
	var encryptRot string

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z'{
			encryptRot += string('a' + (ch - 'a' + 13 )%26)
		}else if ch >= 'A' && ch <= 'Z'{
			encryptRot += string('A' + (ch - 'A' + 13 )%26)
		}else{
			encryptRot += string(ch)
		}
	}
	return encryptRot
}

// // Encrypt the message with reverse
func Encrypt_reverse(s string) string {
	var encryptReverse string

	for _, ch := range s {
		if ch >= 'a' && ch <= 'z'{
			encryptReverse += string('z' - (ch - 'a'))
		}else if ch >= 'A' && ch <= 'Z'{
			encryptReverse += string('Z' - (ch - 'A'))
		}else{
			encryptReverse += string(ch)
		}
	}
	return encryptReverse
}


func Encrypt_caesar(s string) string {
	var encryptCeaser string
	for _, ch:= range s {
		if ch >= 'a' && ch <= 'z' {

			encryptCeaser += string('a' + (ch -'a' + 3)%26 )
		} else if ch >= 'A' && ch <= 'Z' {
			encryptCeaser += string('A' + (ch -'A' + 3)%26)
		}else{
			encryptCeaser += string(ch)
		}
	}
	return encryptCeaser

}


/*======================Decrypting=============================*/



// // Decrypt the message with rot13
func Decrypt_rot13(s string) string {
	decrypt := Encrypt_rot13(s)
	return decrypt
}

// // Decrypt the message with reverse
func Decrypt_reverse(s string) string {
	decrypt_reverse := Encrypt_reverse(s)
	return decrypt_reverse
}




func Decrypt_caesar(s string) string {
	var decryptCeaser string
	for _, ch:= range s {
		if ch >= 'a' && ch <= 'z' {

			decryptCeaser += string('a' + (ch -'a' - 3 + 26)%26 )
		} else if ch >= 'A' && ch <= 'Z' {
			decryptCeaser += string('A' + (ch -'A' - 3 + 26 )%26)
		}else{
			decryptCeaser += string(ch)
		}
	}
	return decryptCeaser
}

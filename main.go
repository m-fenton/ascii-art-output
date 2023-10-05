package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func readData(banner string) string {
	// This function reads the data from the chosen banner file. If no banner is chosen an error message will be logged.
	data, err := os.ReadFile(banner)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	// If there are is only 1 arg or more than 4 args then print below statements and then exit.
	if len(os.Args) == 1 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	// Use textChecker to ascertain which arg is being used as our text variable.
	text := textChecker()
	// Use bannerChecker to ascertain which arg will be used to define the banner we want.
	banner := bannerChecker()
	splitStr := strings.Split(string(readData(banner+".txt")), "\n")
	// The string in our text variable noted above will be split by new line.
	splitArgs := strings.Split(string(text), "\\n")
	// for each slice of arg...
	for a := 0; a < len(splitArgs); a++ {
		runeArgs := []rune(splitArgs[a])
		for _, char := range runeArgs {
			if char != 10 && (char < 32 || char > 126) {
				fmt.Println("Character included in string which is not present in banner files.")
				os.Exit(0)
			}
		}
		// if the slice of arg contains nothing then print a new line
		if len(splitArgs[a]) == 0 {
			fmt.Println()
			// otherwise print the first line of each letter in the slice of arg, then 2nd - 8th
		} else {
			// If we have two args we will use the second arg as our string input. This string will be 'asciified' and printed to terminal.
			if len(os.Args) == 2 {
				for i := 1; i <= 8; i++ {
					for j := 0; j <= len(runeArgs)-1; j++ {
						letterArgs := runeArgs[j]
						// rune to line number
						lineNumber := (int(letterArgs)-32)*9 + i
						// prints from line number to the next 8 lines
						// We have named our completed answer to the variable name asciiText
						asciiText := splitStr[lineNumber]
						// We have used print to print our asciiText result to terminal.
						fmt.Print(asciiText)
					}
					// We have used print to print the necessary new line to terminal.
					fmt.Print("\n")
				}
				fmt.Println()
			} else {
				// The arg at position 1 is the flag, referred to as [OPTION], containing our file name.
				filename := os.Args[1]
				// We use strings.TrimPrefix to remove the flag leaving us with only the file name.
				filename = strings.TrimPrefix(filename, "--output=")
				// We create a file with this new trimmed filename. It is named to the variable destination.
				destination, err := os.Create(filename)
				// If the error does not equal nil then the process will close.
				if err != nil {
					fmt.Println("os.Create:", err)
					return
				}
				defer destination.Close()
				// If our args count is equal to 3 or 4 then we use this version of our ascii function to then print to our desired destination.

				for a := 0; a < len(splitArgs); a++ {
					runeArgs := []rune(splitArgs[a])
					for _, char := range runeArgs {
						if char != 10 && (char < 32 || char > 126) {
							fmt.Println("Character included in string which is not present in banner files.")
							os.Exit(0)
						}
					}
					for i := 1; i <= 8; i++ {
						for j := 0; j <= len(runeArgs)-1; j++ {
							letterArgs := runeArgs[j]
							// rune to line number
							lineNumber := (int(letterArgs)-32)*9 + i
							// prints from line number to the next 8 lines
							// We have named our completed answer to the variable name asciiText
							asciiText := splitStr[lineNumber]
							// We have used Fprint to print to our asciiText result to our variable destination.
							fmt.Fprint(destination, asciiText)
						}
						// We have used Fprint to print the necessary new line to our variable destination.
						fmt.Fprint(destination, "\n")
					}
				}
			}
		}
	}
}

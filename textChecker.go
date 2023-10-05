package main

import "os"

func textChecker() string {
	text := ""
	// If our args equal 2 then the 2nd arg, at position 1, will be used as our string.
	if len(os.Args) == 2 {
		text := os.Args[1]
		return text
	} else {
		// If our args are more than 2 then the 3rd arg, at position 2, will be used as our string.
		if len(os.Args) > 2 {
			text := os.Args[2]
			return text
		}
	}
	return text
}

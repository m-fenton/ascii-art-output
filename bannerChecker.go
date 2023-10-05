package main

import "os"

func bannerChecker() string {
	// This checks to see if the final argument in our command line is one of the three fonts.
	argsLen := len(os.Args) - 1
	if os.Args[argsLen] == "shadow" || os.Args[argsLen] == "standard" || os.Args[argsLen] == "thinkertoy" {
		// If so we return that font.
		banner := os.Args[argsLen]
		return banner
	} else {
		// Otherwise return "standard", acting as a default font.
		banner := "standard"
		return banner
	}
}

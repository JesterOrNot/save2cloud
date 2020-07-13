package main

import (
	"fmt"
	"os"
)

func PrintError(reason string, code int) {
	fmt.Println("\033[1;31mError:\033[m", reason)
	os.Exit(code)
}

func PrintWarning(warning string)  {
	fmt.Println("\033[1;33mError:\033[m", warning)	
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

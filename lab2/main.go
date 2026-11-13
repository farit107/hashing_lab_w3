package main

import (
	"fmt"
	"os"
	"strings"

	"lab2/utils/crack"
)

func main() {
	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	
	data, err := os.ReadFile("nord_vpn.txt")
	if err != nil {
		fmt.Println("Error: Could not read the file.")
		return
	}


	words := strings.Split(string(data), "\n")


	for _, word := range words {
		

		word = strings.TrimSpace(word)


		if word == "" {
			continue
		}


		currentHash := crack.Hashsha1(word)
		
		fmt.Printf("Checking: %s\n", word) 

		if currentHash == targetHash {
			fmt.Println("\n========================================")
			fmt.Println("PASSWORD FOUND:", word)
			fmt.Println("========================================")
			return 
		}
	}

	fmt.Println("Password not in list.")
}
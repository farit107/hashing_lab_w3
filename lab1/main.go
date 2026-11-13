package main

import (
	"fmt"
	"os"
	"strings"

	"lab1/utils/crack"
)

func main() {
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	
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


		currentHash := crack.HashMD5(word)
		
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
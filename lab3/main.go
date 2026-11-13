package main

import (
	"fmt"
	"os"
	"strings"

	"lab3/utils/crack"
)

func main() {
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	
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


		currentHash := crack.Hashsha512(word)
		
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
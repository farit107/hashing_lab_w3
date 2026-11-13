package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
	"crypto/sha3"
)

func hashMD5(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func hashSHA1(input string) string {
	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func hashSHA256(input string) string {
	hash := sha256.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func hashSHA512(input string) string {
	hash := sha512.Sum512([]byte(input))
	return hex.EncodeToString(hash[:])
}

func hashSHA3(input string) string {
	hash := sha3.Sum256([]byte(input))
	return hex.EncodeToString(hash[:])
}

func proofMe(input1, input2 string) {
	fmt.Println("\n========Name + Hashing Program========")
	
	md5A := hashMD5(input1)
	md5B := hashMD5(input2)
	
	sha1A := hashSHA1(input1)
	sha1B := hashSHA1(input2)
	
	sha256A := hashSHA256(input1)
	sha256B := hashSHA256(input2)
	
	sha512A := hashSHA512(input1)
	sha512B := hashSHA512(input2)
	
	sha3A := hashSHA3(input1)
	sha3B := hashSHA3(input2)
	
	fmt.Println("\n--- Quick Match Summary ---")
	
	matchMD5 := "No Match!"
	if md5A == md5B {
		matchMD5 = "Match!"
	}
	fmt.Printf("Hash (MD5): %s / %s => %s\n", md5A, md5B, matchMD5)
	
	matchSHA1 := "No Match!"
	if sha1A == sha1B {
		matchSHA1 = "Match!"
	}
	fmt.Printf("Hash (SHA1): %s / %s => %s\n", sha1A, sha1B, matchSHA1)
	
	matchSHA256 := "No Match!"
	if sha256A == sha256B {
		matchSHA256 = "Match!"
	}
	fmt.Printf("Hash (SHA256): %s / %s => %s\n", sha256A, sha256B, matchSHA256)
	
	matchSHA512 := "No Match!"
	if sha512A == sha512B {
		matchSHA512 = "Match!"
	}
	fmt.Printf("Hash (SHA512): %s / %s => %s\n", sha512A, sha512B, matchSHA512)
	
	matchSHA3 := "No Match!"
	if sha3A == sha3B {
		matchSHA3 = "Match!"
	}
	fmt.Printf("Hash (SHA3): %s / %s => %s\n", sha3A, sha3B, matchSHA3)
	
	fmt.Println("\n--- Detailed Hash Outputs ---")
	
	fmt.Printf("\nHash (MD5):\n")
	fmt.Printf("Output A = %s\n", md5A)
	fmt.Printf("Output B = %s\n", md5B)
	fmt.Printf("=> %s\n", matchMD5)
	
	fmt.Printf("\nHash (SHA1):\n")
	fmt.Printf("Output A = %s\n", sha1A)
	fmt.Printf("Output B = %s\n", sha1B)
	fmt.Printf("=> %s\n", matchSHA1)
	
	fmt.Printf("\nHash (SHA256):\n")
	fmt.Printf("Output A = %s\n", sha256A)
	fmt.Printf("Output B = %s\n", sha256B)
	fmt.Printf("=> %s\n", matchSHA256)
	
	fmt.Printf("\nHash (SHA512):\n")
	fmt.Printf("Output A = %s\n", sha512A)
	fmt.Printf("Output B = %s\n", sha512B)
	fmt.Printf("=> %s\n", matchSHA512)
	
	fmt.Printf("\nHash (SHA3-256):\n")
	fmt.Printf("Output A = %s\n", sha3A)
	fmt.Printf("Output B = %s\n", sha3B)
	fmt.Printf("=> %s\n", matchSHA3)
}

func main() {
	
	fmt.Print("Please input value 1: ")
	var input1 string
	fmt.Scanln(&input1)
	
	fmt.Print("Please input value 2: ")
	var input2 string
	fmt.Scanln(&input2)
	
	if strings.TrimSpace(input1) == "" || strings.TrimSpace(input2) == "" {
		fmt.Println("Error: Both inputs must be non-empty strings")
		os.Exit(1)
	}
	
	proofMe(input1, input2)
}
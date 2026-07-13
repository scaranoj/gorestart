package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	pw := "chode123"

	sb, err := bcrypt.GenerateFromPassword([]byte(pw), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}

	loginPword1 := "chode123"
	// Nope, line below is totally unecessary. Just use the `sb`` var since it already contains a byte slice value
	// hashedPword2 := []byte{36, 50, 97, 36, 48, 52, 36, 75, 72, 102, 57, 55, 110, 103, 112, 82, 116, 113, 57, 47, 118, 113, 56, 46, 54, 77, 110, 103, 117, 48, 53, 113, 65, 48, 114, 99, 116, 86, 56, 75, 99, 84, 86, 119, 112, 57, 101, 79, 119, 85, 111, 49, 116, 87, 72, 85, 107, 111, 48, 105}

	err = bcrypt.CompareHashAndPassword(sb, []byte(loginPword1))
	if err != nil {
		fmt.Println(err)

	}

	fmt.Println(pw)
	fmt.Println(sb)
}

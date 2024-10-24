package main

import (
	"fmt"
	"os"

	"github.com/mrkouhadi/go-storage/filesystem"
	"github.com/mrkouhadi/go-storage/storageAPIs"
	"github.com/mrkouhadi/go-storage/utils"
)

func main() {
	err := os.MkdirAll("data", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	tokens := utils.TokenData{
		AccessToken:  "iam@access@token",
		RefreshToken: "iam@refresh@token",
	}
	// file system
	err = filesystem.SaveTokensToFile("./data/jwtTokens", tokens)
	if err != nil {
		fmt.Println("Error saving toke in file system: ", err)
	}
	d, err := filesystem.LoadTokensFromFile("./data/jwtTokens")
	if err != nil {
		fmt.Println("error reading tokens from file system: ", err)
	}
	fmt.Println(d.AccessToken)
	// keychain (MACOS) & Credential manager (windows) & GNOME Keyring (Lunix)
	err = storageAPIs.SaveTokens(tokens)
	if err != nil {
		fmt.Println("error saving tokens in keychain/Credential manager: ", err)
	}
	tokens, err = storageAPIs.LoadTokens()
	if err != nil {
		fmt.Println("error saving tokens in keychain/Credential manager: ", err)
	}
	fmt.Println(tokens)
	err = storageAPIs.ClearTokens()
	if err != nil {
		fmt.Println("error clearing tokens: ", err)
	}
	fmt.Println(tokens)
}

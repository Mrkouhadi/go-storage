package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mrkouhadi/go-storage/filesystem"
	"github.com/mrkouhadi/go-storage/leveldbstorage"
	"github.com/mrkouhadi/go-storage/storageAPIs"
	"github.com/mrkouhadi/go-storage/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

func main() {
	tokens := utils.TokenData{
		AccessToken:  "iam@access@token",
		RefreshToken: "iam@refresh@token",
	}
	///////////////////////////////////////////////////////////// file system
	err := os.MkdirAll("data", 0755)
	if err != nil {
		fmt.Println("Error creating folder:", err)
		return
	}
	// Save tokens
	err = filesystem.SaveTokensToFile("./data/jwtTokens", tokens)
	if err != nil {
		fmt.Println("Error saving toke in file system: ", err)
	}
	// Load tokens
	d, err := filesystem.LoadTokensFromFile("./data/jwtTokens")
	if err != nil {
		fmt.Println("error reading tokens from file system: ", err)
	}
	fmt.Println(d.AccessToken)
	// Clear tokens
	err = filesystem.ClearTokens("./data/jwtTokens")
	if err != nil {
		fmt.Println("Error clearing tokens:", err)
	} else {
		fmt.Println("Tokens cleared successfully.")
	}
	///////////////////////////////////////////////////////////// Platform specific APIs
	// keychain (MACOS) & Credential manager (windows) & GNOME Keyring (Lunix)
	// Save tokens
	err = storageAPIs.SaveTokens(tokens)
	if err != nil {
		fmt.Println("error saving tokens in keychain/Credential manager: ", err)
	}
	// Load tokens
	tokens, err = storageAPIs.LoadTokens()
	if err != nil {
		fmt.Println("error saving tokens in keychain/Credential manager: ", err)
	}
	fmt.Println(tokens)
	// Clear tokens
	err = storageAPIs.ClearTokens()
	if err != nil {
		fmt.Println("error clearing tokens: ", err)
	}
	fmt.Println(tokens)

	///////////////////////////////////////////////////////////////  LEVEL DB
	// Open the LevelDB database
	db, err := leveldb.OpenFile("data/db", nil)
	if err != nil {
		log.Fatal("Error opening LevelDB:", err)
	}
	defer db.Close()
	// Save tokens to LevelDB
	err = leveldbstorage.SaveTokensToDB(db, "jwtTokens", tokens)
	if err != nil {
		log.Fatal("Error saving tokens to LevelDB:", err)
	}

	// Load tokens from LevelDB
	loadedTokens, err := leveldbstorage.LoadTokensFromDB(db, "jwtTokens")
	if err != nil {
		log.Fatal("Error loading tokens from LevelDB:", err)
	}
	fmt.Println("Loaded AccessToken:", loadedTokens.AccessToken)

	// Clear tokens from LevelDB
	err = leveldbstorage.ClearTokensFromDB(db, "jwtTokens")
	if err != nil {
		log.Fatal("Error clearing tokens from LevelDB:", err)
	}
}

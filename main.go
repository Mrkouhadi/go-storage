package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/mrkouhadi/go-storage/filesystem"
	"github.com/mrkouhadi/go-storage/leveldbstorage"
	"github.com/mrkouhadi/go-storage/sqlite3storage"
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
	tokensFromFileSystem, err := filesystem.LoadTokensFromFile("./data/jwtTokens")
	if err != nil {
		fmt.Println("error reading tokens from file system: ", err)
	}
	fmt.Println("Tokens from File system: ", tokensFromFileSystem)
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
	tokensFromPlatformSpecificAPIs, err := storageAPIs.LoadTokens()
	if err != nil {
		fmt.Println("error saving tokens in keychain/Credential manager: ", err)
	}
	fmt.Println("Tokens From Platform Specific APIs: ", tokensFromPlatformSpecificAPIs)
	// Clear tokens
	err = storageAPIs.ClearTokens()
	if err != nil {
		fmt.Println("error clearing tokens: ", err)
	} else {
		fmt.Println("Tokens cleared successfully.")
	}
	///////////////////////////////////////////////////////////////  LEVEL DB
	// Open the LevelDB database
	level_db, err := leveldb.OpenFile("data/level_db", nil)
	if err != nil {
		log.Fatal("Error opening LevelDB:", err)
	}
	defer level_db.Close()
	// Save tokens to LevelDB
	err = leveldbstorage.SaveTokensToDB(level_db, "jwtTokens", tokens)
	if err != nil {
		log.Fatal("Error saving tokens to LevelDB:", err)
	}
	// Load tokens from LevelDB
	tokensFromLevelDb, err := leveldbstorage.LoadTokensFromDB(level_db, "jwtTokens")
	if err != nil {
		log.Fatal("Error loading tokens from LevelDB:", err)
	}
	fmt.Println("Tokens from Level DB:", tokensFromLevelDb)
	// Clear tokens from LevelDB
	err = leveldbstorage.ClearTokensFromDB(level_db, "jwtTokens")
	if err != nil {
		log.Fatal("Error clearing tokens from LevelDB:", err)
	} else {
		fmt.Println("Tokens cleared successfully.")
	}
	///////////////////////////////////////////////////////////////  sqlite3 DB
	// Open the SQLite3 database (this will create the file if it doesn't exist)
	sqlite_db, err := sql.Open("sqlite3", "data/sqlite3/tokens.db")
	if err != nil {
		log.Fatal("Error opening SQLite3 DB:", err)
	}
	defer sqlite_db.Close()

	// Initialize the database
	err = sqlite3storage.InitializeDB(sqlite_db)
	if err != nil {
		log.Fatal("Error initializing SQLite3 DB:", err)
	}

	// Save tokens to SQLite3
	err = sqlite3storage.SaveTokensToDB(sqlite_db, "jwtTokens", tokens)
	if err != nil {
		log.Fatal("Error saving tokens to SQLite3:", err)
	}

	// Load tokens from SQLite3
	loadedTokens, err := sqlite3storage.LoadTokensFromDB(sqlite_db, "jwtTokens")
	if err != nil {
		log.Fatal("Error loading tokens from SQLite3:", err)
	}
	fmt.Println("Tokens from sqlite3 :", loadedTokens)

	// Clear tokens from SQLite3
	err = sqlite3storage.ClearTokensFromDB(sqlite_db, "jwtTokens")
	if err != nil {
		log.Fatal("Error clearing tokens from SQLite3:", err)
	} else {
		fmt.Println("Tokens cleared successfully.")
	}
}

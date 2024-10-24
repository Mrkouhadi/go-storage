package sqlite3storage

import (
	"database/sql"
	"encoding/json"
	"errors"

	_ "github.com/mattn/go-sqlite3" // SQLite3 driver
	"github.com/mrkouhadi/go-storage/utils"
)

// InitializeDB initializes the SQLite3 database and ensures the tokens table exists
func InitializeDB(db *sql.DB) error {
	// Create tokens table if it does not exist
	query := `CREATE TABLE IF NOT EXISTS tokens (
		id TEXT PRIMARY KEY,
		data BLOB
	);`
	_, err := db.Exec(query)
	return err
}

// SaveTokensToDB saves encrypted tokens to SQLite3
func SaveTokensToDB(db *sql.DB, key string, tokenData utils.TokenData) error {
	// Serialize the token data to JSON
	jsonData, err := json.Marshal(tokenData)
	if err != nil {
		return err
	}

	// Encrypt the data
	encryptedData, err := utils.Encrypt(jsonData, utils.EncryptionKey)
	if err != nil {
		return err
	}

	// Insert or replace the token data into the SQLite3 database
	query := `INSERT OR REPLACE INTO tokens (id, data) VALUES (?, ?)`
	_, err = db.Exec(query, key, encryptedData)
	return err
}

// LoadTokensFromDB loads and decrypts tokens from SQLite3
func LoadTokensFromDB(db *sql.DB, key string) (utils.TokenData, error) {
	// Retrieve the encrypted data from the SQLite3 database
	query := `SELECT data FROM tokens WHERE id = ?`
	row := db.QueryRow(query, key)

	var encryptedData []byte
	err := row.Scan(&encryptedData)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return utils.TokenData{}, errors.New("tokens not found")
		}
		return utils.TokenData{}, err
	}

	// Decrypt the data
	decryptedData, err := utils.Decrypt(encryptedData, utils.EncryptionKey)
	if err != nil {
		return utils.TokenData{}, err
	}

	// Deserialize decrypted JSON data into TokenData struct
	var tokenData utils.TokenData
	err = json.Unmarshal(decryptedData, &tokenData)
	if err != nil {
		return utils.TokenData{}, err
	}

	return tokenData, nil
}

// ClearTokensFromDB removes the token entry from SQLite3
func ClearTokensFromDB(db *sql.DB, key string) error {
	// Delete the token data from the SQLite3 database
	query := `DELETE FROM tokens WHERE id = ?`
	_, err := db.Exec(query, key)
	return err
}

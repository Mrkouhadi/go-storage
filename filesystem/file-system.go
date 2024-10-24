package filesystem

import (
	"encoding/json"
	"os"

	"github.com/mrkouhadi/go-storage/utils"
)

// SaveTokensToFile saves encrypted tokens to file
func SaveTokensToFile(filename string, tokenData utils.TokenData) error {
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

	// Write encrypted data to a file
	return os.WriteFile(filename, encryptedData, 0644)
}

// LoadTokensFromFile loads and decrypts tokens from file
func LoadTokensFromFile(filename string) (utils.TokenData, error) {
	// Read the encrypted data from the file
	encryptedData, err := os.ReadFile(filename)
	if err != nil {
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

// ClearTokens removes the token file from the filesystem
func ClearTokens(filename string) error {
	// Remove the file from the filesystem
	err := os.Remove(filename)
	if os.IsNotExist(err) {
		// If the file does not exist, it's not an error
		return nil
	}
	return err
}

package filesystem

import (
	"encoding/json"
	"os"

	"github.com/mrkouhadi/go-storage/utils"
)

// Save encrypted tokens to file
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

// Load and decrypt tokens from file
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

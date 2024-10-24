// store data securely using platform specific APIs (automatically data gets encrypted )
// keychain (MACOS) & Credential manager (windows) & GNOME Keyring (Lunix)

package storageAPIs

import (
	"encoding/json"

	"github.com/mrkouhadi/go-storage/utils"

	"github.com/zalando/go-keyring"
)

// Define a service name and key
const (
	serviceName = "Brifel"
	tokenKey    = "jwtTokens"
)

// Save tokens to the keychain or credential manager
func SaveTokens(tokenData utils.TokenData) error {
	// Serialize the token data to JSON
	jsonData, err := json.Marshal(tokenData)
	if err != nil {
		return err
	}

	// Store data in the keychain or credential manager
	return keyring.Set(serviceName, tokenKey, string(jsonData))
}

// Load tokens from the keychain or credential manager
func LoadTokens() (utils.TokenData, error) {
	// Retrieve the data from the keychain or credential manager
	encryptedData, err := keyring.Get(serviceName, tokenKey)
	if err != nil {
		return utils.TokenData{}, err
	}

	// Deserialize JSON data into TokenData struct
	var tokenData utils.TokenData
	err = json.Unmarshal([]byte(encryptedData), &tokenData)
	if err != nil {
		return utils.TokenData{}, err
	}
	return tokenData, nil
}

// Clear tokens from the keychain or credential manager
func ClearTokens() error {
	// Delete the entry from the keychain or credential manager
	return keyring.Delete(serviceName, tokenKey)
}

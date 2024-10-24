package leveldbstorage

import (
	"encoding/json"
	"errors"

	"github.com/mrkouhadi/go-storage/utils"
	"github.com/syndtr/goleveldb/leveldb"
)

// SaveTokensToDB saves encrypted tokens to LevelDB
func SaveTokensToDB(db *leveldb.DB, key string, tokenData utils.TokenData) error {
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

	// Save the encrypted data to LevelDB
	err = db.Put([]byte(key), encryptedData, nil)
	return err
}

// LoadTokensFromDB loads and decrypts tokens from LevelDB
func LoadTokensFromDB(db *leveldb.DB, key string) (utils.TokenData, error) {
	// Retrieve the encrypted data from LevelDB
	encryptedData, err := db.Get([]byte(key), nil)
	if err != nil {
		if err == leveldb.ErrNotFound {
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

// ClearTokensFromDB removes the token entry from LevelDB
func ClearTokensFromDB(db *leveldb.DB, key string) error {
	// Delete the token data from LevelDB
	err := db.Delete([]byte(key), nil)
	return err
}

package security

import (
	"crypto/rand"
	"errors"
	"github.com/flasherup/gradtage.de/common"
	"github.com/tomogoma/go-api-guard"
)

const notFoundError = "key not found"

type KeyType struct {
	value []byte
}

func (k KeyType) Value() []byte {
	return k.value
}


type Store struct {
	Keys map[string]KeyType
}

func (s Store)IsNotFoundError(err error) bool {
	if err != nil && err.Error() == notFoundError {
		return true
	}
	return false
 }

func (s Store)InsertAPIKey(userID string, key []byte) (api.Key, error) {
	k := KeyType{ value:key}
	s.Keys[userID] = k
	return k, nil
}

func (s Store)APIKeyByUserIDVal(userID string, key []byte) (api.Key, error) {
	if k, ok := s.Keys[userID]; ok {
		return k,nil
	}
	s.InsertAPIKey(userID, key)
	return s.Keys[userID],errors.New(notFoundError)
}

type KeyGen struct {

}

func (kg KeyGen)SecureRandomBytes(length int) ([]byte, error) {
	nonce := make([]byte, length)
	if _, err := rand.Read(nonce); err != nil {
		panic(err.Error())
	}

	fixed, err := common.GenerateRandomString(length)
	if err != nil {
		return nil,err
	}

	return []byte(fixed),nil
}


type KeyManager struct {
	KeyGuard api.Guard
	KeyStore Store
}


func NewKeyManager() (*KeyManager, error) {
	s := Store{ Keys:make(map[string]KeyType) }
	kg := KeyGen{}

	guard, err := api.NewGuard(s, api.WithKeyGenerator(kg))
	if err != nil {
		return nil, err
	}

	res := KeyManager{*guard, s}

 	return &res,nil
}

func (km *KeyManager) RestoreKeys( keys map[string]string) error {
	for k,v := range keys {
		km.KeyStore.InsertAPIKey(k,[]byte(v))
	}
	return nil
}

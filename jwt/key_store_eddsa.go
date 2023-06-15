package jwt

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
)

// KeyStoreEdDSA is a key store for EdDSA keys
type KeyStoreEdDSA struct {
	PublicKey  ed25519.PublicKey
	PrivateKey ed25519.PrivateKey
}

// LoadOrGenerateKeysEdDSA creates a new EdDSA key store from the given private and public key paths or generates a new key pair if the files do not exist
func LoadOrGenerateKeysEdDSA(privPath, pubPath string) (*KeyStoreEdDSA, error) {
	ks := &KeyStoreEdDSA{}

	if fileExist(privPath) {
		if err := ks.LoadPrivateKeyFromFile(privPath); err != nil {
			return ks, err
		}
	} else {
		ks.GenerateKeys()
		if err := ks.SavePrivateKey(privPath); err != nil {
			return ks, err
		}
	}

	if fileExist(pubPath) {
		if err := ks.LoadPublicKeyFromFile(pubPath); err != nil {
			return ks, err
		}
	} else {
		if ks.PublicKey == nil {
			publicKey, ok := ks.PrivateKey.Public().(ed25519.PublicKey)
			if !ok {
				ks.PublicKey = publicKey
			} else {
				return ks, fmt.Errorf("invalid public key")
			}
		}

		if err := ks.SavePublicKey(pubPath); err != nil {
			return ks, err
		}
	}
	return ks, nil
}

// GenerateKeys generates a new Ed25519 key pair
func (ks *KeyStoreEdDSA) GenerateKeys() error {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return err
	}

	ks.PrivateKey = privateKey
	ks.PublicKey = publicKey

	return nil
}

// LoadPrivateKeyFromFile loads the private key from the specified path
func (ks *KeyStoreEdDSA) LoadPrivateKeyFromFile(path string) error {
	// Read the private key file
	privateKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Parse the PEM-encoded private key
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		return fmt.Errorf("invalid private key format")
	}

	// Extract the raw private key bytes
	privateKey := ed25519.NewKeyFromSeed(block.Bytes)

	ks.PrivateKey = privateKey
	return nil
}

// LoadPublicKeyFromFile loads the public key from the specified path
func (ks *KeyStoreEdDSA) LoadPublicKeyFromFile(path string) error {
	// Read the public key file
	publicKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	// Parse the PEM-encoded public key
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return fmt.Errorf("invalid public key format")
	}

	ks.PublicKey = block.Bytes

	return nil
}

// LoadPrivateKeyFromString loads the private key from a string
func (ks *KeyStoreEdDSA) LoadPublicKeyFromString(str string) error {
	publicKeyBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

	// Parse the PEM-encoded public key
	block, _ := pem.Decode(publicKeyBytes)
	if block == nil || block.Type != "PUBLIC KEY" {
		return fmt.Errorf("invalid public key format")
	}

	ks.PublicKey = block.Bytes
	return nil
}

// LoadPublicKeyFromString loads the public key from a string
func (ks *KeyStoreEdDSA) LoadPrivateKeyFromString(str string) error {
	privateKeyBytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return err
	}

	// Parse the PEM-encoded private key
	block, _ := pem.Decode(privateKeyBytes)
	if block == nil || block.Type != "PRIVATE KEY" {
		return fmt.Errorf("invalid private key format")
	}

	// Extract the raw private key bytes
	privateKey := ed25519.NewKeyFromSeed(block.Bytes)

	ks.PrivateKey = privateKey
	return nil
}

// SavePrivateKey saves the private key to the specified path
func (ks *KeyStoreEdDSA) SavePrivateKey(path string) error {
	// Encode the private key to PEM format
	privateKeyPEM := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: ed25519PrivateKeyToDER(ks.PrivateKey),
	}
	privateKeyPEMBytes := pem.EncodeToMemory(privateKeyPEM)

	// Save private key to file
	err := ioutil.WriteFile(path, privateKeyPEMBytes, 0600)
	if err != nil {
		return err
	}

	return nil
}

func ed25519PrivateKeyToDER(privateKey ed25519.PrivateKey) []byte {
	// Ed25519 private key is already in the correct format (raw bytes)
	return privateKey.Seed()
}

// SavePublicKey saves the public key to the specified path
func (ks *KeyStoreEdDSA) SavePublicKey(path string) error {
	// Encode the public key to PEM format
	publicKeyPEM := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: ks.PublicKey,
	}
	publicKeyPEMBytes := pem.EncodeToMemory(publicKeyPEM)

	// Save public key to file
	err := ioutil.WriteFile(path, publicKeyPEMBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

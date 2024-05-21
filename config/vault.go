package config

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/vault/api"
	"github.com/joho/godotenv"
)

type VaultService struct {
	Vault *api.Client
}

func InitVault(address, token string) (*VaultService, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}

	client, err := api.NewClient(&api.Config{
		Address: address,
	})
	if err != nil {
		return nil, err
	}

	client.SetToken(token)

	vs := &VaultService{
		Vault: client,
	}

	return vs, nil
}

func (vs *VaultService) DecryptValue(query string, value string) (string, error) {
	path := strings.ToLower("transit/decrypt/"+os.Getenv("VAULT_APP_NAME")) + query
	secret, err := vs.Vault.Logical().Write(path, map[string]interface{}{
		"ciphertext": value,
	})
	if err != nil {
		return "", err
	}
	if secret == nil || secret.Data == nil {
		return "", fmt.Errorf("decryption failed")
	}

	decodeString, err := Base64Decode(secret.Data["plaintext"].(string))
	if err != nil {
		return "", fmt.Errorf(decodeString)
	}
	return decodeString, nil
}

func Base64Decode(encrypt string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return "", err
	}
	return string(decoded), nil
}

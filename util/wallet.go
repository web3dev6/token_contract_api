package util

import (
	"log"
	"regexp"
)

// CreateWallet creates a pvt-pub key pair, stores it securely and returns the walletAddress
func CreateWallet() (string, error) {
	// TODO: implement CreateWallet
	return "0x5D7E7B133E5f16C75A18e3b04Ac9Af85451C209c", nil
}

// IsValidEthAddress returns true if provided eth address is in valid format
func IsValidEthAddress(address string) bool {
	// Regular expression to match a typical Ethereum wallet address
	pattern := `^0x[a-fA-F0-9]{40}$`
	match, err := regexp.MatchString(pattern, address)
	if err != nil {
		log.Println("error in regexp matching", err)
		return false
	}

	return match
}

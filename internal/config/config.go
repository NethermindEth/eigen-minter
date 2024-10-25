package config

import (
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/spf13/viper"
)

var chainIDs = map[string]uint64{
	"mainnet": 1,
	"holesky": 17000,
}

var networkRPCs = map[string]string{
	"mainnet": "https://eth.llamarpc.com",
	"holesky": "https://ethereum-holesky-rpc.publicnode.com",
}

var contractAddresses = map[string]string{
	"holesky": "0x8DaaE33cB2da8dA23595ADB19f271EF41E34bd8C",
	"mainnet": "0x0ffC6AC10515EE0F83fEE71FCaf5Ea5805256563",
}

func ChainID(network string) (uint64, error) {
	chainID, ok := chainIDs[network]
	if !ok {
		return 0, fmt.Errorf("network %s is not supported", network)
	}
	return chainID, nil
}

func NetworkRPC(network string) (string, error) {
	rpc, ok := networkRPCs[network]
	if !ok {
		return "", fmt.Errorf("network %s is not supported", network)
	}
	return rpc, nil
}

func ContractAddress(network string) (string, error) {
	address, ok := contractAddresses[network]
	if !ok {
		return "", fmt.Errorf("network %s is not supported", network)
	}
	return address, nil
}

func ValidateConfig() (Config, error) {
	cfg := Config{}

	cfg.Network = viper.GetString("network")
	if _, ok := contractAddresses[cfg.Network]; !ok {
		return cfg, fmt.Errorf("network %s is not supported", cfg.Network)
	}

	cfg.RPC = viper.GetString("rpc-endpoint")
	if cfg.RPC == "" {
		slog.Info(fmt.Sprintf("rpc-endpoint is not set, using default for %s: %s", cfg.Network, networkRPCs[cfg.Network]))
		cfg.RPC = networkRPCs[cfg.Network]
	}

	cfg.Contract = viper.GetString("contract-address")
	if cfg.Contract == "" {
		slog.Info(fmt.Sprintf("contract-address is not set, using default for %s: %s", cfg.Network, contractAddresses[cfg.Network]))
		cfg.Contract = contractAddresses[cfg.Network]
	}

	pvKeyPath := viper.GetString("private-key-file")
	if pvKeyPath != "" {
		pvKeyBytes, err := os.ReadFile(pvKeyPath)
		if err != nil {
			slog.Error(fmt.Sprintf("failed to read private key file: %v", err))
		}
		if err := validatePrivateKey(string(pvKeyBytes)); err != nil {
			slog.Error(fmt.Sprintf("invalid private key: %v", err))
			// Try now to get the private key from the private-key setting
		} else {
			slog.Debug("private key is valid")
			cfg.PrivateKey = string(pvKeyBytes)
		}
	}
	if cfg.PrivateKey == "" {
		cfg.PrivateKey = viper.GetString("private-key")
		if err := validatePrivateKey(cfg.PrivateKey); err != nil {
			return cfg, fmt.Errorf("invalid private key: %v", err)
		}
	}

	cfg.FromAddress = viper.GetString("from-address")
	if cfg.FromAddress == "" {
		pvKey, err := crypto.HexToECDSA(cfg.PrivateKey)
		if err != nil {
			return cfg, fmt.Errorf("failed to convert private key to ECDSA: %v", err)
		}
		publicKeyECDSA, ok := pvKey.Public().(*ecdsa.PublicKey)
		if !ok {
			return cfg, fmt.Errorf("failed to get public key")
		}
		address := crypto.PubkeyToAddress(*publicKeyECDSA)
		slog.Info(fmt.Sprintf("from-address is not set, using default for %s: %s", cfg.Network, address.Hex()))
		cfg.FromAddress = address.Hex()
	}

	return cfg, nil
}

func validatePrivateKey(privateKey string) error {
	// Remove "0x" prefix if present
	privateKey = strings.TrimPrefix(privateKey, "0x")

	// Check if the private key has the correct length
	if len(privateKey) != 64 {
		return fmt.Errorf("invalid private key length: expected 64 characters, got %d", len(privateKey))
	}

	// Try to parse the private key
	_, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return fmt.Errorf("invalid private key: %v", err)
	}

	return nil
}

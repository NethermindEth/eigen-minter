package cli

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log/slog"
	"math/big"
	"os"
	"strings"

	"github.com/NethermindEth/eigen-minter/internal/client"
	"github.com/NethermindEth/eigen-minter/internal/contract"
	"github.com/NethermindEth/eigen-minter/internal/metrics"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logLevel string

type Config struct {
	Network     string
	RPC         string
	Contract    string
	PushGateway string
	PrivateKey  string
	FromAddress string
}

var contractAddresses = map[string]string{
	"holesky": "0x8DaaE33cB2da8dA23595ADB19f271EF41E34bd8C",
	"mainnet": "0x0ffC6AC10515EE0F83fEE71FCaf5Ea5805256563",
}

var chainIDs = map[string]uint64{
	"mainnet": 1,
	"holesky": 17000,
}

var networkRPCs = map[string]string{
	"mainnet": "https://eth.llamarpc.com",
	"holesky": "https://ethereum-holesky-rpc.publicnode.com",
}

func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "eigen-minter",
		Short: "Eigen Minter is a CLI application for interacting with the TokenHopper smart contract",
		Long:  "A cronjob to call Eigenlayer Token Hopper contracts functions periodically to mint EIGEN tokens for programmatic incentives.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			logLevel = viper.GetString("log-level")
			level := slog.Level(slog.LevelInfo)
			if err := level.UnmarshalText([]byte(logLevel)); err != nil {
				return err
			}
			slog.SetLogLoggerLevel(level)
			return nil
		},
		RunE: run,
	}
	// Disable completion default cmd
	cmd.CompletionOptions.DisableDefaultCmd = true

	cmd.Flags().String("network", "holesky", "Ethereum network to use (mainnet or holesky)")
	cmd.Flags().String("rpc-endpoint", "", "RPC endpoint URL")
	cmd.Flags().String("contract-address", "", "TokenHopper contract address")
	cmd.Flags().String("pushgateway-url", "", "Prometheus Pushgateway URL")
	cmd.Flags().String("private-key", "", "Private key to use for transactions")
	cmd.Flags().String("from-address", "", "From address to use for transactions")
	cmd.Flags().String("private-key-file", "", "File containing the private key to use for transactions. If a private key is found in the file, it will be used instead of the private-key flag")
	cmd.Flags().String("log-level", "info", "Log level (debug, info, warn, error)")

	viper.BindPFlag("network", cmd.Flags().Lookup("network"))
	viper.BindPFlag("rpc-endpoint", cmd.Flags().Lookup("rpc-endpoint"))
	viper.BindPFlag("contract-address", cmd.Flags().Lookup("contract-address"))
	viper.BindPFlag("pushgateway-url", cmd.Flags().Lookup("pushgateway-url"))
	viper.BindPFlag("private-key", cmd.Flags().Lookup("private-key"))
	viper.BindPFlag("from-address", cmd.Flags().Lookup("from-address"))
	viper.BindPFlag("private-key-file", cmd.Flags().Lookup("private-key-file"))
	viper.BindPFlag("log-level", cmd.Flags().Lookup("log-level"))

	viper.BindEnv("network", "EIGEN_MINTER_NETWORK")
	viper.BindEnv("rpc-endpoint", "EIGEN_MINTER_RPC_ENDPOINT")
	viper.BindEnv("contract-address", "EIGEN_MINTER_CONTRACT_ADDRESS")
	viper.BindEnv("pushgateway-url", "EIGEN_MINTER_PUSHGATEWAY_URL")
	viper.BindEnv("private-key", "EIGEN_MINTER_PRIVATE_KEY")
	viper.BindEnv("from-address", "EIGEN_MINTER_FROM_ADDRESS")
	viper.BindEnv("private-key-file", "EIGEN_MINTER_PRIVATE_KEY_FILE")

	viper.SetEnvPrefix("EIGEN_MINTER")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	return cmd
}

func run(cmd *cobra.Command, args []string) error {
	var m *metrics.Metrics
	pushGatewayURL := viper.GetString("pushgateway-url")
	if pushGatewayURL != "" {
		m = metrics.NewMetrics(pushGatewayURL)
		m.RecordTrigger()
		slog.Debug(fmt.Sprintf("Pushing metrics to %s", pushGatewayURL))
		if err := m.Push(); err != nil {
			slog.Error(fmt.Sprintf("Failed to push metrics: %v\n", err))
		}
	}

	cfg, err := validateConfig()
	if err != nil {
		return err
	}

	slog.Debug(fmt.Sprintf("Connecting to RPC: %s", cfg.RPC))
	rpcClient, err := client.ConnectClient(chainIDs[cfg.Network], cfg.RPC)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %v", err)
	}

	slog.Debug(fmt.Sprintf("Connecting to contract: %s", cfg.Contract))
	c, err := contract.NewContract(common.HexToAddress(cfg.Contract), rpcClient)
	if err != nil {
		return fmt.Errorf("failed to create contract: %v", err)
	}

	slog.Info("Checking if can press")
	canPress, err := c.CanPress(nil)
	if err != nil {
		return fmt.Errorf("failed to check if can press: %v", err)
	}

	if canPress {
		slog.Info("Pressing button")
		if err := pressButton(cfg, chainIDs[cfg.Network], rpcClient, c); err != nil {
			slog.Error(fmt.Sprintf("failed to press button: %v", err))
			if m != nil {
				m.RecordPressButtonFailure()
			}
		} else {
			slog.Info("Button pressed successfully")
			if m != nil {
				m.RecordPressButtonSuccess()
			}
		}
	} else {
		slog.Info("Cannot press button at this time")
		if m != nil {
			m.RecordPressButtonFailure()
		}
	}

	if m != nil {
		if err := m.Push(); err != nil {
			slog.Error(fmt.Sprintf("Failed to push metrics: %v\n", err))
		}
	}

	return nil
}

func validateConfig() (Config, error) {
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

func pressButton(cfg Config, chainID uint64, rpcClient *ethclient.Client, c *contract.Contract) error {
	// Create a new transactor
	pvKey, err := crypto.HexToECDSA(cfg.PrivateKey)
	if err != nil {
		return fmt.Errorf("failed to convert private key to ECDSA: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(pvKey, big.NewInt(int64(chainID))) // Use appropriate chain ID
	if err != nil {
		return fmt.Errorf("failed to create transactor: %v", err)
	}

	// Get the latest nonce
	nonce, err := rpcClient.PendingNonceAt(context.Background(), common.HexToAddress(cfg.FromAddress))
	if err != nil {
		return fmt.Errorf("failed to retrieve account nonce: %v", err)
	}

	// Set up transaction options
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei

	// Estimate gas limit
	gasLimit, err := estimateGas(auth, rpcClient, common.HexToAddress(cfg.Contract))
	if err != nil {
		return fmt.Errorf("failed to estimate gas: %v", err)
	}
	// Add a buffer to the estimated gas limit
	auth.GasLimit = uint64(float64(gasLimit) * 1.2) // 20% buffer

	// Get suggested gas price and add a small buffer
	suggestedGasPrice, err := rpcClient.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to suggest gas price: %v", err)
	}
	auth.GasPrice = new(big.Int).Mul(suggestedGasPrice, big.NewInt(105))
	auth.GasPrice = auth.GasPrice.Div(auth.GasPrice, big.NewInt(100)) // 5% higher than suggested

	tx, err := c.PressButton(auth)
	if err != nil {
		return fmt.Errorf("failed to press button: %v", err)
	}

	slog.Info(fmt.Sprintf("Transaction sent: %s", tx.Hash().Hex()))

	// Wait for the transaction to be mined
	receipt, err := bind.WaitMined(context.Background(), rpcClient, tx)
	if err != nil {
		slog.Error(fmt.Sprintf("failed to wait for transaction to be mined: %v", err))
	}

	slog.Info(fmt.Sprintf("Transaction mined in block %d", receipt.BlockNumber.Uint64()))
	return nil
}

func estimateGas(auth *bind.TransactOpts, rpcClient *ethclient.Client, contractAddress common.Address) (uint64, error) {
	// Parse contract ABI
	contractABI, err := abi.JSON(strings.NewReader(contract.ContractMetaData.ABI))
	if err != nil {
		return 0, fmt.Errorf("failed to parse contract ABI: %v", err)
	}

	// Create a new transaction object
	data, err := contractABI.Pack("pressButton")
	if err != nil {
		return 0, fmt.Errorf("failed to pack data: %v", err)
	}

	// Estimate gas
	msg := ethereum.CallMsg{
		From:     auth.From,
		To:       &contractAddress,
		Gas:      0,
		GasPrice: auth.GasPrice,
		Value:    auth.Value,
		Data:     data,
	}

	gasLimit, err := rpcClient.EstimateGas(context.Background(), msg)
	if err != nil {
		return 0, fmt.Errorf("failed to estimate gas: %v", err)
	}

	return gasLimit, nil
}

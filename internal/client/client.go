/*
Copyright 2022 Nethermind

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package client 

import (
	"context"
	"fmt"
	"math/big"
	"log/slog"

	"github.com/ethereum/go-ethereum/ethclient"
)

// ConnectClient returns a new Ethereum client connected to the given network.
func ConnectClient(chainID uint64, RPC string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(RPC)
	if err == nil {
		// Try to get the chain ID, which is a basic operation that should work for any Ethereum client
		clientChainID, err := client.ChainID(context.Background())
		if err == nil {
			if clientChainID.Cmp(new(big.Int).SetUint64(chainID)) == 0 {
				// If we successfully got the chain ID and it matches the expected one,
				// we can be reasonably sure this is the correct Ethereum client
				slog.Debug(fmt.Sprintf("connected to %s", RPC))
				return client, nil
			} else {
				slog.Error(fmt.Sprintf("chain ID mismatch: expected %d, got %d", chainID, clientChainID.Uint64()))
			}
		}
		// If there was an error or chain ID mismatch, close the client and continue to the next URL
		client.Close()
	}

	return nil, fmt.Errorf("failed to connect to the RPC URL, either the RPC is wrong or the network is not supported")
}

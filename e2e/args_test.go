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
package e2e

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestE2E_ValidArgs_Network_Holesky(t *testing.T) {
	var cmd *exec.Cmd
	var containerID string
	var err error

	e2eTest := newe2eTestCase(
		t,
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			containerID, err = startPushgatewayContainer(t, 9089)
			if err != nil {
				stopPushgatewayContainer(t, containerID)
			}
			return map[string]string{
				"EIGEN_MINTER_LOG_LEVEL":   "debug",
				"EIGEN_MINTER_PRIVATE_KEY": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			}, err
		},
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = runCommandCMD(t, binaryPath, "eigen-minter", "--network", "holesky", "--pushgateway-url", "http://localhost:9089")
			time.Sleep(5 * time.Second)
			return cmd
		},
		func(t *testing.T) {
			defer stopPushgatewayContainer(t, containerID)
			checkPushgatewayMetrics(t, expectedMetrics, 9089)

			cmd.Process.Signal(os.Interrupt)

			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				assert.NoError(t, err)
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill()
			}
		},
	)

	e2eTest.run()
}

func TestE2E_ValidArgs_Network_Mainnet(t *testing.T) {
	var cmd *exec.Cmd
	var containerID string
	var err error

	e2eTest := newe2eTestCase(
		t,
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			containerID, err = startPushgatewayContainer(t, 9091)
			if err != nil {
				stopPushgatewayContainer(t, containerID)
			}
			return map[string]string{
				"EIGEN_MINTER_LOG_LEVEL":   "debug",
				"EIGEN_MINTER_PRIVATE_KEY": "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			}, err
		},
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = runCommandCMD(t, binaryPath, "eigen-minter", "--network", "mainnet", "--pushgateway-url", "http://localhost:9091")
			time.Sleep(10 * time.Second) // Execution on mainnet takes more time
			return cmd
		},
		func(t *testing.T) {
			defer stopPushgatewayContainer(t, containerID)
			checkPushgatewayMetrics(t, expectedMetrics, 9091)

			cmd.Process.Signal(os.Interrupt)

			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				assert.NoError(t, err)
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill()
			}
		},
	)

	e2eTest.run()
}

func TestE2E_ValidArgs_CustomRPC(t *testing.T) {
	var cmd *exec.Cmd
	var containerID string
	var err error

	e2eTest := newe2eTestCase(
		t,
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			containerID, err = startPushgatewayContainer(t, 9092)
			if err != nil {
				stopPushgatewayContainer(t, containerID)
			}
			return map[string]string{
				"EIGEN_MINTER_PUSHGATEWAY_URL": "http://localhost:9092",
				"EIGEN_MINTER_LOG_LEVEL":       "debug",
				"EIGEN_MINTER_PRIVATE_KEY":     "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			}, err
		},
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = runCommandCMD(t, binaryPath, "eigen-minter", "--network", "holesky", "--rpc-endpoint", "https://1rpc.io/holesky")
			time.Sleep(5 * time.Second)
			return cmd
		},
		func(t *testing.T) {
			defer stopPushgatewayContainer(t, containerID)
			checkPushgatewayMetrics(t, expectedMetrics, 9092)

			cmd.Process.Signal(os.Interrupt)

			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				assert.NoError(t, err)
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill()
			}
		},
	)

	e2eTest.run()
}

func TestE2E_ValidEnv_All(t *testing.T) {
	var cmd *exec.Cmd
	var containerID string
	var err error

	e2eTest := newe2eTestCase(
		t,
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			containerID, err = startPushgatewayContainer(t, 9093)
			if err != nil {
				return map[string]string{}, err
			}

			return map[string]string{
				"EIGEN_MINTER_PUSHGATEWAY_URL":  "http://localhost:9093",
				"EIGEN_MINTER_NETWORK":          "holesky",
				"EIGEN_MINTER_RPC_ENDPOINT":     "https://ethereum-holesky-rpc.publicnode.com",
				"EIGEN_MINTER_CONTRACT_ADDRESS": "0x8DaaE33cB2da8dA23595ADB19f271EF41E34bd8C",
				"EIGEN_MINTER_PRIVATE_KEY":      "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
				"EIGEN_MINTER_LOG_LEVEL":        "debug",
			}, nil
		},
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = runCommandCMD(t, binaryPath, "eigen-minter")
			time.Sleep(5 * time.Second)
			return cmd
		},
		func(t *testing.T) {
			defer stopPushgatewayContainer(t, containerID)
			checkPushgatewayMetrics(t, expectedMetrics, 9093)

			cmd.Process.Signal(os.Interrupt)

			done := make(chan error, 1)
			go func() {
				done <- cmd.Wait()
			}()

			select {
			case err := <-done:
				assert.NoError(t, err)
			case <-time.After(5 * time.Second):
				t.Error("Process did not exit within the timeout period")
				cmd.Process.Kill()
			}
		},
	)

	e2eTest.run()
}

func TestE2E_InvalidArgs_Network(t *testing.T) {
	var cmd *exec.Cmd
	var containerID string
	var err error

	e2eTest := newe2eTestCase(
		t,
		func(t *testing.T, binaryPath string) (map[string]string, error) {
			containerID, err = startPushgatewayContainer(t, 9094)
			if err != nil {
				stopPushgatewayContainer(t, containerID)
			}
			return map[string]string{
				"EIGEN_MINTER_PUSHGATEWAY_URL": "http://localhost:9094",
				"EIGEN_MINTER_LOG_LEVEL":       "debug",
				"EIGEN_MINTER_PRIVATE_KEY":     "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
			}, err
		},
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = runCommandCMD(t, binaryPath, "eigen-minter", "--network", "invalid_network")
			return cmd
		},
		func(t *testing.T) {
			defer stopPushgatewayContainer(t, containerID)

			err := cmd.Wait()
			assert.Error(t, err, "eigen-minter command should fail with invalid network")
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "eigen-minter command should fail with invalid network")
		},
	)

	e2eTest.run()
}

func TestE2E_InvalidArgs_PrivateKey(t *testing.T) {
	var cmd *exec.Cmd

	e2eTest := newe2eTestCase(
		t,
		nil,
		func(t *testing.T, binaryPath string) *exec.Cmd {
			cmd = runCommandCMD(t, binaryPath, "eigen-minter", "--network", "holesky", "--private-key", "invalid_key", "--log-level", "debug")
			return cmd
		},
		func(t *testing.T) {
			err := cmd.Wait()
			assert.Error(t, err, "eigen-minter command should fail with invalid private key")
			assert.Equal(t, 1, cmd.ProcessState.ExitCode(), "eigen-minter command should fail with invalid private key")
		},
	)

	e2eTest.run()
}

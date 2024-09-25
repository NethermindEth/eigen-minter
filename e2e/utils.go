package e2e

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/cenkalti/backoff/v4"
)

func runCommand(t *testing.T, path string, binaryName string, args ...string) error {
	_, _, err := runCommandOutput(t, path, binaryName, args...)
	return err
}

func runCommandCMD(t *testing.T, path string, binaryName string, args ...string) *exec.Cmd {
	t.Helper()
	t.Logf("Running command: %s %s", binaryName, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		t.Fatalf("Failed to start command: %s %s", binaryName, strings.Join(args, " "))
	}
	return cmd
}

func runCommandOutput(t *testing.T, path string, binaryName string, args ...string) ([]byte, *exec.Cmd, error) {
	t.Helper()
	t.Logf("Binary path: %s", path)
	t.Logf("Running command: %s %s", binaryName, strings.Join(args, " "))
	cmd := exec.Command(path, args...)
	out, err := cmd.CombinedOutput()
	t.Logf("===== OUTPUT =====\n%s\n==================", out)
	return out, cmd, err
}

func logAndPipeError(t *testing.T, prefix string, err error) error {
	t.Helper()
	if err != nil {
		t.Log(prefix, err)
	}
	return err
}

func checkGoInstalled(t *testing.T) {
	t.Helper()
	err := exec.Command("go", "version").Run()
	if err != nil {
		t.Fatalf("error checking Go installation: %v", err)
	} else {
		t.Logf("Go installed")
	}
}

// Make a GET request to the given URL. Uses exponential retries with backoff.
func getRequest(t *testing.T, url string, retryDuration time.Duration) (*http.Response, error) {
	t.Helper()
	var response *http.Response

	// Adding exponential retry
	b := backoff.NewExponentialBackOff()
	b.MaxElapsedTime = retryDuration

	err := backoff.Retry(func() (err error) {
		// To make a request with custom headers, use NewRequest and Client.Do.
		response, err = http.Get(url)
		if err != nil {
			t.Logf("request failed. Error: %v", err)
			t.Logf("Retrying request")
			return err
		} else if response.StatusCode != 200 {
			t.Logf("bad response, got: %d", response.StatusCode)
		}
		return nil
	}, b)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func startPushgatewayContainer(t *testing.T, port int) (string, error) {
	t.Helper()
	cmd := exec.Command("docker", "run", "-d", "-p", fmt.Sprintf("%d:9091", port), "prom/pushgateway")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to start pushgateway container: %v", err)
	}
	containerID := strings.TrimSpace(string(out))
	return containerID, nil
}

func stopPushgatewayContainer(t *testing.T, containerID string) error {
	t.Helper()
	cmd := exec.Command("docker", "stop", containerID)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to stop pushgateway container: %v", err)
	}

	cmd = exec.Command("docker", "rm", containerID)
	_, err = cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to remove pushgateway container: %v", err)
	}

	return nil
}

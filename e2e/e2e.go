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
	"path/filepath"
	"syscall"
	"testing"
)

type (
	e2eArranger func(t *testing.T, appPath string) (map[string]string, error)
	e2eAct      func(t *testing.T, appPath string) *exec.Cmd
	e2eAssert   func(t *testing.T)
)

type e2eTestCase struct {
	t                    *testing.T
	testDir              string
	repoPath             string
	binaryName           string
	arranger             e2eArranger
	act                  e2eAct
	assert               e2eAssert
	pid                  int
	envVars              []string
}

func newe2eTestCase(t *testing.T, arranger e2eArranger, act e2eAct, assert e2eAssert) *e2eTestCase {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	tc := &e2eTestCase{
		t:          t,
		testDir:    t.TempDir(),
		repoPath:   filepath.Dir(wd),
		binaryName: "eigen-minter",
		arranger:   arranger,
		act:        act,
		assert:     assert,
	}
	t.Logf("Creating new E2E test case (%p). TestDir: %s", tc, tc.testDir)
	checkGoInstalled(t)
	tc.installGoModules()
	tc.build()
	return tc
}

func (e *e2eTestCase) binaryPath() string {
	binaryName := e.binaryName
	return filepath.Join(e.testDir, binaryName)
}

func (e *e2eTestCase) installGoModules() {
	e.t.Helper()
	cmd := exec.Command("go", "mod", "download")
	cmd.Dir = e.repoPath
	e.t.Logf("Installing Go modules in %s", e.repoPath)
	if err := cmd.Run(); err != nil {
		e.t.Fatalf("error installing Go modules: %v", err)
	} else {
		e.t.Logf("Go modules installed")
	}
}

func (e *e2eTestCase) build() {
	e.t.Helper()
	e.t.Logf("Building binary to %s", e.binaryPath())
	err := exec.Command("go", "build", "-o", e.binaryPath(), filepath.Join(e.repoPath, "cmd", e.binaryName, "main.go")).Run()
	if err != nil {
		e.t.Fatalf("error building binary: %v", err)
	} else {
		e.t.Logf("binary built")
	}
}

func (e *e2eTestCase) run() {
	// Cleanup environment before and after test
	e.cleanup()
	defer e.cleanup()
	if e.arranger != nil {
		env, err := e.arranger(e.t, e.binaryPath())
		if err != nil {
			e.t.Fatalf("error in Arrange step: %v", err)
		}

		for key, value := range env {
			e.setEnvVariable(key, value)
		}
	}
	if e.act != nil {
		cmd := e.act(e.t, e.binaryPath())
		e.pid = cmd.Process.Pid
	}
	if e.assert != nil {
		e.assert(e.t)
	}
}

func (e *e2eTestCase) setEnvVariable(key, value string) {
	e.t.Helper()
	e.envVars = append(e.envVars, key)
	os.Setenv(key, value)
}

func (e *e2eTestCase) cleanup() {
	if e.pid != 0 {
		process, err := os.FindProcess(e.pid)
		if err == nil {
			process.Signal(syscall.SIGTERM)
		}
	}

	// Restore environment variables
	for _, env := range e.envVars {
		os.Unsetenv(env)
	}
}

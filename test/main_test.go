package test

import (
	"os/exec"
	"testing"
)

func TestMain(t *testing.T) {
	cmd := exec.Command("go", "run", "/home/gamp/L2E Fellowship/piscine-prompt/ascii-art/cmd/main.go", "Hello House")
	result, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("Mani execution failed!.")
	}

	t.Logf(string(result))
}

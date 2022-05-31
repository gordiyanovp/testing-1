package coverage

import (
	"bytes"
	"os/exec"
	"testing"
)

func TestStudents(t *testing.T) {
	cmd := exec.Command("go", "test", "-cover", "-run", "^TestConcatenator$")
	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		t.Errorf("Error during CMD execution: %v", err)
	}
	coverage := stdout.Bytes()
	if !bytes.Contains(coverage, []byte("100.0%")) {
		t.Error("Error: coverage is not 100.0%")
	}
}

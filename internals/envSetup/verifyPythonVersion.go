package envSetup

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func VerifyPythonVersion() bool {
	cmd := exec.Command("python3", "--version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("python3 command not found: %v", err)
	}

	version := strings.Fields(string(output))

	if len(version) < 2 {
		log.Fatal("unexpected version output:", string(output))
	}

	parts := strings.Split(version[1], ".")

	major, _ := strconv.Atoi(parts[0])
	minor, _ := strconv.Atoi(parts[1])
	patch, _ := strconv.Atoi(parts[2])

	if major > 3 ||
		(major == 3 && minor > 12) ||
		(major == 3 && minor == 12 && patch >= 12) {
		return true
	} else {
		return false
	}

}

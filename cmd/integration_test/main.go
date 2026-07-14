package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("========================================")
	fmt.Println("   FINBOARD INTEGRATION TESTS (DOCKER)")
	fmt.Println("========================================")
	fmt.Println()

	fmt.Println("IMPORTANT: Docker must be running!")
	fmt.Println("These tests require PostgreSQL container.")
	fmt.Println()

	fmt.Println("Running integration tests...")
	fmt.Println()

	cmd := exec.Command("go", "test", "-v", "-count=1", "./src/integration/...")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println()
		fmt.Println("========================================")
		fmt.Println("Integration tests failed or Docker not available.")
		fmt.Println("========================================")
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("========================================")
	fmt.Println("All integration tests passed!")
	fmt.Println("========================================")
}

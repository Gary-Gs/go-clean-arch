package usecase

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Initializing tests ...")

	exitVal := m.Run()

	fmt.Println("Shutting down tests ...")
	os.Exit(exitVal)
}

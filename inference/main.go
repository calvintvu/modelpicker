//go:build ignore

package main

import (
	"fmt"
	"os"

	"github.com/calvinvu/ez-model-router/inference"
)

func main() {
	session, err := inference.LoadInferenceSession()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load inference session:", err)
		os.Exit(1)
	}
	defer session.CloseInferenceSession()

	fmt.Println("inference session loaded successfully")
}

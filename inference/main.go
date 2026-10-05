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

	ticketText := "I cancelled my subscription two weeks ago and I still have not received my refund."
	questionText := "Which team should handle this ticket?"
	options := []string{"billing: payments, refunds, invoices", "support: product help and bugs", "sales: new purchases and upgrades"}
	temperature := 1.7601518630981445
	probs, err := session.Choose(questionText, ticketText, options, temperature)
	maxValue := float64(0)
	maxValueIndex := -1
	for i := range probs {
		if probs[i] > maxValue {
			maxValue = float64(probs[i])
			maxValueIndex = i
		}
	}
	if maxValue != -1 {
		fmt.Println(options[maxValueIndex])
	}
}

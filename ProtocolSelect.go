package main

import (
	"fmt"
	"strings"
)

func ProtocolSelect() (string, error) {

	fmt.Println("Please enter a port or protocol by name:")

	input := ""

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Printf("Error reading input: %w", err)
	}

	inputToUpper := strings.ToUpper(input)

	p, err := ProtocolBuilder(inputToUpper)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("Protocol: %s Port: %s", p.Name, p.Port), nil
}

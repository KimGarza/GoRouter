package main

import "fmt"

func ProtocolBuilder(input string) (*Protocol, error) {

	switch input {
	case "80", "HTTP":
		return &Protocol{Name: "HTTP", Port: "80"}, nil
	case "22", "SSH":
		return &Protocol{Name: "SSH", Port: "22"}, nil
	case "3389", "RDP":
		return &Protocol{Name: "RDP", Port: "3389"}, nil
	default:
		return nil, fmt.Errorf("Input did not match any protocols in directory")
	}
}

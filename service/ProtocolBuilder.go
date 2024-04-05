package service

import (
	"fmt"
	"go-router/model"
)

func ProtocolBuilder(input string) (*model.Protocol, error) {

	switch input {
	case "80", "HTTP":
		return &model.Protocol{Name: "HTTP", Port: "80"}, nil
	case "22", "SSH":
		return &model.Protocol{Name: "SSH", Port: "22"}, nil
	case "3389", "RDP":
		return &model.Protocol{Name: "RDP", Port: "3389"}, nil
	default:
		return nil, fmt.Errorf("Input did not match any protocols in directory")
	}
}

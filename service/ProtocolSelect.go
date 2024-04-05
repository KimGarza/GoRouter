package service

import (
	"fmt"
	"go-router/model"
	"strings"
)

func ProtocolSelect(value string) (model.Protocol, error) {

	inputToUpper := strings.ToUpper(value)

	p, err := ProtocolBuilder(inputToUpper)
	if err != nil {
		return model.Protocol{}, err
	}

	fmt.Sprintf("Protocol: %s Port: %s", p.Name, p.Port)

	return *p, nil
}

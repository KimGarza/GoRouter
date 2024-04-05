package service

import (
	"fmt"
	"go-router/model"
)

func ProtocolBuilder(input string) (*model.Protocol, error) {

	switch input {
	case "80", "HTTP":
		return &model.Protocol{
				Name:            "HTTP",
				Port:            "80",
				OSI:             "Application",
				ConnnectionType: []string{"stateful", "stateless"},
				Security:        "none",
				Payload:         []string{"HTML", "CSS", "JavaScript", "Multimedia", "JSON", "XML", "Binary Data", "Form Data"}, // content type
				Headers:         []string{"Content-Type", "Content-Length", "Request Headers", "Repsonse Headers", "Security Headers", "Entity Headers"},
				Standards:       []string{"RFC 1945", "RFC 2616", "RFC 7230", "RFC 7231", "RFC 7232", "RFC 7233", "RFC 7234", "RFC 7235", "RFC 7540", "RFC 9114"},
				Description:     "HTTP is an application layer protocol designed to transfer information between networked devices and runs on top of other layers of the network protocol stack. -Cloudflare",
				LatestVersion:   3},
			nil
	// case "22", "SSH":
	// 	return &model.Protocol{Name: "SSH", Port: "22"}, nil
	// case "3389", "RDP":
	// 	return &model.Protocol{Name: "RDP", Port: "3389"}, nil
	default:
		return nil, fmt.Errorf("Input did not match any protocols in directory")
	}
}

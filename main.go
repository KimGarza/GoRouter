package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Hello Mars!")

	// take input of either port or protocol to upper, predict the other based on entered input
	p, err := ProtocolSelect()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(p)

	// once protocol is decided, type send ( eventually to be click on send )
	// now router will execute. Will show your data and break it apart in to frames. Disects the frame to depict each element in the frame/data quantity
	// replace graphic visual of connection from computer to dest through router or just the signal to go through router with ascii or similar like
	// internet <----------- ROUTER <----------- localPC
	// breif explain each element in process of OSI model and variation between protocols. Ex: rdp will have a constant connection, http will send and be done
}

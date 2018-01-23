package main

import (
	"fmt"
	"log"
	"os"

	"github.com/byuoitav/password-utility/passwords"
	"github.com/fatih/color"
)

func main() {

	//fetches the arguments without the path to the executable
	args := os.Args[1:]

	result, err := passwords.GetPassword(args[0])
	if err != nil {
		msg := fmt.Sprintf("passord not found: %s", err.Error())
		log.Printf("%s", color.HiRedString("[password-utility] %s", msg))
	}

	fmt.Printf(result)

}

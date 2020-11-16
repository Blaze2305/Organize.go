package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Blaze2305/Organize.go/internal/utils"
)

func main() {

	_, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	err = utils.MakeInitFolders()
	if err != nil {
		log.Fatal(err)
		return
	}

}

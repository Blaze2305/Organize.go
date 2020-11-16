package main

import (
	"log"
	"os"

	"github.com/Blaze2305/Organize.go/internal/utils"
)

func main() {

	_, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	err = utils.MakeInitFolders()
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = utils.OrganizeFolder()
	if err != nil {
		log.Fatalln(err)
		return
	}

}

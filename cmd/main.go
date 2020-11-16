package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Blaze2305/Organize.go/internal/utils"
)

func main() {
	fmt.Println("Organize this folder? Action is irreversible y / n")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	if string(input[0]) != "y" {
		fmt.Println("Good Bye :D")
		return
	}
	err := utils.MakeInitFolders()
	if err != nil {
		log.Fatalln(err)
		return
	}

	err = utils.OrganizeFolder()
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Done")

}

package main // import "github.com/xmlking/go_project/cmd/great1"

import (
	"flag"
	"fmt"

	"github.com/xmlking/go_project/utils"
)

func init() {
	fmt.Println("Command ==>", "Greet1")
}

func main() {
	name := flag.String("name", "", "name")
	flag.Parse()

	if err := utils.Greet(*name); err != nil {
		fmt.Printf("Failed to greet you: %v", err)
	}
}

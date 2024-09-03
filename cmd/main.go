package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/luisnogit/luisnobd.git/cmd/Utils"
	"github.com/luisnogit/luisnobd.git/cmd/parser"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	wellcomeprompt()
	for {
		fmt.Print(">>")
		if !input.Scan() || input.Err() != nil {
			log.Fatal(input.Err())
		}
		if strinput := input.Text(); strinput[0] == '.' {
			if utils.MetaCommand(strinput) != nil {
				fmt.Println("Invalid meta command", strinput)
				continue
			} else {
				continue
			}
		}
		statement, err := parser.NewStatement(input.Text())
		if err != nil {
			log.Println(err)
			continue
		}
		execute(statement)

	}
}
func wellcomeprompt() {
	fmt.Printf("LuisnoBD - v0.0.1\n System page size: %d\n Wellcome:", os.Getpagesize())
}



func execute(statement utils.Statement) {
	switch statement.Tp {
	case utils.INSERT:
		fmt.Println("This would insert")
	case utils.SELECT:
		fmt.Println("This would select")
	}

}

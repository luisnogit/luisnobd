package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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
			if MetaCommand(strinput) != nil {
				fmt.Println("Invalid meta command", strinput)
				continue
			} else {
				continue
			}
		}
		statement, err := NewStatement(input.Text())
		if err != nil {
			log.Println(err)
			continue
		}
		execute(statement)

	}
}
func wellcomeprompt() {
	fmt.Println("LuisnoBD - v0.0.1\n Wellcome:")
}

func MetaCommand(command string) error {
	if command == ".exit" {
		os.Exit(0)
	} else {
		return ErrInvalidMeta
	}
	return nil
}

type StatementType uint8

const (
	INSERT = iota
	SELECT
)

type Statement struct {
	Tp    StatementType
	Input string
}

func execute(statement Statement) {
	switch statement.Tp {
	case INSERT:
		fmt.Println("This would insert")
	case SELECT:
		fmt.Println("This would select")
	}

}

func NewStatement(input string) (Statement, error) {
	if strings.Contains(input, "SELECT") {
		return Statement{Tp: SELECT, Input: input}, nil
	} else if strings.Contains(input, "INSERT")  {
		return Statement{Tp: INSERT, Input: input}, nil
	}
	return Statement{}, ErrUnrecognizedStm
}

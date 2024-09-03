package parser

import (
	"fmt"
	"os"
	"strings"

	utils "github.com/luisnogit/luisnobd.git/cmd/Utils"
)

func MetaCommand(command string) error {
	if command == ".exit" {
		os.Exit(0)
	} else {
		return utils.ErrInvalidMeta
	}
	return nil
}

func NewStatement(input string) (utils.Statement, error) {
	var operation string
	var table utils.Row
	_, err := fmt.Sscanf(input, "%s %d %s %s", &operation, &table.Id, &table.Username, &table.Email)
	if err != nil {
		return utils.Statement{}, utils.ErrUnrecognizedStm
	}
	fmt.Println(table)
	if strings.ToUpper(operation) == "SELECT" {
		return utils.Statement{Tp: utils.SELECT, Input: input}, nil
	} else if strings.ToUpper(operation) == "INSERT" {

		return utils.Statement{Tp: utils.INSERT, Input: input}, nil
	}
	return utils.Statement{}, utils.ErrUnrecognizedStm
}

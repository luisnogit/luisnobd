package utils

import (
	"bytes"
	"encoding/gob"
	"os"
)

func Serialize(param any) ([]byte, error) {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	if err := encoder.Encode(param); err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func Deserialize(param []byte, result any) error {
	var buff = bytes.NewBuffer(param)
	decoder := gob.NewDecoder(buff)
	if err := decoder.Decode(result); err != nil {
		return err
	}
	return nil

}
func MetaCommand(command string) error {
	if command == ".exit" {
		os.Exit(0)
	} else {
		return ErrInvalidMeta
	}
	return nil
}

package utils

import (
	"testing"
)

func TestSerialization(t *testing.T) {
	var testcase Row = Row{
		Id:       1,
		Email:    "luis@gmail.com",
		Username: "luis lindo",
	}
	result, err := Serialize(testcase)
	if err != nil {
		t.Fatalf("deu pau no processo: %v", err)
	}
	var final Row
	err = Deserialize(result, &final)
	if err != nil {
		t.Fatalf("deu pau no desserializar %v", err)
	}

	if testcase != final {
		t.Fatalf("OS DOIS ESTÃO ERRADOS %v || %v\n", testcase, final)
	}
	t.Log("testcases used:", final, result)
}

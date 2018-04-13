package main

import (
	"testing"
	"fmt"
)

func TestHandleArguments(t *testing.T) {
	t.Log("Call app - Checking for missing files names")
	args := []string{"nameProgram", "file01.cvc", "file02.cvc"}

	if e, _ := HandleArguments(args); e != nil {
		t.Error("Filenames requires to call this app")
	}
}

func TestCSVProcessing(t *testing.T) {
	t.Log("Call app - Checking for missing files")
	if e, _ := ProcessCSV("data/contas.csv"); e != nil {
		fmt.Println(e)
		t.Error("CSV files doesn't exist")
	}
}

func TestCSVHandling(t *testing.T) {
	t.Log("Call app - Handle with incorrect csv file")
	if e, _ := HandleCSV("wrong.csv"); e != nil {
		t.Error("Doesn't work with incorrect csv file")
	}
}

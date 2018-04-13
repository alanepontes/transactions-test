package main

import (
    "testing"
)

func TestTransactionCreation(t *testing.T) {
	t.Log("Creating new transaction.")
	if e, _ := NewTransaction("1", "string"); e == nil {
	    t.Error("Invalid data")
	}
}
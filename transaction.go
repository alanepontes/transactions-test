package main

import (
	"strconv"
)

type Transaction struct {
	Id  int
	Value int
}

func NewTransaction(Id string, Value string) (error, Transaction) {
	var err error
	var convId, convValue int
	
	if convId, err = strconv.Atoi(Id); err != nil {
		return &CustomError{Id, "Cannot convert value to int"}, Transaction{}
	}

	if convValue, err = strconv.Atoi(Value); err != nil {
		return &CustomError{Value, "Cannot convert value to int"}, Transaction{}
	}

	return nil, Transaction {
		Id: convId,
		Value: convValue,
	}
}
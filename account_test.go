package main

import (
    "testing"
    "github.com/shopspring/decimal"
    "fmt"
)

func TestAccountCreation(t *testing.T) {
	t.Log("Creating new account.")
	if e, _ := NewAccount("1", "string"); e == nil {
		t.Error("Invalid data")
	}
}

func TestExpectedBalance(t *testing.T) {
    accountId := 1
    balance := 2000
    var transactions []Transaction
    _, t1:= NewTransaction("1", "1000")
    _, t2 := NewTransaction("1", "1000")
    _, t3 := NewTransaction("1", "-2000")
    _, t4 := NewTransaction("1", "-1230")
    transactions = append(transactions, t1)  
    transactions = append(transactions, t2) 
    transactions = append(transactions, t3)  
    transactions = append(transactions, t4)  
    
    expectedBalance := decimal.NewFromFloat(7.7)
    currentBalance := CurrentBalance(accountId, balance, transactions)

    t.Log("Test result from CurrentBalance")
    if !(currentBalance.StringFixed(2) == expectedBalance.StringFixed(2)) { 
        t.Error(fmt.Sprintf("Current balance returned wrong value - currentBalance: %s | expected currentBalance %s", currentBalance.StringFixed(2), expectedBalance.StringFixed(2)))      
    }
}
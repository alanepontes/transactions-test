package main

import (
    "strconv"
		"github.com/shopspring/decimal"
)

const fine float64 = 5.0
type Account struct {
	Id int
	Balance int
}

func NewAccount(Id string, Balance string) (error, Account) {
	var err error
	var convId, convBalance int
	
	if convId, err  = strconv.Atoi(Id); err != nil {
		return &CustomError{Id, "Cannot convert value to int"}, Account{}
	}

	if convBalance, err = strconv.Atoi(Balance); err != nil {
		return &CustomError{Balance, "Cannot convert value to int"}, Account{}
	}

	return nil, Account {
		Id: convId,
		Balance: convBalance,
	}
}

func hasFine(balance decimal.Decimal) bool {
	return balance.LessThan(decimal.NewFromFloat(0))
}

func applyFine(balance decimal.Decimal) decimal.Decimal {
	return decimal.Sum(balance, decimal.NewFromFloat(-fine))
}


func filterTransactions(accountId int, transactions []Transaction) []Transaction {
	var passed []Transaction

	for _, transaction := range(transactions) {
			if (transaction.Id == accountId) {
					passed = append(passed, transaction)
			}
	}
	return passed
}

func appliedTransactions(balance int, transactions []Transaction) decimal.Decimal {
	currentBalance := FormatToMoney(balance)
	
	for _, transaction := range(transactions) {
		  currentBalance = decimal.Sum(currentBalance, FormatToMoney(transaction.Value))
			if (hasFine(currentBalance)) {
					currentBalance = applyFine(currentBalance)
			}
	}
	return currentBalance
}
	
func CurrentBalance(accountId int, balance int, transactions []Transaction) decimal.Decimal {
	transactionsInAccount := filterTransactions(accountId, transactions)
	if (len(transactionsInAccount) == 0) {
			return FormatToMoney(balance)
	}
	return appliedTransactions(balance, transactionsInAccount)
}
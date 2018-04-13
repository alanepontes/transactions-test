package main

import (
    "os"
    "fmt"
)

func main() {
    if err, nameFiles := HandleArguments(os.Args); err != nil {
        fmt.Println("Failed:", err)
    } else {
        var err error
        var accountFile, transactionFile string
        
        if err, accountFile = ProcessCSV(nameFiles[0]); err != nil {
            fmt.Println("Error:", err)
        }
        
        if err, transactionFile  = ProcessCSV(nameFiles[1]); err != nil {
            fmt.Println("Error:", err)
        }
            
        var dataAccountFile, dataTransactionFile [][]string
        if err, dataAccountFile  = HandleCSV(accountFile); err != nil {
            fmt.Println("Error:", err)
        }
            
        if err, dataTransactionFile  = HandleCSV(transactionFile); err != nil {
            fmt.Println("Error:", err)
        }
        
        var accounts []Account
        for _, line := range(dataAccountFile) {
            var err error
            var account Account

            if err, account = NewAccount(line[0], line[1]); err != nil {
                fmt.Println("Invalid Account data:", err)
                continue
            }

            accounts = append(accounts, account)
        }

        var transactions []Transaction
        for _, line := range(dataTransactionFile) {
            var err error
            var transaction Transaction

            if err, transaction = NewTransaction(line[0], line[1]); err != nil {
                fmt.Println("Invalid Transaction data:", err)
                continue
            }

            transactions = append(transactions, transaction)
        }

        for _, account := range(accounts) {
            currentBalance := CurrentBalance(account.Id, account.Balance, transactions)
            fmt.Println(CSVnify(account.Id, currentBalance))
        }
  }
}

package main

import "fmt"

type BankAccount struct {
    owner   string
    balance int
}

func NewBankAccount(owner string, balance int) *BankAccount {
    return &BankAccount{
        owner:   owner,
        balance: balance,
    }
}

func (a *BankAccount) Deposit(amount int) int {
    if amount > 0 {
        a.balance += amount
        fmt.Printf("Deposited %d\n", amount)
    }
    return a.balance
}

func (a *BankAccount) Withdraw(amount int) int {
    if 0 < amount && amount <= a.balance {
        a.balance -= amount
        fmt.Printf("Withdrew %d\n", amount)
    } else {
        fmt.Println("Insufficient funds")
    }
    return a.balance
}

func main() {
    account := NewBankAccount("Bob", 200)
    account.Deposit(100)
    account.Withdraw(50)
    fmt.Printf("Balance: %d\n", account.balance)
}

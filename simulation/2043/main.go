package main

import "fmt"

type Bank struct {
	bal []int64
}

func Constructor(balance []int64) Bank {
	b := make([]int64, len(balance))
	copy(b, balance)
	return Bank{bal: b}
}

func (this *Bank) valid(account int) bool {
	return account >= 1 && account <= len(this.bal)
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if !this.valid(account1) || !this.valid(account2) {
		return false
	}
	a := account1 - 1
	b := account2 - 1
	if this.bal[a] < money {
		return false
	}
	this.bal[a] -= money
	this.bal[b] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if !this.valid(account) {
		return false
	}
	this.bal[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if !this.valid(account) {
		return false
	}
	a := account - 1
	if this.bal[a] < money {
		return false
	}
	this.bal[a] -= money
	return true
}

func main() {
	b := Bank{}
	fmt.Println(b.Withdraw(3, 10))
}

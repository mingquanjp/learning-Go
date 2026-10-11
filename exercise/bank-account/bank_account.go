package bankaccount

import "sync"

type Account struct {
	balance int64
	isOpen bool
	mu sync.Mutex
}

func Open(amt int64) *Account {
	if amt <0 {
		return nil
	}
	newAcc := Account{
		balance: amt,
		isOpen: true,
	}
	return &newAcc
	
		
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance, a.isOpen
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance + amt < 0 {
		return a.balance, false
	}
	if a.isOpen{
		a.balance += amt
		return a.balance,true
	}

	return a.balance, false
	
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isOpen && amt >=0 && a.balance >= amt{
		a.balance -= amt
		return a.balance, true
	}
	return a.balance, false
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.isOpen {
		a.isOpen = false
		pay := a.balance
		a.balance = 0
		return pay, true
	}
	return a.balance, false
}

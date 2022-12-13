package account

import "sync"

// Define the Account type here.
type Account struct {
	balance int64
	open    bool
	mx      sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}

	return &Account{
		balance: amount,
		open:    true,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.mx.RLock()
	defer a.mx.RUnlock()

	if !a.open {
		return 0, false
	}

	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mx.Lock()
	defer a.mx.Unlock()

	if !a.open {
		return 0, false
	}

	if amount < 0 && a.balance < -amount {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mx.Lock()
	defer a.mx.Unlock()

	if !a.open {
		return 0, false
	}

	oldBalance := a.balance
	a.balance = 0
	a.open = false

	return oldBalance, true
}

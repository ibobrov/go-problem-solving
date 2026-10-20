package pointer

import (
	"errors"
	"math"
)

type Account struct {
	Balance int64
}

func (a *Account) Deposit(amount int64) error {
	if a == nil {
		return errors.New("account is nil")
	}

	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if amount > math.MaxInt64-a.Balance {
		return errors.New("balance overflow")
	}

	a.Balance += amount
	return nil
}

func (a *Account) Withdraw(amount int64) error {
	if a == nil {
		return errors.New("account is nil")
	}

	if amount <= 0 {
		return errors.New("amount must be greater than zero")
	}

	if a.Balance < amount {
		return errors.New("cannot withdraw")
	}

	a.Balance -= amount
	return nil
}

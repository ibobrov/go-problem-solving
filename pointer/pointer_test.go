package pointer

import (
	"strings"
	"testing"
)

func TestAccount_Deposit(t *testing.T) {
	tests := []struct {
		name        string
		account     *Account
		amount      int64
		wantBalance int64
		wantErr     string
	}{
		{
			name:        "deposit to empty account",
			account:     &Account{Balance: 0},
			amount:      100,
			wantBalance: 100,
		},
		{
			name:        "deposit to existing balance",
			account:     &Account{Balance: 200},
			amount:      50,
			wantBalance: 250,
		},
		{
			name:        "deposit zero",
			account:     &Account{Balance: 100},
			amount:      0,
			wantBalance: 100,
			wantErr:     "amount must be greater than zero",
		},
		{
			name:        "deposit negative amount",
			account:     &Account{Balance: 100},
			amount:      -50,
			wantBalance: 100,
			wantErr:     "amount must be greater than zero",
		},
		{
			name:    "nil account",
			account: nil,
			amount:  100,
			wantErr: "account is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Deposit(tt.amount)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("expected error %q, got %q", tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.account == nil {
				return
			}

			if tt.account.Balance != tt.wantBalance {
				t.Errorf(
					"Balance = %d, want %d",
					tt.account.Balance,
					tt.wantBalance,
				)
			}
		})
	}
}

func TestAccount_Withdraw(t *testing.T) {
	tests := []struct {
		name        string
		account     *Account
		amount      int64
		wantBalance int64
		wantErr     string
	}{
		{
			name:        "successful withdrawal",
			account:     &Account{Balance: 200},
			amount:      50,
			wantBalance: 150,
		},
		{
			name:        "withdraw entire balance",
			account:     &Account{Balance: 100},
			amount:      100,
			wantBalance: 0,
		},
		{
			name:        "insufficient funds",
			account:     &Account{Balance: 100},
			amount:      150,
			wantBalance: 100,
			wantErr:     "cannot withdraw",
		},
		{
			name:        "withdraw from empty account",
			account:     &Account{Balance: 0},
			amount:      50,
			wantBalance: 0,
			wantErr:     "cannot withdraw",
		},
		{
			name:        "withdraw zero",
			account:     &Account{Balance: 100},
			amount:      0,
			wantBalance: 100,
			wantErr:     "amount must be greater than zero",
		},
		{
			name:        "withdraw negative amount",
			account:     &Account{Balance: 100},
			amount:      -50,
			wantBalance: 100,
			wantErr:     "amount must be greater than zero",
		},
		{
			name:    "nil account",
			account: nil,
			amount:  100,
			wantErr: "account is nil",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.account.Withdraw(tt.amount)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantErr)
				}
				if !strings.Contains(err.Error(), tt.wantErr) {
					t.Errorf("expected error %q, got %q", tt.wantErr, err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if tt.account == nil {
				return
			}

			if tt.account.Balance != tt.wantBalance {
				t.Errorf(
					"Balance = %d, want %d",
					tt.account.Balance,
					tt.wantBalance,
				)
			}
		})
	}
}

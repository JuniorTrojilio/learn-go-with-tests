package pointeranderrors

import (
	"errors"
	"fmt"
)

// ErrorInsuficientBalance is a default error to insuficient balance
const ErrorInsuficientBalance string = "saldo insuficiente"

// Wallet is a object
type Wallet struct {
	balance Bitcoin
}

// Bitcoin is a currency
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Stringer is a interface to String method
type Stringer interface {
	String() string
}

// Deposite bitcoin to balance
func (w *Wallet) Deposite(value Bitcoin) {
	w.balance += value
}

// Balance return current value
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw bitcoin of balance
func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return errors.New(ErrorInsuficientBalance)
	}

	w.balance -= value
	return nil
}

func main() {

}

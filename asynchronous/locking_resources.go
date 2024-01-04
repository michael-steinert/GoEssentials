package asynchronous

import (
	"fmt"
	"sync"
)

func LockingResources() {
	var waitGroup sync.WaitGroup
	account := NewAccount()
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func() {
			account.Deposit(42)
			account.Withdraw(42)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	fmt.Printf("Account Balance is  %d\n", account.Balance())
}

type Account struct {
	mutex   sync.RWMutex
	balance int
}

func NewAccount() *Account {
	return &Account{
		balance: 0,
	}
}

func (account *Account) Balance() int {
	// Account is only locked when a Function (i.e. GoRoutine) is Writing
	account.mutex.RLock()
	// Account is only unlocked when a Function (i.e. GoRoutine) is Writing
	defer account.mutex.RUnlock()
	return account.balance
}

func (account *Account) Deposit(amount int) {
	// Locking Account, therefore this Function (i.e. GoRoutine) has exclusive Access to this Account
	account.mutex.Lock()
	// Unlocking Account, therefore all other Functions (i.e. GoRoutines) have Access to this Account
	defer account.mutex.Unlock()
	account.balance += amount
}

func (account *Account) Withdraw(amount int) {
	account.mutex.Lock()
	defer account.mutex.Unlock()
	account.balance -= amount
}

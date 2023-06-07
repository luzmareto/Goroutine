package goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 1; i <= 1000; i++ {
		go func() {
			for j := 1; j <= 100; j++ {
				mutex.Lock()
				x = x + 1
				mutex.Unlock()
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("counter = ", x)
}

/*
Mutex berfungsi pada struct yang akan digunakan oleh banyak Goroutine
Kelebihan Mutex adalah mengunci Goroutine sehingga data bisa diolah secara berurutan
*/

type BankAccount struct {
	RWMutex  sync.RWMutex
	Balancee int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balancee = account.Balancee + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balancee
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutext(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("total balance", account.GetBalance())
}

// deadlock
type UserBalance struct {
	sync.Mutex //menulis variable/field Mutex adalah optional
	Name       string
	Balance    int
}

// method
func (user *UserBalance) lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount

}

// uang user1 ditransfer ke user2
func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

func TestDeadlock(t *testing.T) {
	//user1(Luz) mempunyai uang 1jt
	user1 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Luz",
		Balance: 1000000,
	}

	//user2(Mareto) mempunyai uang 1jt
	user2 := UserBalance{
		Mutex:   sync.Mutex{},
		Name:    "Mareto",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 100000)

	time.Sleep(10 * time.Second)

	fmt.Println("User ", user1.Name, "Balance", user1.Balance)
	fmt.Println("User ", user2.Name, "Balance", user2.Balance)
}

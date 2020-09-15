package main

import "fmt"

type bank struct {
	firstName string
	lastName  string
	account   string
	balance   int
}

func (b bank) email() string {
	return b.firstName + "." + b.lastName + "@email.com"
}

func (b bank) display(detail bool) {
	if detail {
		fmt.Printf("Name: %s %s, Account: %s, Balance: %d, email: %s\n", b.firstName, b.lastName, b.account, b.balance, b.email())
	} else {
		fmt.Printf("Name: %s %s, Balance: %d\n", b.firstName, b.lastName, b.balance)
	}

}

func (b *bank) deposit(cash int) {
	b.balance = b.balance + cash
}

func (b *bank) withdraw(cash int) {
	temp := b.balance - cash
	if temp < 0 {
		fmt.Println("Warning: insufficient fund")
	} else {
		b.balance = temp
	}
}

func main() {

	p1 := bank{firstName: "John", lastName: "Smith", account: "111-222", balance: 100}
	p1.display(true)

	p2 := bank{firstName: "Katie", lastName: "Sanero", account: "111-333", balance: 500}
	p2.display(true)

	p3 := bank{firstName: "Lisa", lastName: "Shu", account: "111-555"}
	p3.display(true)

	p3.deposit(1000)
	p3.display(false)
	p2.withdraw(100)
	p2.display(false)
	// p3.balance += 100
	// p3.display(false)
}

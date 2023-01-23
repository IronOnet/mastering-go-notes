package main 

/**
	The Facade design pattern is used to abstract subsystem interfaces with 
	a helper. The facade design pattern is used in scenarios when the number 
	of interfaces increases and the system gets complicated. 

	The facade pattern is made up of the facade class, module classes and 
	a client/ 

	- THe facade delegates the requests from the clien tto the module classes. 
	The facade class hides the complexities of the subsystem logic and rules 

	- Module classes implement the behaviors and functionalities of the module 
	subsystem. 

	- The client invokes the facade method. The facade class functionality can be 
	spread accross multiple packages and assemblies.

	For example account, customer and transaction are the classes that have account, 
	customer, and transaction creation methods. BranchManagerFacade can be used by the 
	client to create an account, customer and transaction.
*/


import (
	"fmt"
)

// Account struct 
type Account struct{
	id string 
	accountType string
}

// Account class method create - creates account given AccountType 
func (account *Account) create(accountType string) *Account{
	account.accountType = accountType 
	return account 
}

// Account class method getById given by id string 
func (account *Account) getById(id string) *Account{
	fmt.Println("getting account by id") 
	return account 
}

// Account class method deleteById given id string 
func (account *Account) deleteById(id string){
	fmt.Println("delete account by id") 

}

type Customer struct{
	name string 
	id int 
}

// Customer class method create - create Customer given name 
func (customer *Customer) create(name string) *Customer{
	fmt.Println("creating a customer") 
	customer.name = name 
	return customer
}

// Transaction struct 
type Transaction struct{
	id string 
	amount float32 
	srcAccountId string 
	destAccountId string 
}

// Transaction class method create Transactions 
func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction{
	fmt.Println("creating transaction") 
	transaction.srcAccountId = srcAccountId 
	transaction.destAccountId = destAccountId 
	transaction.amount = amount 
	return transaction
}

// BranchManagerFacade struct 
type BranchManagerFacade struct{
	account *Account 
	customer *Customer 
	transaction *Transaction 
}

// method NewBranchManagerFacade 
func NewBranchManagerFacade() *BranchManagerFacade{
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

// BranchManagerFacad has the createCustomerAccount method, which calls 
// the create method on the customer class instance, as 

// BranchManagerFacade class method createCustomerAccount 
func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account){
	var customer = facade.customer.create(customerName) 
	var account = facade.account.create(customerName) 
	return customer, account
}

// BranchManagerFacade class method createTransaction 
func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction{
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

// THe main method calls the NewBranchManagerFacade method to create a facade. THe methods 
// on facade are invoked to create customer and account 

func main(){
	var facade = NewBranchManagerFacade() 
	var customer *Customer 
	var account *Account 
	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings") 
	fmt.Println(customer.name) 
	fmt.Println(account.accountType) 
	var transaction = facade.createTransaction("21456", "87345", 1000) 
	fmt.Println(transaction.amount)
}


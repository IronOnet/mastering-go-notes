// The private class data pattern secures the data within a class. 
// This pattern encapsulates the initialization of the class data. 
// the Write priviledges of properties within the private class are 
// protected, and properties are set during construction. 
// The private class pattern prints the exposure of information 
// by securing it in a class that retains the state. The encapsulation 
// of class data initialization is a scenario where this pattrn 
// is applicable. 


package main 

// Account is a class with account details and customer name. AccountDetails is 
// the private attribute of Account, and CustomerName is the public attribute. 
// JSON  marshalling of Account has CustomerName as public property. AccountDetails is 
// the package property in Go

import (
	"encoding/json" 
	"fmt"
)

// AccountDetails struct 
type AccountDetails struct{
	id string 
	accountType string
}

// Account struct 
type Account struct{
	details *AccountDetails 
	CustomerName string
}

// Account class method setDetails 
func (account *Account) setDetails(id string, accountType string){
	account.details = &AccountDetails{id, accountType}
}

// As shown in the following code, the Account class has the getId method, 
// Which returns the id private class attribute: 

// Account class method getId 
func (account *Account) getId() string{
	return account.details.id 
}

// Account class method getAccountType 
func (account *Account) getAccountType() string{
	return account.details.accountType
}

// THE main method calls the Account initializer with CustomerName. The details 
// of the account are set details with the setDetails method 

func main(){
	var account *Account = &Account{CustomerName: "John Smith"} 
	account.setDetails("4532", "current") 
	jsonAccount, _ := json.Marshal(account) 
	fmt.Println("Private class hidden", string(jsonAccount)) 
	fmt.Println("Account ID", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
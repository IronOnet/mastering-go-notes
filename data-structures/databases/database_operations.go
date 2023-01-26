package main 

import (
	"fmt" 
	"database/sql" 
	_ "github.com/go-sql-driver/mysql"
)

// Customer Class 
type Customer struct{
	CustomerId int 
	CustomerName string 
	SSN string
}

// GetConnection method with returns sql.DB 
func GetConnection() (database *sql.DB){
	databaseDriver := "mysql" 
	databaseUser := "newUser"
	databasePass := "newUser" 
	databaseName := "crm" 
	database, error := sql.Open(databaseDriver, 
	databaseUser + ": " + databasePass + "@/" + databaseName) 

	if error != nil{
		panic(error.Error())
	}

	return database 
}

// GetCustomers method returns Customer Array 
func GetCustomers() []Customer{
	var database *sql.DB 
	database = GetConnection() 

	var error error 
	var rows *sql.Rows 
	rows, error = database.Query("SELECT * FROM Customer ORDER BY CustomerId DESC") 
	if error != nil{
		panic(error.Error)
	}

	var customer Customer 
	customer = Customer{} 

	var customers []Customer 
	customers = []Customer{} 
	for rows.Next(){
		var customerId int  
		var customerName string 
		var ssn string 
		error = rows.Scan(&customerId, &customerName, &ssn) 
		if error != nil{
			panic(error.Error())
		}
		customer.CustomerId = customerId 
		customer.CustomerName = customerName 
		customer.SSN = ssn  
		customers = append(customers, customer)
	}

	defer database.Close()
	return customers
}

// InsertCustomer method with parameter customer 
func InsertCustomer(customer Customer){
	var database *sql.DB 
	database = GetConnection() 

	var error error 
	var insert *sql.Stmt 
	insert, error = database.Prepare("INSERT INTO CUSTOMER(CustomerName, SSN) VALUES(?,?)") 
	if error != nil{
		panic(error.Error())
	}
	insert.Exec(customer.CustomerName, customer.SSN)

	defer database.Close()
}

// The update operation 
// tHE UPDATE Operation is as follows. The UpdateCustomer method takes 
// The Customer parameter and creates a prepared statement for the UPDATE 
// sttement. The statement is used to update a customer row in the 
// table 

func UpdateCustomer(customer Customer){
	var database *sql.DB 
	database = GetConnection() 
	var error error 
	var update *sql.Stmt 
	update, error = database.Prepare("UPDATE CUSTOMER SET CustomerName=?, SSN=? WHERE CustomerId=?") 
	if error != nil{
		panic(error.Error())
	}

	update.Exec(customer.CustomerName, customer.SSN, customer.CustomerId) 
	defer database.Close()
}


// The delete operation is as follows. THe DeleteCustomer method takes 
// the Customer parameter and creates a prepared statement for the DELETE 
// statement. The statement is used to execute the deletion of a customer 
// row in the table 

func DeleteCustomer(customer Customer){
	var database *sql.DB 
	database = GetConnection() 
	var error error 
	var delete *sql.Stmt 
	delete, error = database.Prepare("DELETE FROM Customer WHERE Customerid=?") 
	if error != nil{
		panic(error.Error())
	}
	delete.Exec(customer.CustomerId) 
	defer database.Close()
}

func main(){
	//var customers []Customer 
	customers := GetCustomers() 
	fmt.Println("Ccustomers ", customers) 
	var customer Customer 
	customer.CustomerName = "George Thompson" 
	customer.SSN = "23233432" 
	customer.CustomerId = 5 
	//DeleteCustomer(customer)
	UpdateCustomer(customer) 
	customers = GetCustomers() 
	fmt.Println("After Update", customers)
}
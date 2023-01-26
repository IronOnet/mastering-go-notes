package main  

import (
	"fmt" 
	"net/http" 
	"text/template" 
	"log"
	
)

var template_html = template.Must(template.ParseGlob("templates/*")) 

func Home(writer http.ResponseWriter, request *http.Request){
	var customers []Customer 
	customers = GetCustomers() 
	log.Println(customers) 
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Create(writer http.ResponseWriter, request *http.Request){
	// var customers []Customer 
	// customers = GetCustomers() 
	// log.Println(customers) 
	template_html.ExecuteTemplate(writer, "Create", nil)
}

func Insert(writer http.ResponseWriter, request *http.Request){
	var customer Customer 
	customer.CustomerName = request.FormValue("customername") 
	customer.SSN = request.FormValue("ssn") 
	var customers []Customer 
	customers = GetCustomers() 
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Alter(writer http.ResponseWriter, request *http.Request){
	var customer Customer 
	var customerId int 
	var customerIdStr = request.FormValue("id") 
	fmt.Scanf(customerIdStr, "%d", &customerId) 
	customer.CustomerId = customerId 
	customer.CustomerName = request.FormValue("customername") 
	customer.SSN = request.FormValue("ssn") 
	UpdateCustomer(customer) 
	var customers []Customer 
	customers = GetCustomers() 
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func Update(writer http.ResponseWriter, request *http.Request){

	var customerId int 
	var customerIdStr string 
	customerIdStr = request.FormValue("id") 
	fmt.Sscanf(customerIdStr, "%d", &customerId) 
	var customer Customer
	customer = GetCustomerById(customerId) 
	template_html.ExecuteTemplate(writer, "update", customer)
}

func Delete(writer http.ResponseWriter, request *http.Request){
	var customerId int 
	var customerIdStr string 
	customerIdStr = request.FormValue("id") 
	fmt.Sscanf(customerIdStr, "%d", &customerId) 
	var customer Customer 
	customer = GetCustomerById(customerId) 
	DeleteCustomer(customer) 
	var customers []Customer 
	customers = GetCustomers() 
	template_html.ExecuteTemplate(writer, "Home", customers)
}

func View(writer http.ResponseWriter, request *http.Request){
	var customerId int 
	var customerIdStr string 
	customerIdStr = request.FormValue("id") 
	fmt.Sscanf(customerIdStr, "%d", &customerId) 
	var customer Customer
	customer = GetCustomerById(customerId) 
	fmt.Println(customer) 
	var customers []Customer 
	customers = []Customer{customer} 

	template_html.ExecuteTemplate(writer, "View", customers)
}

func main(){
	log.Println("Server started on: http://localhost:8000") 
	http.HandleFunc("/", Home) 
	http.HandleFunc("/alter", Alter) 
	http.HandleFunc("/update", Update) 
	http.HandleFunc("/view", View) 
	http.HandleFunc("/insert", Insert)  
	http.HandleFunc("/delete", Delete) 
	http.ListenAndServe(":8000", nil)
}
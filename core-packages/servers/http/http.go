package main 

// A HTTP server impl 

import (
	"net/http" 
	"io"
) 

func hello(res http.ResponseWriter, req *http.Request){
	res.Header().Set(
		"Content-Type", 
		"text/html",
	)

	io.WriteString(
		res, 
		`<doctype html>
		<html>
		    <head>
			   <title>Hello World</title>
			</head>
			<body>
				Hello World! 
			</body> 
		</html>`,
	)
}

func runHttpServer(){
	http.HandleFunc("/hello", hello) 
	http.ListenAndServe(":9000", nil)
}

func main(){
	go runHttpServer()
}
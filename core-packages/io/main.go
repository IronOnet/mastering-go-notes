package main 

import (
	"bytes" 
	"fmt" 
	"os"
	"io/ioutil" 
	"path/filepath"
)
//import "io"

/**
	The IO package consists of a few functions, but mostly interfaces 
	used in other packages. The two main interfaces are Reader and Writer
	Readers support reading via the Read method. Writers support writting 
	via the Write method
**/

// To read or write to a []bytre or a string you can use 
// the Buffer struct found in the bytes package
func doSomeIo(){
	var buf bytes.Buffer 
	fmt.Println(buf.Write([]byte("test")))
}

func fileOpener(){
	file, err := os.Open("some.txt") 
	if err != nil{
		// handle error here 
		fmt.Printf("An Error Happened While OPening the file %q", err.Error())
		return 
	}
	// We use defer file.Close() right after opening the file 
	// to make sure the file is closed as soon as the function 
	// completes
	defer file.Close()

	// get the file size 
	stat, err := file.Stat() 
	if err != nil{
		return 
	}

	// read the file 
	bs := make([]byte, stat.Size()) 
	_, err = file.Read(bs) 
	if err != nil{
		return 
	}

	str := string(bs) 
	fmt.Println(str)
}

// Reading files is very common, so there's a shorter way to do this: 
func fileOpenerII(){
	bs, err := ioutil.ReadFile("some.txt") 
	if err != nil{
		return 
	}

	str := string(bs) 
	fmt.Println(str)
}

// Here's how we can create a file 
func fileCreator(){
	file, err := os.Create("test.txt") 
	if err != nil{
		return 
	}
	defer file.Close() 
	file.WriteString("Speaking on the topic of race and color, Afrika is the most genetically diverse continent")
}

// TO get the contents of a directlry we use the same 
// os.Open function but give it a directory path instead of 
// a filename. Then we call the Readdir method
func dirOpener(){
	dir, err := os.Open(".") 
	if err != nil{
		return 
	}
	defer dir.Close() 

	fileInfos, err := dir.Readdir(-1) 
	if err != nil{
		return 
	}
	for _, fi := range fileInfos{
		fmt.Println(fi.Name())
	}
}

// Often we want to recursively walk a folder (read the folder's content) 
// all the sub folders, all the sub-sub folder. To make this easier 
// there's a Walk function provided in the path/filepath package 
func walkPath(){
	// THe function you pass to Walk is called for every file and folder 
	// in the root folder (in this case .)
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error{
		fmt.Println(path) 
		return nil 
	})
}

/**
	Files and folders. 
	To open a file in Go use the Open function from the os 
	package. Here is an example of how to read the contents 
	of a file and display them on the terminal
**/

func main(){
	doSomeIo()
	fileOpener() 
	fileOpenerII()
	fileCreator()
	dirOpener()
}
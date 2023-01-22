package main

/**
Hashes & Cryptography

A Hash function takes a set of data and reduces it to a smaller
fixed size. Hashes are frequently used in programing for everything
from lookiing up data to easily detecting changes. Hash functions in
Go are broken into two categories: cryptographic and non-cryptographic.
**/

import (
	"fmt"
	"hash/crc32"
	"io/ioutil"
)

func getHash(filename string) (uint32, error){
	bs , err := ioutil.ReadFile("test1.txt") 
	if err != nil{
		return 0, err 
	}

	h := crc32.NewIEEE() 
	h.Write(bs) 
	return h.Sum32(), nil 
}

func compareFiles(){
  h1, err := getHash("test1.txt") 
  if err != nil{
	return 
  } 
  h2, err := getHash("test2.txt") 
  if err != nil{
	return 
  }

  fmt.Println(h1, h2, h1 == h2)
}
func main(){
	h := crc32.NewIEEE() 
	h.Write([]byte("test")) 
	v := h.Sum32() 
	fmt.Println(v)
}
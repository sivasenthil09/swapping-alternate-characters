package main

import (
	"bufio"
	"os"
	"fmt"
)

func main(){
	var input string
	fmt.Print("Enter a sentence : ")
	reader:=bufio.NewReader(os.Stdin)
	input,_ = reader.ReadString('\n')
	size:=len(input)
	input1:=[]byte(input)
	for i:=2;i<size;i=i+4{
		input1[i],input1[i+1]=input1[i+1],input1[i]
	}
	input=string(input1)
	fmt.Print(input)
}
package main

import(
	"bufio"
	"fmt"
	"os"
)
// a function to check the input
func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	if err:= in.Err(); err!=nil{
		fmt.Println("There is something wrong with your file")
	}
	for in.Scan() {
		fmt.Println("Scanned Text: ", in.Text())
	}
	//fmt.Println("Scanned Text: ",in.Text())
}
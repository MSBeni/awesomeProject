package main

import ( "fmt"; "os" )
import "awesomeProject/printer"
import "awesomeProject/printer/timenow"
import "awesomeProject/printer/pathsplitter"
var speed int

func main(){

	var myage, yourage int
	_ = yourage
	var age = 25
	myage = 20
	theages := 12
	realname, ageofrealname, Accessstatus := "Hossein", 12, true
	//realname := "Hossein"
	printer.Hello()
	fmt.Println(speed)
	fmt.Println("my age is:", myage)
	fmt.Println("the age is:", age)
	fmt.Println(theages, realname)
	fmt.Println(realname, ageofrealname, Accessstatus)
	timenow.Timenowcalc()
	pathsplitter.Splitpath("css/main.css")
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("Path", os.Args[0])
	fmt.Printf("1st Args", os.Args[1])
	fmt.Printf("2nd Args", os.Args[2])
	fmt.Printf("3rd Args", os.Args[3])
	fmt.Printf("Args Length", len(os.Args))
}

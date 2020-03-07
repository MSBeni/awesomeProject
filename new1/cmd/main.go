package main

import (
	"os" 
	"fmt" 
	"strconv"
	"unicode/utf8"
	"math"
)

func main() {
	var s string

	var(
		ab uint8 = 255
		bc = 255
	)

	ab++
	if (int(ab)<bc){
		fmt.Println(ab,bc,"bc is greater than ab")	
	}
	fmt.Println(math.MaxInt8)
	arg := os.Args[1]
	feet, _ := strconv.ParseFloat(arg,64)
	// fmt.Printf("%v\n",feet )
	s = `I am there`
	text :=`
			I am
			there`
	fmt.Println(feet)
	fmt.Println(s)
	fmt.Println(text)

	// Concatenate
	a := "1+2"
	b := 3
	fmt.Println(a+"="+strconv.Itoa(b))
	fmt.Println(strconv.FormatBool(true)+" "+strconv.FormatBool(false))
	fmt.Println(len(a))	
	fmt.Println(len("علی"))
	fmt.Println(utf8.RuneCountInString("علی"))
}


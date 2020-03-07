package main

import (
	"fmt"
	//"strconv"
	"time"
)

func main() {
	type placeholder [5]string

	zero:=placeholder{
		"|||||",
		"|| ||",
		"|| ||",
		"|| ||",
		"|||||",
	}
	one:=placeholder{
		" ||| ",
		"   | ",
		"   | ",
		"   | ",
		"  |||",
	}
	two:=placeholder{
		"||||||",
		"    ||",
		"||||||",
		"||    ",
		"||||||",
	}
	three:=placeholder{
		"||||||",
		"    ||",
		"||||||",
		"    ||",
		"||||||",
	}
	four:=placeholder{
		"|| ||  ",
		"|| ||  ",
		"|||||||",
		"   ||  ",
		"   ||  ",
	}
	five:=placeholder{
		"||||||",
		"||    ",
		"||||||",
		"    ||",
		"||||||",
	}
	six:=placeholder{
		"||||||",
		"||    ",
		"||||||",
		"||  ||",
		"||||||",
	}
	seven:=placeholder{
		"||||||",
		"   || ",
		"  ||  ",
		" ||   ",
		"||    ",
	}
	eight:=placeholder{
		"||||||",
		"||  ||",
		"||||||",
		"||  ||",
		"||||||",
	}
	nine:=placeholder{
		"||||||",
		"||  ||",
		"||||||",
		"    ||",
		"||||||",
	}
	num := [...]placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
	//clock := [...]placeholder{
	//	num[0],num[1],
	//	num[4],num[5],
	//	num[3],num[7],
	//}


	//switch timenow := strconv.Itoa(time.Now().Hour());{
	//case timenow[0] == "0":
	//	_
	//}
	for i:=0; i<100; i++ {
		hour, min, sec := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		clock := [...]placeholder{
			num[hour/10], num[hour%10],
			num[min/10], num[min%10],
			num[sec/10], num[sec%10],
		}
		if time.Now().Hour() > 12 {
			for i, _ := range one {
				for j := range clock {
					fmt.Print(clock[j][i] + " ")
				}
				fmt.Println()
			}
		}
		time.Sleep(time.Second)

	}
	//hour, min, sec := time.Now().Hour(), time.Now().Minute(), time.Now().Second()
	//fmt.Println(hour/10, hour%10,min/10, min%10,sec/10, sec%10)
	//if time.Now().Hour() > 12{
	//	for i, _ := range one{
	//		for j:= range clock{
	//			fmt.Print(clock[j][i] + " ")
	//		}
	//		fmt.Println()
	//	}
	//
	//}
}